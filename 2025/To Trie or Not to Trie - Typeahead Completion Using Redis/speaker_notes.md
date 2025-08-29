# Speaker Notes

Hi, I'm Garrett, a principal engineer on our Platform Team at Skool

In this talk I'll be walking you through how we revamped our type ahead completion for posts and comments on on our site

***NEXT SLIDE***

This should be a familiar feature for anyone who has used a site or product where you can tag or mention other users.

When you start to type the name or handle of another user you’re presented with a nice list of users which match what you’ve typed so far

***NEXT SLIDE***

These queries happen whenever someone starts to @ mention another user by name or handle

The way this was originally implemented it would query postgres using the LIKE operator with some minimal debouncing and cancellation of requests if the user keeps typing

The users table isn’t the only table we need to query for this completion, we also need to support filtering the set of matched users by Skool group membership

***NEXT SLIDE***

So what did want to get out of this effort

The P90 latencies in postgres for these calls were regularly around 100ms, but under load would could spike to 250ms, and in worst case fully time out at 3s

We needed a solution that would scale predictably as our user base continues to grow

An ideal solution would move this workload out of postgres and break the dependency on our users table

***NEXT SLIDE***

And here’s what we did.

Using a prefix trie we can achieve insert and lookup times of O(n) and a space complexity of O(n\*k) where k is the average length of the usernames that we are storing

With this structure each node in the tree represents a “prefix” that can be searched. Each node would store the set of user ids for users whose names match the prefix.

This means that given a prefix to search you can directly reference the full set of user ids which are a match.

And finally, we’ll put all this data in Redis since we need somewhere to store all these prefix keys

***NEXT SLIDE***

So how do we actually do this. At a high level these are the steps

We Compute each possible prefix for both the first and last names of a user to use as keys in Redis

Then store a hash set which contains the users first and last names and a key to sort matches on, as well as maintain a set of users IDs on a per group basis which we can use for filtering by group membership

To actually find a match we just need to pass in our prefix and we will return a result a set of ids sorted by the user name

To save on repeat searches we store the sorted ID set for a short amount of time and can reuse it if the same prefix is passed. This helps in cases where a user may make a typo and backspace their way through previously matched sets

***NEXT SLIDE***

And now we’re going to look at some code\!

This how we break apart the users name into it’s full set of prefixes. Because we support completions by both first and last names we need to create a set of keys for each.

When we actually call Redis to insert the keys we’re running them as a Lua script to cut down on network overhead so we don’t need to make a call to insert each prefix individually.

***NEXT SLIDE***

As we computed each prefix we sanitized the strings to remove diacritical marks and normalize them such that we’re left only with lowercase letters and numbers.

***NEXT SLIDE***

For the matching we build a set of keys out of the prefix we want to match on, provide a max number of matches we’re interested in, and fire off the redis command. Like the inserts we do this all in a Lua script on the redis side to cut down that network overhead.

***NEXT SLIDE***

And here’s what we found when we actually got this deployed.

It’s fast, it’s scales, and it’s isolated from our most critical db tables\!

One of the biggest things we noticed other than lower latency in general, is that the variance in that latency is much smaller than it was when we were serving this out of postgres. The biggest thing here is that we’re not having to compete with locking and other things happening in our postgres database.

***NEXT SLIDE***

That said, there are some gotchas

Because we store user details in both the DB and in this redis instance extra care needs to be taken to keep them in sync

The memory usage of this implementation will grow over time with user growth (moreso with longer names)

To support additional types of filtering beyond the group filter we’ve implemented today we would need to add additional keyspaces to redis (and in turn add to the memory pressure)

***NEXT SLIDE***

And here are some of our key takeaways from this effort.

At the end of the day, if you need to look something up based on a prefix a prefix trie is likely to be a pretty good fit

Carving this typeahead completion out of Postgres and putting it into Redis alleviated load on our database and made this typeahead feature faster.

One of the big tradeoffs we found in this approach is that your memory needs will scale with growth. We can continue to scale redis, but at some point we may cross a threshold where we will want to revisit

***NEXT SLIDE***

And that’s all I've got for you\! If anyone has had any questions feel free send over an email come find me at our booth during the next break. 

Thanks everyone\!  

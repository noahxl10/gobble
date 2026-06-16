# Gobble

No AI has touched this project except for the test suite. This could be a bad thing, or it could be a good thing depending on how you look at it. For me, it's a good thing, as this is a severely helpful learning excercise. Maybe Mythos can rebuild it for me later...

A light framework for treating S3 like a db

(why the name gobble? well, go + object = gobject, so naturally, gobble is the more fun iteration)

Object storage:
- Pricing varies depending on availability needs, but roughly:
- $.02 - $.11per GB per month for storage depending on level of retrieval
- $.0125 for infrequent
- First 100GB out is FREE!
- cheap ($0 in, $0.05-$0.09 per GB out after 100GB)
- Retrieval requests are as low as $0.0004 per 1000 requests!

Downsides:
- speed (cold reads hurt!)
- setup

Why not block storage?
- Limited scale
- Little metadata
- Throughput is much more expensive
Why not file storage?
- Performance is worse, as fs is designed for human-readable structures.


Core features:
- WAL
- Auto storage for file types (so far): CSV
- Indexing
- Backups
- Buffer cache
- Checkpoints
- Cold/Hot storage preferences
- Auto-budgeting? COMING SOON


Potential future improvements:
- RBAC?
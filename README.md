# Hybrid cache library written
- In memory
- Distributed cache with Redis
- Auto sync with RedisBus
- Geo, List, Set
- Diagnostics
- Stats
- System clock
- Node init
- Auto expiration
- Benchmarks
- Cache replacements LRU, LFU, FIFO, MRU, Clock
- Sharding (djb2, djb2a, murmur, sdbm, fnv)
- Consistant hash
- Peers
- Generics
- Sinks
- Response cache
- Json
- Binary
- Protobuf
- Mutex lock
- Sleep
- Expiration scan frequency
- Size limit
- Logging
- Middlewares
- Compressor
- Absolute expiration
- Sliding expiration
- Distributed lock
- High concurrency (buckets)
- murmur2
- GC effiency (no pointers in map)
- large maps cause significant GC pauses
- Main cache, Hot cache

https://easycaching.readthedocs.io/en/latest/Hybrid/

Cache Operations: Support common cache operations such as GET, SET, DEL, EXPIRE, INCR, DECR, etc. These operations enable storing and retrieving data from the cache.

Data Distribution: Utilize Redis as a distributed cache to store data across multiple nodes or instances. This allows for horizontal scaling and improved performance by distributing the data across the Redis cluster.

Cache Expiration: Support setting TTL (time-to-live) for cache entries to automatically expire and evict stale data from the cache. This helps manage cache size and ensures data freshness.

Pub/Sub Messaging: Utilize Redis Pub/Sub functionality for cache synchronization across multiple instances. This enables real-time notifications and updates to be propagated to all instances when the cached data changes.

Cache Invalidation: Provide mechanisms to invalidate or update cache entries when the corresponding data changes. This can be achieved by using Redis Pub/Sub to publish cache invalidation messages when relevant data is updated.

Cache Miss Handling: Handle cache misses by fetching the data from the data source (e.g., database or external API) and populating the cache. This ensures that subsequent requests for the same data can be served from the cache.

Cache Read-through and Write-through: Support read-through and write-through caching patterns, where cache operations trigger corresponding data retrieval or updates in the data source, ensuring data consistency between the cache and the underlying data store.

Cache Serialization: Serialize and deserialize cache data to facilitate storage and retrieval. This may involve converting data into a suitable format such as JSON, binary, or protocol buffers.

Cache Compression: Optionally support data compression techniques to minimize memory usage and storage requirements in the cache.

Cache Monitoring: Provide monitoring and metrics related to cache usage, hit rates, eviction rates, memory consumption, and other cache performance indicators. This helps track cache efficiency and identify potential bottlenecks.

Cache Configuration: Allow configuration of cache parameters such as cache size, eviction policies, Redis connection settings, Pub/Sub channels, and other relevant configuration options.

Cache Partitioning and Sharding: Support partitioning or sharding of cache data across multiple Redis instances to distribute the data and workload effectively, ensuring high availability and performance.

Cache Consistency: Handle cache consistency issues, such as race conditions or concurrent updates, to maintain data integrity across distributed cache instances.

Error Handling and Retry: Implement appropriate error handling and retry mechanisms for cache operations and Redis connection failures, ensuring reliable cache access and availability.

Security and Authentication: Provide secure communication with Redis by supporting authentication and encryption mechanisms to protect cache data and ensure secure connectivity.

Integration with Other Systems: Allow integration with other systems or frameworks, such as popular caching libraries or frameworks in the chosen programming language, to provide seamless interoperability and ease of use.

Documentation and Examples: Provide comprehensive documentation, usage examples, and code samples to facilitate easy adoption and integration of the hybrid cache library in various applications and environments.

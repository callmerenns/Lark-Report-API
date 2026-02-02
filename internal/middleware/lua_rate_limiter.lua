-- KEYS[1] = rate limit key
-- ARGV[1] = now (unix nano)
-- ARGV[2] = window start (unix nano)
-- ARGV[3] = limit
-- ARGV[4] = ttl seconds

redis.call('ZREMRANGEBYSCORE', KEYS[1], 0, ARGV[2])
redis.call('ZADD', KEYS[1], ARGV[1], ARGV[1])

local count = redis.call('ZCARD', KEYS[1])

if count == 1 then
	redis.call('EXPIRE', KEYS[1], ARGV[4])
end

if count > tonumber(ARGV[3]) then
	return {0, count}
end

return {1, count}

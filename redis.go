package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/elliotchance/phpserialize"
	"os"
)

//缓存
var Cache redis.Conn

func main()  {
	var err error
	//--------------------缓存初始化开始
	Cache, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {

	}
	//-------------------缓存初始化结束
	reg := CacheSet("name", "小明","100")
	println(reg)
	os.Exit(1)
	/*

	//打印变量类型
	fmt.Println("type:", reflect.TypeOf(CONFIG))
	os.Exit(1)

	//添加字符串值
	reg := CacheSet("name", "小明","100")
	println(reg)
	os.Exit(1)

	//查询字符串值
	key := CacheGet("name")
	println(key)
	os.Exit(1)

	//关闭链接
	defer Cache.Close()
	*/
}



/*
 * 设置添加数组缓存
 */
func CacheSetArray(key string, val map[string]string , past string) (bool) {

	out, err := phpserialize.Marshal(val, nil)
	if err != nil {
		return false
	}

	//缓存前缀
	CACHE_PREFIX := "CACHE_PREFIX"

	_, err = Cache.Do("SET", CACHE_PREFIX + ":" + key, string(out),"EX",past)

	if err != nil {
		return false
	}

	return true

}



/*
 * 设置添加字符串缓存
 */
func CacheSet(key, val,past string) (bool) {

	out, err := phpserialize.Marshal(val, nil)
	if err != nil {
		return false
	}

	//缓存前缀
	CACHE_PREFIX := "CACHE_PREFIX"

	_, err = Cache.Do("SET", CACHE_PREFIX + ":" + key, string(out),"EX",past)

	if err != nil {
		return false
	}

	return true

}


/*
 * 查询查询数组缓存
 */
func CacheGetArray(key string) (map[string]string) {
	valNew := make(map[string]string)
	//缓存前缀
	CACHE_PREFIX := "CACHE_PREFIX"

	val, err := redis.String(Cache.Do("GET", CACHE_PREFIX + ":" + key))

	if err != nil {
		valNew["null"] = "null"
	} else {
		//反序列化
		valInterFace, err := phpserialize.UnmarshalAssociativeArray([]byte(val))
		if err != nil {
			valNew["null"] = "null"
		} else {
			//转map[string]string
			for k, v := range valInterFace {
				valNew[k.(string)] = v.(string)
			}
			valNew["null"] = "false"
		}
	}
	return map[string]string(valNew)

}



/*
 * 查询查询字符串缓存
 */
func CacheGet(key string) (string) {

	//缓存前缀
	CACHE_PREFIX := "CACHE_PREFIX"

	val, err := redis.String(Cache.Do("GET", CACHE_PREFIX + ":" + key))
	var valNew string

	if err != nil {
		valNew = "null"
	} else {
		//反序列化
		phpserialize.Unmarshal([]byte(val), &valNew)
	}
	return valNew

}


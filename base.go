package cachelib

import "errors"

// DEFAULTEXPIRE 缺省时间
var DEFAULTEXPIRE = 300

// BaseCache Cache Interface
type BaseCache interface{
	// Get Look up key in the cache and return the value for it
	Get(key string) (result string)
	// CheckGet just like `Get`,but append an error flag
	CheckGet(key string)(result string,err error)
	// GetMany Returns a list of values for the given keys
	GetMany(keys... string)(result []string)
	// CheckGetMany Just like GetMany,but append an error flag
	CheckGetMany(keys... string)(result []string,err error)
	// Set add a new key/value to the cache (overwrites value, 
	//if key already exists in the cache)
	Set(key string, value string) (result string)
	// CheckSet Just like `Set` ,but append an error flag
	CheckSet(key string, value string) (result string,err error)
	// SetMany Sets multiple keys and values from a mapping
	SetMany(keyMap map[string]string) (result []string)
	// CheckSetMany Just like SetMany,but append an error flag
	CheckSetMany(keyMap map[string]string) (result []string,err error)
	// Add Works like `Set` but does not overwrite the values of 
	// already existing keys
	Add(key string, value string) bool
	Delete(key string)bool
	// Has Checks if a key exists in the cache without returning it. This is a
	// cheap operation that bypasses loading the actual data on the backend.
	Has(key string)bool
	// Clear Clears the cache
	Clear()bool
	Inc(key string,delta uint)string
	Dec(key string,delta uint)string
	
}

// CacheBase Cache Base config
type CacheBase struct{
	KeyPrefix string
	Expire uint
}


// NullCache cache that doesn't cache.  This can be useful for unit testing
type NullCache struct{
	CacheBase
}

// Get X
func (c *NullCache) Get(key string)string{
	return ""
}

// CheckGet x
func (c *NullCache) CheckGet(key string)(string,error){
	return "", errors.New("NULL")
}

// GetMany x
func (c *NullCache)GetMany(keys... string)(result []string){
	return []string{}
}
// CheckGetMany x
func (c *NullCache)CheckGetMany(keys... string)(result []string,err error){
	return []string{}, errors.New("NULL")
}

// Set x
func (c *NullCache)Set(key string, value string) (result string){
	return ""
}

// CheckSet x
func (c *NullCache)CheckSet(key string, value string) (result string,err error){
	return "",errors.New("NULL")
}

// SetMany x
func (c *NullCache)SetMany(keyMap map[string]string) (result []string){
	return []string{}
}

// CheckSetMany x
func (c *NullCache)CheckSetMany(keyMap map[string]string) (result []string,err error){
	return []string{},errors.New("NULL")
}

// Add x
func (c *NullCache)Add(key string, value string) bool{
	return false
}

// Delete x
func (c *NullCache)Delete(key string)bool{
	return true
}
// Has x
func (c *NullCache)Has(key string)bool{
	return false
}

// Clear Clears the cache
func (c *NullCache)Clear()bool{
	return true
}

// Inc x
func (c *NullCache)Inc(key string,delta uint)string{
	return ""
}

// Dec x
func (c *NullCache)Dec(key string,delta uint)string{
	return ""
}

// NewNullCache Create New NullCache Instance
func NewNullCache() *NullCache{
	return &NullCache{}
}
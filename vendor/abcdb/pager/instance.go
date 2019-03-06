package pager

//Singleton is something
var Singleton *StructPager

// Instance : create / get the global singleton instance of pager
// see [Singleton Pattern](https://en.wikipedia.org/wiki/Singleton_pattern)
func Instance() *StructPager {
	if Singleton == nil {
		Singleton = new(StructPager)
	}
	return Singleton
}

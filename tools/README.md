Go App Engine SDK Fix
=====================

`go_appengine` was downloaded from
https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_darwin_amd64-1.9.38.zip

`go_appengine_vendor_symlink_fix` was created by downloaded the same SDK as
`go_appengine` and applying `go_appengine_vendor_symlink_fix.patch`

The change adds a new property to devappserver2.ModuleConfiguration containing
the application root without removing symlinks, and uses this new property when
launching `go-app-builder`.

There are probably other issues caused by this same root problem that this
patch does not resolve, but I haven't encountered them.

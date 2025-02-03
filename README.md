# tag-sandbox

This is sample project to demonstrate behaviour of containerimage.ParseReference() function. 

output:

```
./play.ground 
artifactTag = 
artifactDigest = sha256:1420cefd4b20014b3361951c22593de6e9a2476bbbadd1759464eab5bfc0d34f
artifactRepository = my-app
artifactRepository = my-private-repo.company.com
```

As you can see - last version of containerimage (v0.20.3) doesn't return tag.

Last version dated at 15 Jan 2025:
https://github.com/google/go-containerregistry/commits/v0.20.3/

Supposed fix dated 5 Mar 2019. (Far time ago ..):
google/go-containerregistry#39
https://github.com/google/go-containerregistry/pull/391/commits

So it must include this commit.
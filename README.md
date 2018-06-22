If you want to include some files in a package, use this stanza:

    //go:generate include README.md

Now you can include the file like this:

    import "<your path>/internal/assets/README_md"

And you can use the file like this:

    README_md.Bytes()

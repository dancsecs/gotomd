<!---             *****  AUTO GENERATED:  DO NOT MODIFY  ***** -->
<!---          MODIFY TEMPLATE: './testdata/tstpkg/.README.gtm.md' -->
<!---               See: 'https://github.com/dancsecs/gotomd' -->

# Package example1

This project is used by the Szerszam utility function to test its markdown
update methods against an independent standalone project. All features
will be tested against this file so it will be updated and changed often.

The following will be replaced by the go package documentation

```go
package example1
```

Package example1 demonstrates various template options.

# MarkDown Headings can be used in go docs.

# Markdown code formatting may be used in go doc templates.

It will be translated to go doc format (tabbed) when processed.

    #!/bin/bash
    echo "Hello, world."

# Include (and expand) Shared Snippet From .doc.gtm.go Template
# Common Snippet Inclusion

    #!/bin/bash
    echo "Hello, world."

# Include (and expand) Shared Snippet From .README.gtm.md Template
# Common Snippet Inclusion

```bash
#!/bin/bash
echo "Hello, world."
```

# Embedded in .README.gtm.md template

```bash
#!/bin/bash
echo "Hello, world."
```

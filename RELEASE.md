Release procedure
=================

How to release a new version of sleepd.

## Versioning

The version number of sleepd follows [semantic versioning 2.0.0][semver].

## Bump version

1. Determine a new version number. Set the `$VERSION` variable as follows.
    ```console
    $ VERSION=X.Y.Z
    ```

2. Checkout `main` branch.

3. Change the `VERSION` variable in `Makefile`. The value should have `-devel` suffix.

    ```console
    $ sed -i "1cVERSION ?= $VERSION-devel" Makefile
    ```

4. Commit the changes and push it.

    ```console
    $ git commit -a -m "Bump version to $VERSION"
    $ git push
    ```

5. Add a git tag, then push it.

    ```console
    $ git tag "v$VERSION"
    $ git push origin "v$VERSION"
    ```

    When the release CI is successful, the latest container image is uploaded to Docker Hub.
    And the new GitHub release page is created.

6. Edit [the GitHub release page](https://github.com/masa213f/sleepd/releases/latest).

    Add notable changes and make the page publish.

[semver]: https://semver.org/spec/v2.0.0.html

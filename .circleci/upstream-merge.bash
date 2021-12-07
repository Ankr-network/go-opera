UPSTREAM_REPO="Fantom-foundation/go-opera"

UPSTREAM_URL="https://github.com/$UPSTREAM_REPO"
git remote get-url upstream &> /dev/null && git remote remove upstream
git remote add upstream $UPSTREAM_URL
git fetch --tags upstream
LATEST_RELEASE_TAG=`curl --silent https://api.github.com/repos/$UPSTREAM_REPO/releases/latest | jq --raw-output .target_commitish`
git merge upstream/$LATEST_RELEASE_TAG

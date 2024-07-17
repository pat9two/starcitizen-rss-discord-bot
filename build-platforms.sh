echo "Setting linux OS and ARCH"
. set-linux-build.sh
echo "Building..."
go build
echo "Done!"
echo "Setting windows OS and ARCH"
. set-windows-build.sh
echo "Building..."
go build
echo "Done!"
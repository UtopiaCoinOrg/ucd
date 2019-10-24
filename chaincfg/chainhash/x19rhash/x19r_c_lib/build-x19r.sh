#!/bin/sh

ARCHS="arm64 armv7s x86_64 i386 armv7"
CWD=`pwd`
SCRATCH="scratch-x19r"

rm -rf  "$SCRATCH/$ARCH"
for ARCH in $ARCHS
	do
		echo "building $ARCH..."
		mkdir -p "$SCRATCH/$ARCH"
		if [ "$ARCH" = "i386" -o "$ARCH" = "x86_64" ]
		then
		    PLATFORM="iPhoneSimulator"
		    if [ "$ARCH" = "x86_64" ]
		    then
		    	SIMULATOR="-mios-simulator-version-min=7.0"
                        HOST=x86_64-apple-darwin
		    else
		    	SIMULATOR="-mios-simulator-version-min=5.0"
                        HOST=i386-apple-darwin
		    fi
		else
		    PLATFORM="iPhoneOS"
		    SIMULATOR=
                    HOST=arm-apple-darwin
		fi
		
		XCRUN_SDK=`echo $PLATFORM | tr '[:upper:]' '[:lower:]'`
		CC="xcrun -sdk $XCRUN_SDK clang -arch $ARCH"
		CFLAGS="-arch $ARCH $SIMULATOR"
		if ! xcodebuild -version | grep "Xcode [1-6]\."
		then
			CFLAGS="$CFLAGS -fembed-bitcode"
		fi
		CXXFLAGS="$CFLAGS"
		LDFLAGS="$CFLAGS"
		
		CC="$CC" CFLAGS="$CFLAGS" LDFLAGS="$LDFLAGS"
		export SCC="$CC"
		make
		cp -rf *.o "$SCRATCH/$ARCH/"
		cp -rf *.a "$SCRATCH/$ARCH/"
		rm -rf *.o
		rm -rf *.a
	done


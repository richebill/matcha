package pb

//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/app/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/env/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/keyboard/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/layout/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/paint/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/text/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/touch/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/alert/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/android/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/ios/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/button/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/imageview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/progressview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/scrollview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/segmentview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/switchview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. gomatcha.io/matcha/pb/view/textinput/*.proto )"

//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/app/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/env/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/keyboard/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/layout/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/paint/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/text/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/touch/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/alert/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/android/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/ios/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/button/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/imageview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/progressview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/scrollview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/segmentview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/switchview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --java_out=gomatcha.io/matcha/android/Matcha/app/src/main/java gomatcha.io/matcha/pb/view/textinput/*.proto )"

//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/app/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/env/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/keyboard/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/layout/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/paint/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/text/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/touch/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/alert/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/android/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/ios/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/button/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/imageview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/progressview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/scrollview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/segmentview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/switchview/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=gomatcha.io/matcha/ios/Matcha/Protobuf gomatcha.io/matcha/pb/view/textinput/*.proto )"

// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/proto/image.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class MatchaPBImage;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - MatchaPBImageRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface MatchaPBImageRoot : GPBRootObject
@end

#pragma mark - MatchaPBImage

typedef GPB_ENUM(MatchaPBImage_FieldNumber) {
  MatchaPBImage_FieldNumber_Width = 1,
  MatchaPBImage_FieldNumber_Height = 2,
  MatchaPBImage_FieldNumber_Data_p = 3,
  MatchaPBImage_FieldNumber_Stride = 4,
};

@interface MatchaPBImage : GPBMessage

@property(nonatomic, readwrite) int64_t width;

@property(nonatomic, readwrite) int64_t height;

@property(nonatomic, readwrite) int64_t stride;

@property(nonatomic, readwrite, copy, null_resettable) NSData *data_p;

@end

#pragma mark - MatchaPBImageProperties

typedef GPB_ENUM(MatchaPBImageProperties_FieldNumber) {
  MatchaPBImageProperties_FieldNumber_Width = 1,
  MatchaPBImageProperties_FieldNumber_Height = 2,
  MatchaPBImageProperties_FieldNumber_Scale = 3,
};

@interface MatchaPBImageProperties : GPBMessage

@property(nonatomic, readwrite) int64_t width;

@property(nonatomic, readwrite) int64_t height;

@property(nonatomic, readwrite) double scale;

@end

#pragma mark - MatchaPBImageOrResource

typedef GPB_ENUM(MatchaPBImageOrResource_FieldNumber) {
  MatchaPBImageOrResource_FieldNumber_Image = 1,
  MatchaPBImageOrResource_FieldNumber_Path = 2,
};

@interface MatchaPBImageOrResource : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) MatchaPBImage *image;
/** Test to see if @c image has been set. */
@property(nonatomic, readwrite) BOOL hasImage;

@property(nonatomic, readwrite, copy, null_resettable) NSString *path;

@end

#pragma mark - MatchaPBColor

typedef GPB_ENUM(MatchaPBColor_FieldNumber) {
  MatchaPBColor_FieldNumber_Red = 1,
  MatchaPBColor_FieldNumber_Blue = 2,
  MatchaPBColor_FieldNumber_Green = 3,
  MatchaPBColor_FieldNumber_Alpha = 4,
};

@interface MatchaPBColor : GPBMessage

@property(nonatomic, readwrite) uint32_t red;

@property(nonatomic, readwrite) uint32_t blue;

@property(nonatomic, readwrite) uint32_t green;

@property(nonatomic, readwrite) uint32_t alpha;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
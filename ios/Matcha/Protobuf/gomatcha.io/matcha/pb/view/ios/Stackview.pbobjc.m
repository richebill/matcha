// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/ios/stackview.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

 #import "gomatcha.io/matcha/pb/view/ios/Stackview.pbobjc.h"
 #import "gomatcha.io/matcha/pb/Image.pbobjc.h"
 #import "gomatcha.io/matcha/pb/text/Text.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaiOSPBStackviewRoot

@implementation MatchaiOSPBStackviewRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - MatchaiOSPBStackviewRoot_FileDescriptor

static GPBFileDescriptor *MatchaiOSPBStackviewRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.view.ios"
                                                 objcPrefix:@"MatchaiOSPB"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MatchaiOSPBStackChildView

@implementation MatchaiOSPBStackChildView

@dynamic screenId;

typedef struct MatchaiOSPBStackChildView__storage_ {
  uint32_t _has_storage_[1];
  int64_t screenId;
} MatchaiOSPBStackChildView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "screenId",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackChildView_FieldNumber_ScreenId,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackChildView__storage_, screenId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaiOSPBStackChildView class]
                                     rootClass:[MatchaiOSPBStackviewRoot class]
                                          file:MatchaiOSPBStackviewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaiOSPBStackChildView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\003\010\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaiOSPBStackView

@implementation MatchaiOSPBStackView

@dynamic childrenArray, childrenArray_Count;
@dynamic hasTitleTextStyle, titleTextStyle;
@dynamic hasBackTextStyle, backTextStyle;
@dynamic hasBarColor, barColor;

typedef struct MatchaiOSPBStackView__storage_ {
  uint32_t _has_storage_[1];
  NSMutableArray *childrenArray;
  MatchaPBTextStyle *titleTextStyle;
  MatchaPBTextStyle *backTextStyle;
  MatchaPBColor *barColor;
} MatchaiOSPBStackView__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "childrenArray",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaiOSPBStackChildView),
        .number = MatchaiOSPBStackView_FieldNumber_ChildrenArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackView__storage_, childrenArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "titleTextStyle",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBTextStyle),
        .number = MatchaiOSPBStackView_FieldNumber_TitleTextStyle,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackView__storage_, titleTextStyle),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "backTextStyle",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBTextStyle),
        .number = MatchaiOSPBStackView_FieldNumber_BackTextStyle,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackView__storage_, backTextStyle),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "barColor",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBColor),
        .number = MatchaiOSPBStackView_FieldNumber_BarColor,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackView__storage_, barColor),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaiOSPBStackView class]
                                     rootClass:[MatchaiOSPBStackviewRoot class]
                                          file:MatchaiOSPBStackviewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaiOSPBStackView__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\003\002\016\000\003\r\000\004\010\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaiOSPBStackBar

@implementation MatchaiOSPBStackBar

@dynamic title;
@dynamic backButtonHidden;
@dynamic customBackButtonTitle;
@dynamic backButtonTitle;
@dynamic hasTitleView;
@dynamic rightViewCount;
@dynamic leftViewCount;

typedef struct MatchaiOSPBStackBar__storage_ {
  uint32_t _has_storage_[1];
  NSString *title;
  NSString *backButtonTitle;
  int64_t rightViewCount;
  int64_t leftViewCount;
} MatchaiOSPBStackBar__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "title",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackBar_FieldNumber_Title,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackBar__storage_, title),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "customBackButtonTitle",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackBar_FieldNumber_CustomBackButtonTitle,
        .hasIndex = 3,
        .offset = 4,  // Stored in _has_storage_ to save space.
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeBool,
      },
      {
        .name = "backButtonTitle",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackBar_FieldNumber_BackButtonTitle,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackBar__storage_, backButtonTitle),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "hasTitleView",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackBar_FieldNumber_HasTitleView,
        .hasIndex = 6,
        .offset = 7,  // Stored in _has_storage_ to save space.
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeBool,
      },
      {
        .name = "rightViewCount",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackBar_FieldNumber_RightViewCount,
        .hasIndex = 8,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackBar__storage_, rightViewCount),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "leftViewCount",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackBar_FieldNumber_LeftViewCount,
        .hasIndex = 9,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackBar__storage_, leftViewCount),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "backButtonHidden",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackBar_FieldNumber_BackButtonHidden,
        .hasIndex = 1,
        .offset = 2,  // Stored in _has_storage_ to save space.
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeBool,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaiOSPBStackBar class]
                                     rootClass:[MatchaiOSPBStackviewRoot class]
                                          file:MatchaiOSPBStackviewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaiOSPBStackBar__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\006\002\025\000\003\017\000\004\014\000\005\016\000\006\r\000\007\020\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - MatchaiOSPBStackEvent

@implementation MatchaiOSPBStackEvent

@dynamic idArray, idArray_Count;

typedef struct MatchaiOSPBStackEvent__storage_ {
  uint32_t _has_storage_[1];
  GPBInt64Array *idArray;
} MatchaiOSPBStackEvent__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "idArray",
        .dataTypeSpecific.className = NULL,
        .number = MatchaiOSPBStackEvent_FieldNumber_IdArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(MatchaiOSPBStackEvent__storage_, idArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldPacked),
        .dataType = GPBDataTypeInt64,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaiOSPBStackEvent class]
                                     rootClass:[MatchaiOSPBStackviewRoot class]
                                          file:MatchaiOSPBStackviewRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaiOSPBStackEvent__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)

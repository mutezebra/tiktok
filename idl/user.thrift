namespace go api.user

include "base.thrift"

struct UserInfo {
    1: optional i64 ID (go.tag="json:\"id,omitempty\"")
    2: optional string UserName (go.tag="json:\"user_name,omitempty\"")
    3: optional string Email (go.tag="json:\"email,omitempty\"")
    4: optional i8 Gender (go.tag="json:\"gender,omitempty\"")
    5: optional string Avatar (go.tag="json:\"avatar,omitempty\"")
    6: optional i32 Fans (go.tag="json:\"fans,omitempty\"")
    7: optional i32 Follows (go.tag="json:\"follows,omitempty\"")
    8: optional i64 CreateAt (go.tag="json:\"create_at,omitempty\"")
    9: optional i64 UpdateAt (go.tag="json:\"update_at,omitempty\"")
    10: optional bool TotpStatus (go.tag="json:\"totp_status,omitempty\"")
}

struct RegisterReq {
    1: optional string UserName (go.tag="json:\"user_name,omitempty\",form:\"user_name,required\"")
    2: optional string Email    (go.tag="json:\"email,omitempty\",form:\"email\"")
    3: optional string Password (go.tag="json:\"password,omitempty\",form:\"password,required\"")
}

struct RegisterResp {
    1: optional base.Base Base (go.tag="json:\"base,omitempty\"")
}

struct LoginReq {
    1: optional string UserName (go.tag="json:\"user_name,omitempty\",form:\"user_name,required\"")
    2: optional string password (go.tag="json:\"password,omitempty\",form:\"password,required\"")
    3: optional string OTPCode (go.tag="json:\"otp_code,omitempty\"")
}

struct LoginResp {
    1: optional base.Base Base (go.tag="json:\"base,omitempty\"")
    2: optional UserInfo Data (go.tag="json:\"data\"")
}

struct InfoReq {
    1: optional i64 UID (go.tag="json:\"uid,omitempty\"")
    2: optional string Name (go.tag="json:\"name,omitempty\"")
}

struct InfoResp {
    1: optional base.Base Base (go.tag="json:\"base,omitempty\"")
    2: optional UserInfo Data (go.tag="json:\"data\"")
}

struct UploadAvatarReq {
}

struct UploadAvatarResp {
    1: optional base.Base Base (go.tag="json:\"base,omitempty\"")
}

struct TotpQrcodeReq {
}

struct TotpQrcodeData {
    1: optional string Secret (go.tag="json:\"secret,omitempty\"")
    2: optional string Qrcode (go.tag="json:\"qrcode,omitempty\"")
}

struct TotpQrcodeResp {
    1: optional base.Base Base (go.tag="json:\"base,omitempty\"")
    2: optional TotpQrcodeData Data (go.tag="json:\"data\"")
}

struct EnableTotpReq {
    1: optional string Secret (go.tag="json:\"secret,omitempty\"")
    2: optional i32 Code (go.tag="json:\"code,omitempty\"")
}

struct EnableTotpResp {
    1: optional base.Base Base (go.tag="json:\"base,omitempty\"")
}

service UserService {
    RegisterResp Register(1: RegisterReq req)
    LoginResp Login(1: LoginReq req)
    InfoResp Info(1: InfoReq req)
    UploadAvatarResp UploadAvatar(1: UploadAvatarReq req)
    TotpQrcodeResp TotpQrcode(1: TotpQrcodeReq req)
    EnableTotpResp EnableTotp(1: EnableTotpReq req)
}

package middleware

import (
	"encoding/json"
	"memo-server/model/request"
	"memo-server/utils"
	"memo-server/utils/r"

	"github.com/gin-gonic/gin"
)

func DecryptMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		isEncrypted := c.Request.Header.Get("Encrypted")
		if isEncrypted != "true" {
			c.Next()
			return
		}

		// 1.获取传参， 加密请求的参数
		var params request.Encrypted
		err := c.ShouldBind(&params)
		if err != nil {
			r.OkJsonResult(c, r.ERROR_BAD_PARAM, "", nil)
			c.Abort()
			return
		}

		// 2. 从session获取会话对于的私钥
		pem, ok := utils.GetCachedPrivateKeyPEM(c)
		if !ok {
			r.OkJsonResult(c, r.APP_GETRSA_FAILED, "", nil)
			c.Abort()
			return
		}

		//3. 用私钥解密获取rsa加密后的aes (key iv)
		rsaKeyIv, err2 := utils.RSADecryptData(pem, params.Key)
		if err2 != nil {
			r.OkJsonResult(c, r.ERROR_BAD_PARAM, "请求出错了，请重试", nil)
			c.Abort()
			return
		}

		//4.保存aes的key和iv
		var keyIv utils.AESKeyIv
		err3 := json.Unmarshal(rsaKeyIv, &keyIv)
		if err3 != nil {
			r.OkJsonResult(c, r.ERROR_BAD_PARAM, "请求参数错误", nil)
			c.Abort()
			return
		}

		//5. aes解密
		data, err4 := utils.AESDecryptData(keyIv, params.Data)
		if err4 != nil {
			r.OkJsonResult(c, r.ERROR_BAD_PARAM, "请求参数错误-2", nil)
			c.Abort()
			return
		}

		c.Set(utils.ContextDecryptedParams, data)
		c.Next()
		return
	}

}

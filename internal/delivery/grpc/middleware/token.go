// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package middleware

// TODO DELETE

//// validateUser выполняет безопасную проверку токена и возвращает данные пользователя
//func validateUser(ctx context.Context, tokenMiddleware *middleware.TokenMiddleware) (*middleware.UserClaims, error) {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return nil, errors.New("metadata not found")
//	}
//
//	// Извлечение токена из метаданных
//	tokens := md.Get("Authorization")
//	if len(tokens) == 0 {
//		return nil, errors.New("authorization token not provided")
//	}
//
//	token := tokens[len(tokens)-1] // Последний токен (если их несколько)
//	if token == "" {
//		return nil, errors.New("empty authorization token")
//	}
//
//	// Валидация токена
//	userClaims, err := tokenMiddleware.ValidateToken(ctx)
//	if err != nil {
//		return nil, err // Ошибки токена будут обработаны выше
//	}
//
//	return userClaims, nil
//}

//type UserInfo struct {
//	Username string `json:"username"`
//	Id       int    `json:"id"`
//}
//
//type UserClaims struct {
//	*jwt.StandardClaims
//	UserInfo
//}
//
//func (t *TokenMiddleware) CreateToken(user *model.User) (string, error) {
//	claims := &UserClaims{
//		&jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
//			Issuer:    "voo.su",
//		},
//		UserInfo{
//			Username: user.Username,
//			Id:       user.Id,
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(t.Conf.App.Jwt.Secret))
//}
//
//func (t *TokenMiddleware) ParseToken(tokenString string) (*UserClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf(t.Locale.Localize("unexpected_signature_method"), token.Header["alg"])
//		}
//
//		return []byte(t.Conf.App.Jwt.Secret), nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
//		return claims, nil
//	}
//
//	return nil, err
//}
//
//func (t *TokenMiddleware) ValidateToken(ctx context.Context) (*UserClaims, error) {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return nil, errors.New(t.Locale.Localize("failed_to_fetch_metadata"))
//	}
//
//	token := md.Get("Authorization")
//	userClaims, err := t.ParseToken(token[len(token)-1])
//	if err != nil {
//		return nil, err
//	}
//
//	return userClaims, nil
//}

// controller содержит структуры для обработки запросов
//
// HandlerBuilder позволяет создать сервер
//
// TokensPairGetter позволяет получить пару токенов (access и refresh)
//
// TokensRefresher позволяет обновить токены
// Если меняется User-Agent, то токен сбрасывается
//
// TokensIdGetter позволяет получить user Id
//
// TokenUnauthorized позволяет выйти через систему
package controller
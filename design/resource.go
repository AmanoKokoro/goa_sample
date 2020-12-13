package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

// /actionsの定義をする
var _ = Service("actions", func() {
	/*
	   リソースに対してどういった操作を行うか定義する
	   add リソースを追加する
	   list    リソースをリストで取得する
	   delete  リソースを削除する
	   上記のような感じで定義すればおｋです。
	*/
	// Error("BadRequest")
	// Error("NotFound")

	// HTTP(func() {
	// 	// "DivByZero" エラーに HTTP ステータスコード 400 Bad Request を使う。
	// 	Response("BadRequest", StatusBadRequest)
	// })
	// HTTP(func() {
	// 	Response("NotFound", StatusNotFound)
	// })

	Method("ping", func() {
		// アクションの説明
		Description("サーバーとの導通確認")
		// 返したいレスポンス
		// 200 OK + MessageTypeで定義しているMediaType
		// Response()
		Result(String)
		// // 400 BadRequest + ErrorMediaというデフォルトで容易されているMediaType
		// // 足りないパラメーターなどがあれば自動的に返される
		//Response(BadRequest, ErrorMedia)
		HTTP(func() {
			PUT("/ping")
		})
	})

	Method("hello", func() {
		Description("挨拶する")

		Payload(func() {
			Attribute("name", String, "Name", func() {
				Default("")
			})
			Required("name")
		})

		Result(String)

		HTTP(func() {
			GET("/hello")
		})
	})
	Method("ID", func() {
		Description("複数アクション（:id）")
		Payload(func() {
			Attribute("id", Int, "id")
		})

		Result(Int)

		HTTP(func() {
			GET("/{id}")
		})
	})
})

// Swaggerをローカルで実行するめの定義
var _ = Service("swagger", func() {
	cors.Origin("*", func() {
		cors.Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})

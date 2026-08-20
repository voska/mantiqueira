// Package mantiqueira describes the Mantiqueira em Casa storefront
// (www.mantiqueiraemcasa.com.br), the direct-to-consumer egg brand of Grupo
// Mantiqueira.
package mantiqueira

import "github.com/voska/vtexkit/store"

// Store is the Mantiqueira em Casa descriptor.
//
// Account is the one field that cannot be derived. Every other store so far
// has an account name matching its domain, so store.AccountName falls back to
// the host — which would yield "mantiqueiraemcasa" here. The real account is
// "grupomantiqueira" (__RUNTIME__.account on the storefront). Getting this
// wrong breaks two things quietly: the auth cookie would be sent under
// VtexIdclientAutCookie_mantiqueiraemcasa, and the Subscriptions API would be
// addressed on the wrong account host.
//
// Nothing else needs declaring. Probed live on 2026-08-20: classic password
// and emailed access code are both enabled, Google is the only OAuth provider
// (so no OAuthDriver is needed), and Intelligent Search REST and the catalog
// API both return results, so Search stays at the default SearchAuto.
//
// No MinOrder and no Quirks: neither has any evidence behind it at this store.
//
// No Wishlist hashes. vtex.wish-list persisted queries have to be captured
// per store from a live browser session; the zero value makes `fav` report
// that this store's wishlist is unreachable rather than failing oddly.
var Store = store.Store{
	Name:        "mantiqueira",
	DisplayName: "Mantiqueira em Casa",
	BaseURL:     "https://www.mantiqueiraemcasa.com.br",
	Account:     "grupomantiqueira",
}

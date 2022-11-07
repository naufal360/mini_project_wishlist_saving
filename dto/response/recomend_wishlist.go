package response

type RecommendWishlist struct {
	WishlistId        string
	Name              string
	Insufficient      int
	CountRecommend    int
	ResponseRecommend string
}

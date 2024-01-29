package quotes

import (
	"math/rand"
)

var container = []string{
	"Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.",
	"The wise man learns more from his enemies than the fool does from his friends.",
	"The fool doth think he is wise, but the wise man knows himself to be a fool.",
	"The more I learn, the more I realize how much I don’t know.",
	"Honesty is the first chapter in the book of wisdom.",
	"The art of being wise is the art of knowing what to overlook.",
	"Never, no, never did nature say one thing and wisdom say another.",
	"The wise understand by themselves; fools follow the reports of others.",
	"The greatest wisdom is in simplicity.",
	"A wise man can see more from the bottom of a well than a fool can from a mountain top.",
	"The wise man doesn’t give the right answers; he poses the right questions.",
	"A word to the wise is enough.",
	"The doorstep to the temple of wisdom is a knowledge of our own ignorance.",
	"Patience is the companion of wisdom.",
	"The wise man hath his thoughts in his head; the fool, on his tongue.",
	"It is said that wisdom lies not in seeing things, but seeing through things.",
}

func RandomQuote() string {
	return container[rand.Intn(len(container))]
}

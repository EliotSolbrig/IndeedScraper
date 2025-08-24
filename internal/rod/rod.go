package rod

import (
	"fmt"

	"github.com/go-rod/rod"
	// "github.com/go-rod/rod/lib/launcher"
)

func GetBrowser() *rod.Browser {
	fmt.Println("GetBrowser()")
	return rod.New().MustConnect()
}

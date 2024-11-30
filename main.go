package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chromedp/chromedp"
)

func main() {
	// コマンドライン引数の定義
	htmlFile := flag.String("html", "", "Path to the HTML file to render")
	outputFile := flag.String("output", "output.png", "Path to the output PNG file")
	flag.Parse()

	// 引数の検証
	if *htmlFile == "" {
		fmt.Println("Error: HTML file path is required")
		flag.Usage()
		os.Exit(1)
	}

	// HTMLファイルの絶対パスを解決
	htmlPath, err := filepath.Abs(*htmlFile)
	if err != nil {
		fmt.Printf("Error resolving absolute path: %v\n", err)
		os.Exit(1)
	}

	// HTMLファイルをPNGにレンダリング
	if err := renderHTMLToPNG(htmlPath, *outputFile); err != nil {
		fmt.Printf("Error rendering HTML to PNG: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully rendered %s to %s\n", htmlPath, *outputFile)
}

func renderHTMLToPNG(htmlPath string, outputPath string) error {
	// Chromedpのコンテキストを作成
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 結果を格納するバッファ
	var buf []byte

	// HTMLファイルのフルパス
	htmlURL := "file://" + htmlPath

	// タスクを実行
	err := chromedp.Run(ctx,
		chromedp.Navigate(htmlURL),
		chromedp.FullScreenshot(&buf, 90),
	)
	if err != nil {
		return err
	}

	// ファイルに書き込み
	if err := os.WriteFile(outputPath, buf, 0644); err != nil {
		return err
	}

	return nil
}

package stringx

import (
	"fmt"
	"testing"
)

func TestRemoveStylesOfHtmlTag(t *testing.T) {
	const sample_html  = `
 <!-- wp:paragraph -->
 <p><s>这是</s><strong class="tss-bold">一段</strong><em class="tss-italic">测试</em><span style="text-decoration: underline" class="tss-underline">文本</span></p>
 <!-- /wp:paragraph -->
 
 <!-- wp:paragraph -->
 <p><span style="background-color:#0059FF;font-style:#0059F;" class="tss-background-color"><span style="font-weight:#E60000" class="tss-color">这是一段测试文本</span></span></p>
 <!-- /wp:paragraph -->`

	fmt.Println(RemoveStylesOfHtmlTag(sample_html,  "font-weight", "font-style"))
}
// @author      Liu Yongshuai<liuyongshuai@hotmail.com>
// @date        2019-03-02 15:26

package terminalutils

import (
	"fmt"
	"testing"
)

func TestRuneWrap(t *testing.T) {
	str := "擘画强军蓝图，指引奋进征程。2013年到2018年全国两会期间，中共中央总书记、国家主席、中央军委主席习近平连续出席解放军和武警部队代表团全体会议并发表重要讲话，提出一系列新思想、新论断、新要求。6年来，全军部队认真贯彻习主席重要讲话精神，牢固确立习近平强军思想的指导地位，重振政治纲纪，重塑组织形态，重整斗争格局，重构建设布局，重树作风形象，在中国特色强军之路上迈出坚定步伐。"
	ret, lineNum := RuneWrap(str, ScreenWidth)
	fmt.Println(ret, ScreenWidth, lineNum)
	str = "擘画强军蓝图，\n指引奋进征程。"
	ret, lineNum = RuneWrap(str, 10)
	fmt.Println(ret, ScreenWidth, lineNum)
	str = ""
	ret, lineNum = RuneWrap(str, ScreenWidth)
	fmt.Println(ret, ScreenWidth, lineNum)
}

func TestPrintTextDiff(t *testing.T) {
	text1 := `
45454545454545454  特朗普谈美国向移民发射催泪弹：移民很粗暴
sadfadsad 安徽：助力民企发展壮大 支持民营企业在行动
xcvxcvxc 特朗普喊话中美洲移民:如有必要 将永久关闭边境
sss 湖南隆回暂缓"百元车位" 中标单位曾被指可获暴利
】=-、意媒谈D&G风波：中国人记性差 抵制不了多久
*&）……（&暖新闻 带脑瘫儿子跑马拉松 父亲:让儿子少留遗憾
`
	text2 := `
、1222！@￥洞察"号登陆火星传首张照片:可见火星地平线
）（&**&……&……￥%￥##！日本茨城县发生5级地震多县有震感 尚未引发海啸
女子被顶替上学?堂姐夫:她考前已去卖猪肉 没考试
礌lklasdjgfakldgja5岁儿童简历长15页 人民日报:拔苗种不出好"庄稼"
2135457950875607网红自称回深山卖土蜂蜜 所留地址村委会:无此人
||||||||||\、、、、、、暴风雪袭击美国芝加哥地区 近900个航班被取消
`
	PrintTextDiff(text1, text2)
}

func TestPrintDiffTextByGroup(t *testing.T) {

	leftText := [][]string{
		{
			"45454545454545454  特朗普谈美国向移民发射催泪弹：移民很粗暴",
			"sadfadsad 安徽：助力民企发展壮大 支持民营企业在行动",
		},
		{
			"xcvxcvxc 特朗普喊话中美洲移民:如有必要 将永久关闭边境",
		},
		{
			"45454545454545454  特朗普谈美国向移民发射催泪弹：移民很粗暴",
			"sadfadsad 安徽：助力民企发展壮大 支持民营企业在行动",
		},
		{
			"xcvxcvxc 特朗普喊话中美洲移民:如有必要 将永久关闭边境",
		},
	}
	rightText := [][]string{
		{
			"、1222！@￥洞察号登陆火星传首张照片:可见火星地平线",
		},
		{
			"xcvxcvxc 特朗普喊话中美洲移民:如有必要 将永久关闭边境",
		},
		{
			"）（&**&……&……￥%￥##！日本茨城县发生5级地震多县有震感 尚未引发海啸",
			"女子被顶替上学?堂姐夫:她考前已去卖猪肉 没考试",
		},
		{
			"、1222！@￥洞察号登陆火星传首张照片:可见火星地平线",
			"）（&**&……&……￥%￥##！日本茨城县发生5级地震多县有震感 尚未引发海啸",
			"女子被顶替上学?堂姐夫:她考前已去卖猪肉 没考试",
			"礌lklasdjgfakldgja5岁儿童简历长15页 人民日报:拔苗种不出好庄稼",
			"2135457950875607网红自称回深山卖土蜂蜜 所留地址村委会:无此人",
			"||||||||||、、、、、、暴风雪袭击美国芝加哥地区 近900个航班被取消",
		},
	}
	PrintTextDiffByGroup(leftText, rightText)
	PrintTextDiffByGroup(leftText, [][]string{})
	PrintTextDiffByGroup([][]string{}, rightText)
}

func TestWrap(t *testing.T) {
	str := "暴风雪袭击美国芝加哥地区"
	fmt.Println(RuneWrap(str, 7))
	/**
	暴风雪
	袭击美
	国芝加
	哥地区
	*/
}

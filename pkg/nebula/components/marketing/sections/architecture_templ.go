// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package sections

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import models "github.com/onsonr/sonr/internal/orm/marketing"

// ╭───────────────────────────────────────────────────────────╮
// │                         Data Model                        │
// ╰───────────────────────────────────────────────────────────╯

// architecture is the (4th) home page architecture section
var arch = &models.Architecture{
	Heading:  "Onchain Security with Offchain Privacy",
	Subtitle: "Whenever you are ready, just hit publish to turn your site sketches into an actual designs. No creating, no skills, no reshaping.",
	Primary: &models.Technology{
		Title: "Decentralized Identity",
		Desc:  "Sonr leverages the latest specifications from W3C, DIF, and ICF to essentially have an Interchain-Connected, Smart Account System - seamlessly authenticated with PassKeys.",
		Image: &models.Image{
			Src:    HardwareWallet.Src(),
			Width:  "721",
			Height: "280",
		},
	},
	Secondary: &models.Technology{
		Title: "IPFS Vaults",
		Desc:  "Completely distributed, encrypted, and decentralized storage for your data.",
		Image: &models.Image{
			Src:    DecentralizedNetwork.Src(),
			Width:  "342",
			Height: "280",
		},
	},
	Tertiary: &models.Technology{
		Title: "Service Records",
		Desc:  "On-chain validated services created by Developers for secure transmission of user data.",
		Image: &models.Image{
			Src:    DefiDashboard.Src(),
			Width:  "342",
			Height: "280",
		},
	},
	Quaternary: &models.Technology{
		Title: "First-Class Exchange",
		Desc:  "Sonr integrates with the IBC protocol allowing for seamless integration with popular exchanges such as OKX, Binance, and Osmosis.",
		Image: &models.Image{
			Src:    PortfolioBalance.Src(),
			Width:  "342",
			Height: "280",
		},
	},
	Quinary: &models.Technology{
		Title: "Authentication & Authorization",
		Desc:  "Sonr leverages decentralized Macaroons and Multi-Party Computation to provide a secure and decentralized authentication and authorization system.",
		Image: &models.Image{
			Src:    DecentralizedWebNode.Src(),
			Width:  "342",
			Height: "280",
		},
	},
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Render Section View                      │
// ╰───────────────────────────────────────────────────────────╯

// Architecture is the (4th) home page architecture section
func Architecture() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Features #2 --><section><div class=\"py-12 md:py-20\"><div class=\"max-w-5xl mx-auto px-4 sm:px-6\"><div class=\"relative max-w-3xl mx-auto text-center pb-12 md:pb-20\"><h2 class=\"font-inter-tight text-3xl md:text-4xl font-bold text-zinc-900 mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Heading)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 76, Col: 20}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"text-lg text-zinc-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Subtitle)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 79, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><div class=\"max-w-xs mx-auto sm:max-w-none grid sm:grid-cols-2 md:grid-cols-3 gap-8 sm:gap-4 lg:gap-8\"><article class=\"sm:col-span-2 flex flex-col border border-transparent [background:linear-gradient(theme(colors.white),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box] rounded-lg\"><div class=\"grow flex flex-col p-5 pt-6\"><div class=\"flex items-center space-x-3 mb-1\"><svg class=\"inline-flex fill-zinc-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"M17 9c.6 0 1 .4 1 1v6a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h6c.6 0 1 .4 1 1s-.4 1-1 1H4v12h12v-6c0-.6.4-1 1-1Zm-.7-6.7c.4-.4 1-.4 1.4 0 .4.4.4 1 0 1.4l-8 8c-.2.2-.4.3-.7.3-.3 0-.5-.1-.7-.3-.4-.4-.4-1 0-1.4l8-8Z\"></path></svg><h3 class=\"font-inter-tight font-semibold text-zinc-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Primary.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 101, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3></div><p class=\"grow max-w-md text-sm text-zinc-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Primary.Desc)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 105, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><figure><img class=\"h-[280px] object-cover object-left mx-auto sm:object-contain sm:h-auto\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Primary.Image.Src)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 111, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"721\" height=\"280\" alt=\"Feature Post 01\"></figure></article><article class=\"flex flex-col border border-transparent [background:linear-gradient(theme(colors.white),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box] rounded-lg\"><div class=\"grow flex flex-col p-5 pt-6\"><div class=\"flex items-center space-x-3 mb-1\"><svg class=\"inline-flex fill-zinc-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"m6.035 17.335-4-14c-.2-.8.5-1.5 1.3-1.3l14 4c.9.3 1 1.5.1 1.9l-6.6 2.9-2.8 6.6c-.5.9-1.7.8-2-.1Zm-1.5-12.8 2.7 9.5 1.9-4.4c.1-.2.3-.4.5-.5l4.4-1.9-9.5-2.7Z\"></path></svg><h3 class=\"font-inter-tight font-semibold text-zinc-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Secondary.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 134, Col: 31}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3></div><p class=\"grow max-w-md text-sm text-zinc-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Secondary.Desc)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 138, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><figure><img class=\"h-[280px] object-cover object-left mx-auto sm:object-contain sm:h-auto\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Secondary.Image.Src)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 144, Col: 38}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"342\" height=\"280\" alt=\"Feature Post 02\"></figure></article><article class=\"flex flex-col border border-transparent [background:linear-gradient(theme(colors.white),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box] rounded-lg\"><div class=\"grow flex flex-col p-5 pt-6\"><div class=\"flex items-center space-x-3 mb-1\"><svg class=\"inline-flex fill-zinc-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"M8.974 16c-.3 0-.7-.2-.9-.5l-2.2-3.7-2.1 2.8c-.3.4-1 .5-1.4.2-.4-.3-.5-1-.2-1.4l3-4c.2-.3.5-.4.9-.4.3 0 .6.2.8.5l2 3.3 3.3-8.1c0-.4.4-.7.8-.7s.8.2.9.6l4 8c.2.5 0 1.1-.4 1.3-.5.2-1.1 0-1.3-.4l-3-6-3.2 7.9c-.2.4-.6.6-1 .6Z\"></path></svg><h3 class=\"font-inter-tight font-semibold text-zinc-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Tertiary.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 167, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3></div><p class=\"grow max-w-md text-sm text-zinc-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Tertiary.Desc)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 171, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><figure><img class=\"h-[280px] object-cover object-left mx-auto sm:object-contain sm:h-auto\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Tertiary.Image.Src)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 177, Col: 37}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"342\" height=\"280\" alt=\"Feature Post 03\"></figure></article><article class=\"flex flex-col border border-transparent [background:linear-gradient(theme(colors.white),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box] rounded-lg\"><div class=\"grow flex flex-col p-5 pt-6\"><div class=\"flex items-center space-x-3 mb-1\"><svg class=\"inline-flex fill-zinc-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"M9.3 11.7c-.4-.4-.4-1 0-1.4l7-7c.4-.4 1-.4 1.4 0 .4.4.4 1 0 1.4l-7 7c-.4.4-1 .4-1.4 0ZM9.3 17.7c-.4-.4-.4-1 0-1.4l7-7c.4-.4 1-.4 1.4 0 .4.4.4 1 0 1.4l-7 7c-.4.4-1 .4-1.4 0ZM2.3 12.7c-.4-.4-.4-1 0-1.4l7-7c.4-.4 1-.4 1.4 0 .4.4.4 1 0 1.4l-7 7c-.4.4-1 .4-1.4 0Z\"></path></svg><h3 class=\"font-inter-tight font-semibold text-zinc-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Quaternary.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 200, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3></div><p class=\"grow max-w-md text-sm text-zinc-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Quaternary.Desc)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 204, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><figure><img class=\"h-[280px] object-cover object-left mx-auto sm:object-contain sm:h-auto\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Quaternary.Image.Src)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 210, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"342\" height=\"280\" alt=\"Feature Post 04\"></figure></article><article class=\"flex flex-col border border-transparent [background:linear-gradient(theme(colors.white),theme(colors.zinc.50))_padding-box,linear-gradient(120deg,theme(colors.zinc.300),theme(colors.zinc.100),theme(colors.zinc.300))_border-box] rounded-lg\"><div class=\"grow flex flex-col p-5 pt-6\"><div class=\"flex items-center space-x-3 mb-1\"><svg class=\"inline-flex fill-zinc-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\"><path d=\"M16 2H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h8.667l3.733 2.8A1 1 0 0 0 18 17V4a2 2 0 0 0-2-2Zm0 13-2.4-1.8a1 1 0 0 0-.6-.2H4V4h12v11Z\"></path></svg><h3 class=\"font-inter-tight font-semibold text-zinc-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Quinary.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 233, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3></div><p class=\"grow max-w-md text-sm text-zinc-500\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Quinary.Desc)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 237, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div><figure><img class=\"h-[280px] object-cover object-left mx-auto sm:object-contain sm:h-auto\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Quinary.Image.Src)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/nebula/components/marketing/sections/architecture.templ`, Line: 243, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"342\" height=\"280\" alt=\"Feature Post 05\"></figure></article></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

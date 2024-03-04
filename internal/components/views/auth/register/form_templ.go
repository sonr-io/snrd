// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package register

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/sonrhq/sonr/internal/components/base"
)

func CredentialFormView() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{"flex flex-row justify-evenly items-center align-middle pb-2.5 gap-x-1.5"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><sl-input name=\"email\" placeholder=\"sjobs@apple.com\" size=\"large\" autocomplete=\"username webauthn\" required><sl-icon class=\"text-stone-200\" slot=\"prefix\" name=\"at\"></sl-icon></sl-input> <sl-button type=\"submit\" size=\"large\" variant=\"default\"><div class=\"text-md\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = base.PasskeyIcon().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></sl-button><div id=\"name-error\" aria-live=\"polite\" hidden></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func emailIdentifierForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form><div class=\"space-y-4\"><div><label class=\"block text-sm text-slate-300 font-medium mb-1\" for=\"email\">Email <span class=\"text-rose-500\">*</span></label> <input id=\"email\" class=\"form-input w-full\" type=\"email\" placeholder=\"markrossi@company.com\" required></div></div><div class=\"mt-6\"><button class=\"btn text-sm text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white w-full shadow-sm group\"><span class=\"mr-1\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"18\" height=\"17\" fill=\"none\" viewBox=\"0 0 24 24\"><path fill=\"#000\" d=\"M2.45759 5.40673C2.32495 5.61793 2.41154 5.88824 2.62359 6.0195L11.4735 11.498C11.796 11.6977 12.2037 11.6977 12.5262 11.498L21.3764 6.01941C21.5884 5.88814 21.675 5.61782 21.5423 5.40662C21.0117 4.56167 20.0714 4 19 4H5C3.92853 4 2.98825 4.56171 2.45759 5.40673Z\"></path><path fill=\"#000\" d=\"M22 8.88312C22 8.49137 21.5699 8.25179 21.2368 8.45799L13.579 13.1986C12.6114 13.7975 11.3883 13.7975 10.4208 13.1986L2.76318 8.45812C2.43008 8.25192 2 8.4915 2 8.88326V17C2 18.6569 3.34315 20 5 20H19C20.6569 20 22 18.6569 22 17V8.88312Z\"></path></svg></span> Send Confirmation Code <span class=\"tracking-normal text-slate-700 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1\">-&gt;</span></button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func createProfileForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form><div class=\"space-y-4\"><div><label class=\"block text-sm text-slate-300 font-medium mb-1\" for=\"email\">Email <span class=\"text-rose-500\">*</span></label> <input id=\"email\" class=\"form-input w-full\" type=\"username\" placeholder=\"mkiay\" required></div></div><div class=\"mt-6\"><button class=\"btn text-sm text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white w-full shadow-sm group\"><span class=\"mr-1\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"18\" height=\"17\" fill=\"none\" viewBox=\"0 0 24 24\"><path fill=\"#000\" d=\"M2.45759 5.40673C2.32495 5.61793 2.41154 5.88824 2.62359 6.0195L11.4735 11.498C11.796 11.6977 12.2037 11.6977 12.5262 11.498L21.3764 6.01941C21.5884 5.88814 21.675 5.61782 21.5423 5.40662C21.0117 4.56167 20.0714 4 19 4H5C3.92853 4 2.98825 4.56171 2.45759 5.40673Z\"></path><path fill=\"#000\" d=\"M22 8.88312C22 8.49137 21.5699 8.25179 21.2368 8.45799L13.579 13.1986C12.6114 13.7975 11.3883 13.7975 10.4208 13.1986L2.76318 8.45812C2.43008 8.25192 2 8.4915 2 8.88326V17C2 18.6569 3.34315 20 5 20H19C20.6569 20 22 18.6569 22 17V8.88312Z\"></path></svg></span> Send Confirmation Code <span class=\"tracking-normal text-slate-700 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1\">-&gt;</span></button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateCredential(rpName, rpId, challenge string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_CreateCredential_c37e`,
		Function: `function __templ_CreateCredential_c37e(rpName, rpId, challenge){const publicKeyCredentialCreationOptions = {
		challenge: Uint8Array.from(challenge, c => c.charCodeAt(0)),
		rp: {
			name: rpName,
			id: rpId,
		},
		user: {
			id: Uint8Array.from(
				"UZSL85T9AFC", c => c.charCodeAt(0)),
			name: "lee@webauthn.guide",
			displayName: "Lee",
		},
		pubKeyCredParams: [{alg: -7, type: "public-key"}],
		authenticatorSelection: {
			authenticatorAttachment: "cross-platform",
		},
		timeout: 60000,
		attestation: "direct"
	};

	const credential = navigator.credentials.create({
		publicKey: publicKeyCredentialCreationOptions
	}).then((credential) => {
		console.log(credential);
	});
}`,
		Call:       templ.SafeScript(`__templ_CreateCredential_c37e`, rpName, rpId, challenge),
		CallInline: templ.SafeScriptInline(`__templ_CreateCredential_c37e`, rpName, rpId, challenge),
	}
}

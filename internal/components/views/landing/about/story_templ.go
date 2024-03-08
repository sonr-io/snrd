// Code generated by templ - DO NOT EDIT.

<<<<<<< HEAD
// templ: version: v0.2.598
=======
// templ: version: v0.2.543
>>>>>>> master
package about

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func storySection() templ.Component {
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
<<<<<<< HEAD
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Story --><section class=\"relative\"><!-- Blurred shape --><div class=\"absolute top-0 -mt-32 left-1/2 -translate-x-1/2 ml-10 blur-2xl opacity-70 pointer-events-none -z-10\" aria-hidden=\"true\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"434\" height=\"427\"><defs><linearGradient id=\"bs4-a\" x1=\"19.609%\" x2=\"50%\" y1=\"14.544%\" y2=\"100%\"><stop offset=\"0%\" stop-color=\"#A855F7\"></stop> <stop offset=\"100%\" stop-color=\"#6366F1\" stop-opacity=\"0\"></stop></linearGradient></defs> <path fill=\"url(#bs4-a)\" fill-rule=\"evenodd\" d=\"m0 0 461 369-284 58z\" transform=\"matrix(1 0 0 -1 0 427)\"></path></svg></div><div class=\"px-4 sm:px-6\"><div class=\"max-w-5xl mx-auto\"><div class=\"pb-12 md:pb-20\"><!-- Section header --><div class=\"max-w-3xl mx-auto text-center pb-12 md:pb-20\"><h2 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60\">Our story (so far)</h2></div><div class=\"md:flex justify-between space-x-6 md:space-x-8 lg:space-x-14\"><figure class=\"min-w-[240px]\"><img class=\"sticky top-8 mx-auto mb-12 md:mb-0 rounded-lg -rotate-[4deg]\" src=\"https://cdn.sonr.build/images/team.jpg\" width=\"420\" height=\"280\" alt=\"Team\"></figure><div class=\"max-w-[548px] mx-auto\"><div class=\"text-slate-400 space-y-6\"><p>At Sonr, we're committed to reshaping the digital landscape. We believe in a decentralized web that is secure, intuitive, and user-centric, a world where both users and developers can interact with trust and ease. We aim to foster a digital world that champions privacy, enhances trust, and empowers its users through cutting-edge technology and innovation.</p><p>Sonr was born out of the realization that the decentralized web was teetering on a precipice. We witnessed a digital landscape saturated with complex systems that lacked intuitive design, impeding the wider adoption of blockchain technology.</p><p>Our vision is a future where digital freedom is not a privilege reserved for the technologically adept but a fundamental right accessible to everyone. Our aim is to demystify the decentralized web, developing solutions that are straightforward enough for anyone to use while remaining deeply secure and reliable.</p></div></div></div></div></div></div></section>")
=======
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Story --><section class=\"relative\"><!-- Blurred shape --><div class=\"absolute top-0 -mt-32 left-1/2 -translate-x-1/2 ml-10 blur-2xl opacity-70 pointer-events-none -z-10\" aria-hidden=\"true\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"434\" height=\"427\"><defs><linearGradient id=\"bs4-a\" x1=\"19.609%\" x2=\"50%\" y1=\"14.544%\" y2=\"100%\"><stop offset=\"0%\" stop-color=\"#A855F7\"></stop> <stop offset=\"100%\" stop-color=\"#6366F1\" stop-opacity=\"0\"></stop></linearGradient></defs> <path fill=\"url(#bs4-a)\" fill-rule=\"evenodd\" d=\"m0 0 461 369-284 58z\" transform=\"matrix(1 0 0 -1 0 427)\"></path></svg></div><div class=\"px-4 sm:px-6\"><div class=\"max-w-5xl mx-auto\"><div class=\"pb-12 md:pb-20\"><!-- Section header --><div class=\"max-w-3xl mx-auto text-center pb-12 md:pb-20\"><h2 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60\">Our story (so far)</h2></div><div class=\"md:flex justify-between space-x-6 md:space-x-8 lg:space-x-14\"><figure class=\"min-w-[240px]\"><img class=\"sticky top-8 mx-auto mb-12 md:mb-0 rounded-lg -rotate-[4deg]\" src=\"https://cdn.sonr.build/images/team.jpg\" width=\"420\" height=\"280\" alt=\"Team\"></figure><div class=\"max-w-[548px] mx-auto\"><div class=\"text-slate-400 space-y-6\"><p>We came together over a shared excitement about building a product that could solve our own problem of where our next favourite hack is coming from. But also a product that helps everyone thrive in this market: from founders and engineers to companies and investors.</p><p>Sonr is a product that connects people around the topics and ideas that fascinate them. <strong class=\"text-slate-50 font-medium\">The idea that we can use technology to take our experience</strong>, as security lovers, to a new dimension and leave the computer industry in better shape while we're at it.</p><p>You can dive into the atoms that make up a product, share the moments that move you and discuss the ideas you find compelling. We want to create a ground for <strong class=\"text-slate-50 font-medium\">discussion and bring knowledge together</strong>, while making it more accessible and easier to grasp.</p><p>Contrary to popular belief, this product is not random security. It has roots in a piece of classical literature, making it over 5 years old. Richard McClintock, a professor at <a class=\"text-purple-500 font-medium hover:underline\" href=\"#0\">Hampden-Sydney College</a> in Virginia, looked up one of the more obscure words, consectetur from a passage, and going through the cites of the word in classical literature, discovered the undoubtable source.</p><p>We all thrive on learning something new every day and everyone is constantly trying on different hats. We are working with new technologies while rethinking an old industry and are excited about all the possibilities and opportunities to innovate. It's a problem deeply ingrained in traditional sectors like startups and the wider service industry but which has been compounded in the past five to ten years by the emergence of the mostly tech-powered gig economy which has created a new generation of shift workers and indeed</p></div></div></div></div></div></div></section>")
>>>>>>> master
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

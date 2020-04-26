(this["webpackJsonp@jobsgowhere/ui"]=this["webpackJsonp@jobsgowhere/ui"]||[]).push([[0],{36:function(e,n,t){e.exports=t(63)},41:function(e,n,t){},63:function(e,n,t){"use strict";t.r(n);var r=t(0),a=t.n(r),o=t(29),l=t.n(o),c=(t(41),t(4)),u=t(9),i=t(1),m=t(2);function f(){var e=Object(i.a)(["\n        background-color: transparent;\n      "]);return f=function(){return e},e}function d(){var e=Object(i.a)(["\n        background-color: #fff;\n        color: var(--color-blue);\n      "]);return d=function(){return e},e}function v(){var e=Object(i.a)(["\n        ","\n      "]);return v=function(){return e},e}function b(){var e=Object(i.a)(["\n  background-color: var(--color-blue);\n  color: #fff;\n  background-color: var(--color-grey-200);\n  color: var(--color-darkblue);\n  border-radius: 0.875rem;\n  border: none;\n  font-size: 1rem;\n  font-weight: bold;\n  padding: 0.75rem 1.5rem;\n  cursor: pointer;\n\n  &:hover {\n    ","\n  }\n\n  ","\n"]);return b=function(){return e},e}function s(){var e=Object(i.a)(["\n  background-color: var(--color-blue);\n  color: #fff;\n"]);return s=function(){return e},e}var E=Object(m.a)(s()),p=m.b.button(b(),E,(function(e){return e.active||e.primary?Object(m.a)(v(),E):e.secondary?Object(m.a)(d()):e.text?Object(m.a)(f()):void 0}));function g(){var e=Object(i.a)(["\n  font-size: 3rem;\n"]);return g=function(){return e},e}function h(){var e=Object(i.a)(["\n  flex 1 1 auto;\n\n  ul {\n    margin: 0;\n    display: flex;\n    justify-content: flex-end;\n    align-items: center;\n    list-style: none;\n  }\n  li + li {\n    margin-left: 1rem;\n  }\n"]);return h=function(){return e},e}function j(){var e=Object(i.a)(["\n  flex: 0 0 auto;\n  width: 250px;\n  display: flex;\n  align-items: center;\n  justify-content: center;\n"]);return j=function(){return e},e}function y(){var e=Object(i.a)(["\n  grid-area: header;\n  display: flex;\n  padding: 2.25rem;\n"]);return y=function(){return e},e}var O=m.b.div(y()),x=m.b.div(j()),w=m.b.nav(h()),k=m.b.div(g());var C=function(e){return r.createElement(O,null,r.createElement(x,null,r.createElement(c.b,{to:"/"},r.createElement(k,null,"Logo"))),r.createElement(w,null,r.createElement("ul",null,r.createElement("li",null,"Already a member?"),r.createElement("li",null,r.createElement(c.b,{to:"/"},r.createElement(p,{text:!0},"Sign In"))),r.createElement("li",null,r.createElement(c.b,{to:"/"},r.createElement(p,{primary:!0},"Sign Up"))),r.createElement("li",null,r.createElement(c.b,{to:"/"},r.createElement(p,{secondary:!0},"New Post"))),r.createElement("li",null,r.createElement(c.b,{to:"/favourite"},"Favourite")),r.createElement("li",null,r.createElement(c.b,{to:"/"},"Profile")))))};function T(){var e=Object(i.a)(['\n  grid-area: footer;\n  color: var(--color-grey-300);\n\n  ul {\n    display: flex;\n    justify-content: center;\n    margin-top: 2rem;\n    list-style: none;\n    font-size: 0.875rem;\n  }\n  li:not(:last-child):after {\n    content: "\xb7";\n    display: inline-block;\n    margin: 0 0.3rem;\n  }\n  a {\n    color: currentColor;\n    &:hover {\n      color: var(--color-blue);\n    }\n  }\n']);return T=function(){return e},e}var A=m.b.div(T());var B=function(e){return r.createElement(A,null,r.createElement("ul",null,r.createElement("li",null,r.createElement("a",{href:"/"},"About")),r.createElement("li",null,r.createElement("a",{href:"/"},"Privacy")),r.createElement("li",null,r.createElement("a",{href:"/"},"Terms")),r.createElement("li",null,r.createElement("a",{href:"https://github.com/jobsgowhere/jobsgowhere",target:"_blank",rel:"noreferrer noopener"},"Github Repo")),r.createElement("li",null,r.createElement("a",{href:"/"},"Contact Us"))))};function z(){var e=Object(i.a)(['\n  background-color: var(--color-background);\n  display: grid;\n  grid-template-columns: 100px auto 100px;\n  grid-template-rows: 100px auto 100px;\n  grid-template-areas:\n    "header header header"\n    ". main ."\n    "footer footer footer";\n  min-height: 100vh;\n']);return z=function(){return e},e}var J=m.b.div(z());var P=function(e){var n=e.children;return r.createElement(J,null,r.createElement(C,null),n,r.createElement(B,null))};function _(){var e=Object(i.a)(["\n  grid-area: main;\n  display: flex;\n  justify-content: center;\n"]);return _=function(){return e},e}function F(){var e=Object(i.a)(["\n  max-width: 29.75rem;\n  width: 100%;\n  & + & {\n    margin-left: 1.375rem;\n  }\n"]);return F=function(){return e},e}var S=m.b.div(F()),I=m.b.div(_()),G=function(e){var n=e.children;return a.a.createElement(I,null,n)};G.Col=S;var U=G;var M=function(e){return r.createElement(U,null,r.createElement(U.Col,null,r.createElement("p",null,r.createElement(c.b,{to:"/posts"},r.createElement(p,{primary:!0},"Go to jobs board"))),r.createElement("p",null,r.createElement(c.b,{to:"/favourite"},r.createElement(p,{secondary:!0},"Favourite")))),r.createElement(U.Col,null,r.createElement("div",null,r.createElement("h2",null,"Buttons"),r.createElement("p",null,r.createElement(p,null,"Default Button")),r.createElement("p",null,r.createElement(p,{primary:!0},"Primary Button")),r.createElement("p",null,r.createElement(p,{secondary:!0},"Secondary Button")),r.createElement("p",null,r.createElement(p,{active:!0},"Active Button")))))},D=t(15);function R(){var e=Object(i.a)(["\n  background: transparent;\n  border: none;\n  cursor: pointer;\n  fill: ",";\n  padding: 0.1875rem 0.125rem;\n  outline: none;\n\n  &:hover {\n    fill: var(--color-blue);\n  }\n\n  svg {\n    display: block;\n  }\n"]);return R=function(){return e},e}var V=m.b.button(R(),(function(e){return e.active?"var(--color-blue)":"var(--color-grey-300)"})),L=function(e){var n=e.active,t=void 0!==n&&n,r=e.onClick;return a.a.createElement(V,{onClick:r,active:t},a.a.createElement("svg",{width:"22",height:"20",xmlns:"http://www.w3.org/2000/svg"},t?a.a.createElement("path",{d:"M20.633 4.647a6.093 6.093 0 0 0-1.334-1.94A6.219 6.219 0 0 0 14.93.93 6.26 6.26 0 0 0 11 2.315 6.26 6.26 0 0 0 7.07.93 6.219 6.219 0 0 0 2.7 2.706a6.057 6.057 0 0 0-1.825 4.33c0 .78.16 1.593.476 2.42.265.692.644 1.409 1.13 2.133.768 1.146 1.825 2.341 3.138 3.553a35.068 35.068 0 0 0 4.42 3.453l.556.356a.753.753 0 0 0 .808 0l.556-.357a35.517 35.517 0 0 0 4.42-3.452c1.312-1.211 2.37-2.407 3.138-3.553.485-.724.867-1.441 1.13-2.133.316-.827.476-1.64.476-2.42a5.968 5.968 0 0 0-.49-2.39z"}):a.a.createElement("path",{d:"M21.534 3.924a6.349 6.349 0 00-1.389-2.022 6.478 6.478 0 00-4.55-1.85A6.52 6.52 0 0011.5 1.495 6.52 6.52 0 007.406.052a6.478 6.478 0 00-4.551 1.85 6.31 6.31 0 00-1.902 4.51c0 .813.166 1.66.496 2.522.276.72.671 1.467 1.176 2.221.801 1.194 1.902 2.44 3.27 3.702 2.265 2.092 4.509 3.537 4.604 3.596l.579.37a.784.784 0 00.842 0l.579-.37a37.02 37.02 0 004.604-3.596c1.367-1.263 2.468-2.508 3.27-3.702.504-.754.902-1.501 1.176-2.221.33-.862.495-1.71.495-2.522.003-.862-.17-1.7-.51-2.488zM11.5 16.893s-8.691-5.57-8.691-10.481c0-2.488 2.058-4.505 4.597-4.505A4.61 4.61 0 0111.5 4.358a4.61 4.61 0 014.094-2.45c2.54 0 4.597 2.016 4.597 4.504 0 4.912-8.691 10.48-8.691 10.48z"})))};function W(){var e=Object(i.a)(["\n  width: 64px;\n  height: 64px;\n  border-radius: 32px;\n"]);return W=function(){return e},e}function H(){var e=Object(i.a)(["\n  margin-top: auto;\n  font-size: 0.875rem;\n  color: var(--color-grey-300);\n"]);return H=function(){return e},e}function N(){var e=Object(i.a)(["\n  margin-bottom: 1rem;\n  font-size: 1.125rem;\n  line-height: 1.4;\n"]);return N=function(){return e},e}function Y(){var e=Object(i.a)(["\n  margin: 0;\n  font-size: 0.875rem;\n  font-weight: 400;\n  color: var(--color-grey-300);\n"]);return Y=function(){return e},e}function $(){var e=Object(i.a)(["\n  font-size: 1rem;\n  margin: 0.25rem 0;\n"]);return $=function(){return e},e}function q(){var e=Object(i.a)(["\n  margin-left: auto;\n"]);return q=function(){return e},e}function K(){var e=Object(i.a)(["\n  display: flex;\n  margin-bottom: 0.5rem;\n"]);return K=function(){return e},e}function Q(){var e=Object(i.a)(["\n  flex: auto;\n  display: flex;\n  flex-direction: column;\n"]);return Q=function(){return e},e}function X(){var e=Object(i.a)(["\n  flex: 0 0 auto;\n  margin-right: 1.375rem;\n"]);return X=function(){return e},e}function Z(){var e=Object(i.a)(["\n  flex: auto;\n  display: flex;\n  flex-direction: row;\n  padding: 1rem;\n  padding-left: 0.8125rem;\n"]);return Z=function(){return e},e}function ee(){var e=Object(i.a)(['\n  flex: 0 0 auto;\n  display: flex;\n  flex-direction: row;\n  /* height: 172px; */\n  background-color: white;\n  border-radius: 0.875rem;\n  overflow: hidden;\n\n  & + & {\n    margin-top: 1rem;\n  }\n\n  &::before {\n    content: "";\n    width: 0.75rem;\n    background-color: ',";\n    flex: 0 0 auto;\n  }\n"]);return ee=function(){return e},e}var ne=m.b.div(ee(),(function(e){return e.active?"var(--color-blue)":"transparent"})),te=m.b.div(Z()),re=m.b.div(X()),ae=m.b.div(Q()),oe=m.b.div(K()),le=m.b.div(q()),ce=m.b.h2($()),ue=m.b.h3(Y()),ie=m.b.div(N()),me=m.b.div(H()),fe=m.b.img(W());var de=function(e){var n=e.active,t=e.data,a=e.onClick,o=e.handleFavouriteToggle;return r.createElement(c.b,{to:"/posts/".concat(t.id)},r.createElement(ne,{active:n,onClick:a},r.createElement(te,null,r.createElement(re,null,r.createElement(fe,{src:"https://api.adorable.io/avatars/64/abott@adorable.png"})),r.createElement(ae,null,r.createElement(oe,null,r.createElement("div",null,r.createElement(ce,null,"Arthur Simmmons"),r.createElement(ue,null,"Talent Hunter at ABCDEF company")),r.createElement(le,null,r.createElement(L,{active:t.favourite,onClick:o}))),r.createElement(ie,null,t.title),r.createElement(me,null,"Today \xb7 You have connected")))))};function ve(){var e=Object(i.a)(["\n  background-color: white;\n  border-radius: 0.875rem;\n  padding: 1.5rem;\n"]);return ve=function(){return e},e}var be=m.b.div(ve());var se=function(e){var n=e.data;return r.createElement(be,null,n.id,r.createElement("h2",null,n.title),r.createElement("div",null,n.description))},Ee=t(11),pe=t(34),ge=t.n(pe),he={jobs:[],talents:[],activePost:void 0};function je(e,n){switch(n.type){case"UPDATE_JOBS":var t=n.payload;return Object(Ee.a)({},e,{jobs:t.map((function(e){return Object(Ee.a)({},e,{active:!1})}))});case"SET_ACTIVE_JOB":var r=e.jobs.find((function(e){return e.id===n.payload}));return Object(Ee.a)({},e,{jobs:e.jobs.map((function(e){return Object(Ee.a)({},e,{active:e.id===n.payload})})),activePost:r});case"TOGGLE_FAVOURITE_JOB":return Object(Ee.a)({},e,{jobs:e.jobs.map((function(e){return e===n.payload?Object(Ee.a)({},e,{favourite:!e.favourite}):e}))});default:return e}}function ye(){var e=Object(i.a)(["\n  font-size: 0.8rem;\n  font-weight: 300;\n  text-align : center;\n"]);return ye=function(){return e},e}function Oe(){var e=Object(i.a)(["\n  font-size: 1rem;\n  font-weight: 400;\n  text-align : center;\n"]);return Oe=function(){return e},e}function xe(){var e=Object(i.a)(["\n  font-size: 1.5rem;\n  text-align : center;\n"]);return xe=function(){return e},e}function we(){var e=Object(i.a)(["\n  background-color: white;\n  border-radius: 0.875rem;\n  display: flex;\n  flex-direction: column;\n  align-items: center;\n  justify-content: center;\n  padding: 1.5rem;\n"]);return we=function(){return e},e}var ke=m.b.div(we()),Ce=m.b.h1(xe()),Te=m.b.h2(Oe()),Ae=m.b.p(ye());var Be=function(){return r.createElement(ke,null,r.createElement(Ce,null,"Ahh... I see you are looking for a job"),r.createElement(Te,null,"Tap a post on the left"),r.createElement(Ae,null,"\u201cEvery experience in your life is being orchestrated to teach you something you need to know to move forward\u201d",r.createElement("br",null),"- Brian Tracy"))};function ze(){var e=Object(i.a)(["\n  flex: 1;\n"]);return ze=function(){return e},e}function Je(){var e=Object(i.a)(["\n  flex: 1;\n  display: flex;\n  flex-direction: column;\n"]);return Je=function(){return e},e}var Pe=m.b.div(Je()),_e=m.b.div(ze());var Fe=function(){var e=function(){var e,n=a.a.useReducer(je,he),t=Object(D.a)(n,2),r=t[0],o=t[1],l=a.a.useCallback((function(e){o({type:"SET_ACTIVE_JOB",payload:e})}),[]),c=a.a.useCallback((function(e){o({type:"TOGGLE_FAVOURITE_JOB",payload:e})}),[]),i=a.a.useCallback((function(e){o({type:"UPDATE_JOBS",payload:e})}),[]),m=a.a.useMemo((function(){return{setActiveJob:l,toggleFavouriteJob:c,updateJobs:i}}),[l,c,i]),f=Object(u.f)("/posts/:id"),d=null===f||void 0===f||null===(e=f.params)||void 0===e?void 0:e.id;return a.a.useEffect((function(){ge.a.get("/api/jobs").then((function(e){i(e.data.jobs),d&&l(d)}))}),[d,l,i]),a.a.useEffect((function(){d&&l(d)}),[d,l]),[r,m]}(),n=Object(D.a)(e,2),t=n[0],o=n[1].toggleFavouriteJob;return r.createElement(U,null,r.createElement(U.Col,null,r.createElement(Pe,null,t.jobs.map((function(e){return r.createElement(de,{active:e.active,key:e.id,data:e,handleFavouriteToggle:function(n){n.stopPropagation(),o(e)}})})))),r.createElement(U.Col,null,r.createElement(_e,null,t.activePost?r.createElement(se,{data:t.activePost}):r.createElement(Be,null))))};var Se=function(e){return a.a.createElement(U,null,a.a.createElement(U.Col,null,a.a.createElement("h1",null,"Favourite Posts")),a.a.createElement(U.Col,null,"post detail"))};var Ie=function(){return r.createElement(c.a,null,r.createElement(P,null,r.createElement(u.c,null,r.createElement(u.a,{exact:!0,path:"/",component:M}),r.createElement(u.a,{path:"/posts",component:Fe}),r.createElement(u.a,{path:"/favourite",component:Se}))))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));l.a.render(a.a.createElement(a.a.StrictMode,null,a.a.createElement(Ie,null)),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()})).catch((function(e){console.error(e.message)}))}},[[36,1,2]]]);
//# sourceMappingURL=main.85d3fb6e.chunk.js.map
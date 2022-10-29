(()=>{var $=Object.defineProperty;var q=(t,n)=>{for(var e in n)$(t,e,{get:n[e],enumerable:!0})};var m={};q(m,{JSX:()=>d});function C(t,n){let e;n?e=n.querySelectorAll(t):e=document.querySelectorAll(t);let o=[];return e.forEach(r=>{o.push(r)}),o}function N(t,n){let e=C(t,n);switch(e.length){case 0:return;case 1:return e[0];default:console.warn(`found [${e.length}] elements with selector [${t}], wanted zero or one`)}}function l(t,n){let e=N(t,n);return e||console.warn(`no element found for selector [${t}]`),e}function k(t,n){return typeof t=="string"&&(t=l(t)),t.innerHTML=n,t}function d(t,n){let e=document.createElement(t);for(let o in n)if(o&&n.hasOwnProperty(o)){let r=n[o];o==="dangerouslySetInnerHTML"?k(e,r.__html):r===!0?e.setAttribute(o,o):r!==!1&&r!==null&&r!==void 0&&e.setAttribute(o,r.toString())}for(let o=2;o<arguments.length;o++){let r=arguments[o];if(Array.isArray(r))r.forEach(i=>{if(r==null)throw`child array for tag [${t}] is ${r}
${e.outerHTML}`;if(i==null)throw`child for tag [${t}] is ${i}
${e.outerHTML}`;typeof i=="string"&&(i=document.createTextNode(i)),e.appendChild(i)});else{if(r==null)throw`child for tag [${t}] is ${r}
${e.outerHTML}`;r.nodeType||(r=document.createTextNode(r.toString())),e.appendChild(r)}}return e}function h(){for(let t of Array.from(document.querySelectorAll(".menu-container .final")))t.scrollIntoView({block:"nearest"})}var u="mode-light",f="mode-dark";function y(){for(let t of Array.from(document.getElementsByClassName("mode-input"))){let n=t;n.onclick=function(){switch(n.value){case"":document.body.classList.remove(u),document.body.classList.remove(f);break;case"light":document.body.classList.add(u),document.body.classList.remove(f);break;case"dark":document.body.classList.remove(u),document.body.classList.add(f);break}}}}function T(){let t=document.getElementById("flash-container");if(t===null)return;let n=t.querySelectorAll(".flash");n.length>0&&setTimeout(()=>{for(let e of n){let o=e;o.style.opacity="0",setTimeout(()=>o.remove(),500)}},3e3)}function b(){for(let t of Array.from(document.getElementsByClassName("link-confirm"))){let n=t;n.onclick=function(){let e=n.dataset.message;return e&&e.length===0&&(e="Are you sure?"),confirm(e)}}}function E(){document.addEventListener("keydown",t=>{t.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}var M="--selected";function J(t){let n=t.parentElement.parentElement.querySelector("input");if(!n)throw"no associated input found";n.value="\u2205"}function L(){window.rituals.setSiblingToNull=J;let t={},n={};for(let e of Array.from(document.querySelectorAll(".editor"))){let o=e,r=()=>{t={},n={};for(let i of o.elements){let s=i;if(s.name.length>0)if(s.name.endsWith(M))n[s.name]=s;else{(s.type!=="radio"||s.checked)&&(t[s.name]=s.value);let p=()=>{let g=n[s.name+M];g&&(g.checked=t[s.name]!==s.value)};s.onchange=p,s.onkeyup=p}}};o.onreset=r,r()}}var O=[];function x(){let t=document.querySelectorAll(".color-var");if(t.length>0)for(let n of Array.from(t)){let e=n,o=e.dataset.var,r=e.dataset.mode;O.push(o),!(!o||o.length===0)&&(e.oninput=function(){X(r,o,e.value)})}}function X(t,n,e){let o=document.querySelector("#mockup-"+t);if(!o){console.error("can't find mockup for mode ["+t+"]");return}switch(n){case"color-foreground":c(o,".mock-main",e);break;case"color-background":a(o,".mock-main",e);break;case"color-foreground-muted":c(o,".mock-main .mock-muted",e);break;case"color-background-muted":a(o,".mock-main .mock-muted",e);break;case"color-link-foreground":c(o,".mock-main .mock-link",e);break;case"color-link-visited-foreground":c(o,".mock-main .mock-link-visited",e);break;case"color-nav-foreground":c(o,".mock-nav",e),c(o,".mock-nav .mock-link",e);break;case"color-nav-background":a(o,".mock-nav",e);break;case"color-menu-foreground":c(o,".mock-menu",e),c(o,".mock-menu .mock-link",e);break;case"color-menu-background":a(o,".mock-menu",e);break;case"color-menu-selected-foreground":c(o,".mock-menu .mock-link-selected",e);break;case"color-menu-selected-background":a(o,".mock-menu .mock-link-selected",e);break;default:console.error("invalid key ["+n+"]")}}function v(t,n,e){let o=t.querySelectorAll(n);if(o.length==0)throw"empty query selector ["+n+"]";o.forEach(r=>{e(r)})}function a(t,n,e){v(t,n,o=>o.style.backgroundColor=e)}function c(t,n,e){v(t,n,o=>o.style.color=e)}function S(){window.rituals.Socket=H}var H=class{constructor(n,e,o,r,i){this.debug=n,this.open=e,this.recv=o,this.err=r,this.url=w(i),this.connected=!1,this.pauseSeconds=1,this.pendingMessages=[],this.connect()}connect(){this.connectTime=Date.now(),this.sock=new WebSocket(w(this.url));let n=this;this.sock.onopen=()=>{n.connected=!0,n.pendingMessages.forEach(n.send),n.pendingMessages=[],n.debug&&console.log("WebSocket connected"),n.open("todo")},this.sock.onmessage=e=>{let o=JSON.parse(e.data);n.debug&&console.debug("in",o),n.recv(o)},this.sock.onerror=e=>()=>{n.err("socket",e.type)},this.sock.onclose=()=>{n.connected=!1;let e=n.connectTime?Date.now()-n.connectTime:0;0<e&&e<2e3?(n.pauseSeconds=n.pauseSeconds*2,n.debug&&console.debug(`socket closed immediately, reconnecting in ${n.pauseSeconds} seconds`),setTimeout(()=>{n.connect()},n.pauseSeconds*1e3)):(console.debug("socket closed after ["+e+"ms]"),n.connect())}}disconnect(){}send(n){if(this.debug&&console.debug("out",n),!this.sock)throw"not initialized";if(this.connected){let e=JSON.stringify(n,null,2);this.sock.send(e)}else this.pendingMessages.push(n)}};function w(t){if(t||(t="/connect"),t.indexOf("ws")==0)return t;let n=document.location,e="ws";return n.protocol==="https:"&&(e="wss"),t.indexOf("/")!=0&&(t="/"+t),e+`://${n.host}${t}`}function I(){var t=null;return m("div",null,"Test!!!")}function A(){l("table").appendChild(I())}function W(){window.rituals={},window.JSX=d,h(),y(),T(),b(),E(),L(),x(),S(),A()}document.addEventListener("DOMContentLoaded",W);})();
//# sourceMappingURL=client.js.map

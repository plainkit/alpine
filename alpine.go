// Package alpine provides Alpine.js attribute functions for the Plain component library.
package alpine

import (
	"github.com/plainkit/alpine/js"
	"github.com/plainkit/html"
)

// Core Alpine.js directives

// XData declares a new Alpine component and its data.
// Example: XData("{open: false, name: 'Alpine'}")
func XData(data string) html.Global {
	return html.ACustom("x-data", data)
}

// XInit runs code when a component initializes.
// Example: XInit("console.log('Component initialized')")
func XInit(code string) html.Global {
	return html.ACustom("x-init", code)
}

// Display and Rendering

// XShow shows or hides an element based on a condition.
// Example: XShow("open")
func XShow(expression string) html.Global {
	return html.ACustom("x-show", expression)
}

// XIf conditionally renders elements in the DOM.
// Must be used on a <template> tag.
// Example: XIf("user.isAdmin")
func XIf(expression string) html.Global {
	return html.ACustom("x-if", expression)
}

// XFor creates DOM elements by iterating through a list.
// Must be used on a <template> tag.
// Example: XFor("item in items")
func XFor(expression string) html.Global {
	return html.ACustom("x-for", expression)
}

// XHtml sets the inner HTML of an element.
// Example: XHtml("markdownContent")
func XHtml(expression string) html.Global {
	return html.ACustom("x-html", expression)
}

// XText sets the text content of an element.
// Example: XText("userName")
func XText(expression string) html.Global {
	return html.ACustom("x-text", expression)
}

// XCloak hides elements until Alpine is initialized.
// Typically used with CSS: [x-cloak] { display: none !important; }
func XCloak() html.Global {
	return html.ACustom("x-cloak", "")
}

// Event Handling and Binding

// XOn listens for browser events.
// Example: XOn("click", "open = !open")
func XOn(event, handler string) html.Global {
	return html.ACustom("x-on:"+event, handler)
}

// XOnClick is a shorthand for XOn("click", handler).
func XOnClick(handler string) html.Global {
	return html.ACustom("x-on:click", handler)
}

// XOnSubmit is a shorthand for XOn("submit", handler).
func XOnSubmit(handler string) html.Global {
	return html.ACustom("x-on:submit", handler)
}

// XOnChange is a shorthand for XOn("change", handler).
func XOnChange(handler string) html.Global {
	return html.ACustom("x-on:change", handler)
}

// XOnInput is a shorthand for XOn("input", handler).
func XOnInput(handler string) html.Global {
	return html.ACustom("x-on:input", handler)
}

// XOnKeydown is a shorthand for XOn("keydown", handler).
func XOnKeydown(handler string) html.Global {
	return html.ACustom("x-on:keydown", handler)
}

// XOnKeyup is a shorthand for XOn("keyup", handler).
func XOnKeyup(handler string) html.Global {
	return html.ACustom("x-on:keyup", handler)
}

// XOnMouseenter is a shorthand for XOn("mouseenter", handler).
func XOnMouseenter(handler string) html.Global {
	return html.ACustom("x-on:mouseenter", handler)
}

// XOnMouseleave is a shorthand for XOn("mouseleave", handler).
func XOnMouseleave(handler string) html.Global {
	return html.ACustom("x-on:mouseleave", handler)
}

// XBind dynamically sets HTML attributes.
// Example: XBind("class", "{ 'hidden': !open }")
func XBind(attribute, expression string) html.Global {
	return html.ACustom("x-bind:"+attribute, expression)
}

// XBindClass is a shorthand for XBind("class", expression).
func XBindClass(expression string) html.Global {
	return html.ACustom("x-bind:class", expression)
}

// XBindStyle is a shorthand for XBind("style", expression).
func XBindStyle(expression string) html.Global {
	return html.ACustom("x-bind:style", expression)
}

// XBindDisabled is a shorthand for XBind("disabled", expression).
func XBindDisabled(expression string) html.Global {
	return html.ACustom("x-bind:disabled", expression)
}

// XBindValue is a shorthand for XBind("value", expression).
func XBindValue(expression string) html.Global {
	return html.ACustom("x-bind:value", expression)
}

// XModel creates two-way data binding for form inputs.
// Example: XModel("searchQuery")
func XModel(expression string) html.Global {
	return html.ACustom("x-model", expression)
}

// XModelable makes a property modelable from child components.
// Example: XModelable("value")
func XModelable(expression string) html.Global {
	return html.ACustom("x-modelable", expression)
}

// Advanced Directives

// XTransition adds transitions to elements.
// Can be used with modifiers like x-transition:enter, x-transition:leave
func XTransition() html.Global {
	return html.ACustom("x-transition", "")
}

// XTransitionEnter specifies CSS classes for enter transition.
// Example: XTransitionEnter("transition ease-out duration-300")
func XTransitionEnter(classes string) html.Global {
	return html.ACustom("x-transition:enter", classes)
}

// XTransitionEnterStart specifies CSS classes for enter transition start.
// Example: XTransitionEnterStart("transform opacity-0 scale-90")
func XTransitionEnterStart(classes string) html.Global {
	return html.ACustom("x-transition:enter-start", classes)
}

// XTransitionEnterEnd specifies CSS classes for enter transition end.
// Example: XTransitionEnterEnd("transform opacity-100 scale-100")
func XTransitionEnterEnd(classes string) html.Global {
	return html.ACustom("x-transition:enter-end", classes)
}

// XTransitionLeave specifies CSS classes for leave transition.
// Example: XTransitionLeave("transition ease-in duration-300")
func XTransitionLeave(classes string) html.Global {
	return html.ACustom("x-transition:leave", classes)
}

// XTransitionLeaveStart specifies CSS classes for leave transition start.
// Example: XTransitionLeaveStart("transform opacity-100 scale-100")
func XTransitionLeaveStart(classes string) html.Global {
	return html.ACustom("x-transition:leave-start", classes)
}

// XTransitionLeaveEnd specifies CSS classes for leave transition end.
// Example: XTransitionLeaveEnd("transform opacity-0 scale-90")
func XTransitionLeaveEnd(classes string) html.Global {
	return html.ACustom("x-transition:leave-end", classes)
}

// XEffect re-evaluates an expression when dependencies change.
// Example: XEffect("console.log('Count changed:', count)")
func XEffect(expression string) html.Global {
	return html.ACustom("x-effect", expression)
}

// XRef references DOM elements.
// Example: XRef("myInput")
func XRef(name string) html.Global {
	return html.ACustom("x-ref", name)
}

// XTeleport moves elements to another location in the DOM.
// Example: XTeleport("#modal-root")
func XTeleport(selector string) html.Global {
	return html.ACustom("x-teleport", selector)
}

// XIgnore tells Alpine to ignore a block of HTML.
func XIgnore() html.Global {
	return html.ACustom("x-ignore", "")
}

// XId generates unique IDs for elements.
// Example: XId("['input', 'label']")
func XId(expression string) html.Global {
	return html.ACustom("x-id", expression)
}

// Shorthand helpers using @ syntax

// At is a shorthand for event listeners using @ syntax.
// Example: At("click", "open = !open") produces @click="open = !open"
func At(event, handler string) html.Global {
	return html.ACustom("@"+event, handler)
}

// AtClick is a shorthand for @click.
func AtClick(handler string) html.Global {
	return html.ACustom("@click", handler)
}

// AtSubmit is a shorthand for @submit.
func AtSubmit(handler string) html.Global {
	return html.ACustom("@submit", handler)
}

// AtChange is a shorthand for @change.
func AtChange(handler string) html.Global {
	return html.ACustom("@change", handler)
}

// AtInput is a shorthand for @input.
func AtInput(handler string) html.Global {
	return html.ACustom("@input", handler)
}

// AtKeydown is a shorthand for @keydown.
func AtKeydown(handler string) html.Global {
	return html.ACustom("@keydown", handler)
}

// AtKeyup is a shorthand for @keyup.
func AtKeyup(handler string) html.Global {
	return html.ACustom("@keyup", handler)
}

// AtMouseenter is a shorthand for @mouseenter.
func AtMouseenter(handler string) html.Global {
	return html.ACustom("@mouseenter", handler)
}

// AtMouseleave is a shorthand for @mouseleave.
func AtMouseleave(handler string) html.Global {
	return html.ACustom("@mouseleave", handler)
}

// Colon is a shorthand for attribute binding using : syntax.
// Example: Colon("class", "{ 'hidden': !open }") produces :class="{ 'hidden': !open }"
func Colon(attribute, expression string) html.Global {
	return html.ACustom(":"+attribute, expression)
}

// ColonClass is a shorthand for :class.
func ColonClass(expression string) html.Global {
	return html.ACustom(":class", expression)
}

// ColonStyle is a shorthand for :style.
func ColonStyle(expression string) html.Global {
	return html.ACustom(":style", expression)
}

// ColonDisabled is a shorthand for :disabled.
func ColonDisabled(expression string) html.Global {
	return html.ACustom(":disabled", expression)
}

// ColonValue is a shorthand for :value.
func ColonValue(expression string) html.Global {
	return html.ACustom(":value", expression)
}

// Event Modifiers (common ones as helper functions)

// AtClickAway listens for clicks outside the element.
func AtClickAway(handler string) html.Global {
	return html.ACustom("@click.away", handler)
}

// AtClickOutside listens for clicks outside the element (alias for away).
func AtClickOutside(handler string) html.Global {
	return html.ACustom("@click.outside", handler)
}

// AtClickPrevent prevents default behavior.
func AtClickPrevent(handler string) html.Global {
	return html.ACustom("@click.prevent", handler)
}

// AtClickStop stops event propagation.
func AtClickStop(handler string) html.Global {
	return html.ACustom("@click.stop", handler)
}

// AtSubmitPrevent prevents form submission.
func AtSubmitPrevent(handler string) html.Global {
	return html.ACustom("@submit.prevent", handler)
}

// AtKeydownEscape listens for escape key.
func AtKeydownEscape(handler string) html.Global {
	return html.ACustom("@keydown.escape", handler)
}

// AtKeydownEnter listens for enter key.
func AtKeydownEnter(handler string) html.Global {
	return html.ACustom("@keydown.enter", handler)
}

// AtKeydownWindow listens for keydown on window.
func AtKeydownWindow(handler string) html.Global {
	return html.ACustom("@keydown.window", handler)
}

// Model modifiers

// XModelLazy updates the model on change instead of input.
func XModelLazy(expression string) html.Global {
	return html.ACustom("x-model.lazy", expression)
}

// XModelNumber automatically converts the value to a number.
func XModelNumber(expression string) html.Global {
	return html.ACustom("x-model.number", expression)
}

// XModelDebounce debounces the model update.
// Example: XModelDebounce("searchQuery", "500ms")
func XModelDebounce(expression, delay string) html.Global {
	return html.ACustom("x-model.debounce."+delay, expression)
}

// JavaScript returns the embedded Alpine.js JavaScript content.
// This can be used to serve the Alpine.js library directly from your Go application
// without requiring external CDN dependencies.
//
// Example usage:
//
//	http.HandleFunc("/js/alpine.min.js", func(w http.ResponseWriter, r *http.Request) {
//	    w.Header().Set("Content-Type", "application/javascript")
//	    w.Write(alpine.JavaScript())
//	})
func JavaScript() []byte {
	return js.AlpineMinJS
}

document.body.addEventListener("htmx:responseError", function(event) {
  console.log(event.detail);
  showError(event.detail.xhr.response);
});

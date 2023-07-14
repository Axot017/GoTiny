htmx.defineExtension('format-dates', {
    onEvent: function (name, evt) {
        if (name === "htmx:configRequest") {
          const elements_to_format = evt.detail.elt.querySelectorAll('[format-date]');
          for (var i = 0; i < elements_to_format.length; i++) {
            const element = elements_to_format[i];
            const name = element.getAttribute('name');
            if (name && evt.detail.parameters[name]) {
              evt.detail.parameters[name] = new Date(evt.detail.parameters[name]).toISOString();
            }
          }
        }
    },
});

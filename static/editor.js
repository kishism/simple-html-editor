document.addEventListener("DOMContentLoaded", function () {
    const editor = document.getElementById("editor");
    const hiddenInput = document.getElementById("hidden-content");
  
    function syncContent() {
      hiddenInput.value = editor.innerHTML;
    }
  
    // Sync hidden input whenever the editor content changes
    editor.addEventListener("input", syncContent);
  
    // Also sync on page load (in case InitialHTML is preloaded)
    syncContent();
  });
  
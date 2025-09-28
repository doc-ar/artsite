document.body.addEventListener("htmx:afterSwap", () => {
  if (window.dragscroll) {
    window.dragscroll.reset(); // Re-initialize if needed
  }
});

// Scroll buttons
function portfolioScrollLeft() {
  const container = document.getElementById("portfolio-container");
  container.style.scrollBehavior = "smooth";
  container.scrollLeft -= 800;

  setTimeout(() => {
    container.style.scrollBehavior = "auto";
  }, 300);
}
function portfolioScrollRight() {
  const container = document.getElementById("portfolio-container");
  container.style.scrollBehavior = "smooth";
  container.scrollLeft += 800;

  setTimeout(() => {
    container.style.scrollBehavior = "auto";
  }, 300);
}

// Utility to initialize a menu in alpine
function cookieMenu() {
  return {
    menu: getCookie("menu") || "upload",
    init() {
      this.$watch("menu", (value) => setCookie("menu", value));
    },
  };
}

// Utility to get a cookie by name
function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(";").shift();
  return null;
}

// Utility to set a cookie
function setCookie(name, value, days = 7) {
  const expires = new Date(Date.now() + days * 864e5).toUTCString();
  document.cookie = `${name}=${value}; expires=${expires}; path=/`;
}

// Function to close edit art form
function closeArtEditForm() {
  const form_parent = document.querySelector(
    "#admin_art_edit_form",
  ).parentElement;
  if (form_parent) {
    form_parent.outerHTML = "";
  }
}
// Function to close edit series form
function closeSeriesEditForm() {
  const form_parent = document.querySelector(
    "#admin_series_edit_form",
  ).parentElement;
  if (form_parent) {
    form_parent.outerHTML = "";
  }
}

// Function for search bar
function searchList() {
  const searchTerm = document
    .getElementById("gallery_search_input")
    .value.trim()
    .toLowerCase();
  const listItems = document.querySelectorAll("#gallery_item_list li");

  // Step 1: Create array with titles and element references
  const searchableItems = Array.from(listItems).map((li) => ({
    title: li.dataset.title.toLowerCase(),
    element: li,
  }));

  // Step 2: Set up Fuse with just the 'title' key
  const fuse = new Fuse(searchableItems, {
    keys: ["title"],
    threshold: 0.3, // you can make this 0.0 for exact match or increase for fuzzier match
  });

  // Step 3: Search or show all if empty
  let results;
  if (searchTerm === "") {
    results = searchableItems.map((item) => ({ item }));
  } else {
    results = fuse.search(searchTerm);
  }

  // Step 4: Hide all, then show only matched
  searchableItems.forEach((item) => {
    item.element.style.display = "none";
  });

  results.forEach((result) => {
    result.item.element.style.display = "list-item";
  });
}

// Sorting options
function sortGallery() {
  const sortBy = document.getElementById("gallery_sort_select").value;
  const ul = document.getElementById("gallery_item_list");
  const listItems = Array.from(ul.querySelectorAll("li"));

  let sortedItems = [...listItems];

  switch (sortBy) {
    case "title":
      sortedItems.sort((a, b) =>
        a.dataset.title.localeCompare(b.dataset.title),
      );
      break;

    case "dimensions":
      sortedItems.sort((a, b) => {
        const areaA =
          parseFloat(a.dataset.width) * parseFloat(a.dataset.height);
        const areaB =
          parseFloat(b.dataset.width) * parseFloat(b.dataset.height);
        return areaA - areaB; // smallest to largest
      });
      break;

    case "default":
    default:
      // Optional: reset to original order if needed
      return;
  }

  // Re-append sorted items to UL
  sortedItems.forEach((item) => ul.appendChild(item));
}

// Filtering
function filterGallery() {
  const selectedMediums = Array.from(
    document.querySelectorAll("#medium_filters input:checked"),
  ).map((cb) => cb.value);
  const selectedCategories = Array.from(
    document.querySelectorAll("#category_filters input:checked"),
  ).map((cb) => cb.value);

  const listItems = Array.from(
    document.querySelectorAll("#gallery_item_list li"),
  );

  listItems.forEach((li) => {
    const medium = li.dataset.medium;
    const category = li.dataset.category;

    const mediumMatch =
      selectedMediums.length === 0 || selectedMediums.includes(medium);
    const categoryMatch =
      selectedCategories.length === 0 || selectedCategories.includes(category);

    li.style.display = mediumMatch && categoryMatch ? "list-item" : "none";
  });
}

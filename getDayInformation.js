// thanku Bill Gale
  
const descriptions = document.getElementsByClassName("day-desc");
const paragraphs = Array.from(document.getElementsByTagName("p")).filter(
  (i) => i.parentNode.className != "day-desc"
);

const readme = [
  descriptions[0],
  paragraphs[0],
  descriptions[1],
  paragraphs[1],
].map((i) => i.outerHTML);

copy(readme.join(""));
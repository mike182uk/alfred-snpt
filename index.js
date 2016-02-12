var getStdin = require('get-stdin');

getStdin().then(function(input) {
  var items = input.split('\n').map(function(snippet) {
    if (!snippet) {
      return;
    }

    var item = '<item arg="' + snippet + '">';
    item += '<title>' + getTitle(snippet) + '</title>';
    item += '<subtitle>' + getSubTitle(snippet) + '</subtitle>';
    item += '</item>';

    return item;
  });

  console.log('<?xml version="1.0"?><items>' + items.join('') + '</items>');
});

function getTitle(snippet) {
  return snippet.substring(0, snippet.lastIndexOf('-') - 1);
}

function getSubTitle(snippet) {
  snippet = snippet.replace(/(\[(.*)\])/, '').trim();

  return snippet.substring(snippet.lastIndexOf('-') + 2);
}

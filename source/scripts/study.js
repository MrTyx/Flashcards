const UID = {{.ID}}
const ORDER = {{.Order}}
const BUCKET_URL = {{.BucketURL}}
let data = []
let index = 0;

$(document).ready(() => {
  function init() {
    $('#h-text').html('Loading...')
    $.getJSON(`order/${ORDER}`, json => {
      if (json === null || json.length === 0) {
        $('#h-text').html('No cards to study!')
        $('#btn-done').removeClass('hidden')
        return
      }
      data = json;
      $('#h-text').html('What country is this?')
      $('#img-flag').attr('src', getImageURL(data[index].Code));
      $('#btn-reveal').removeClass('hidden')
      $('#progress').removeClass('hidden')
    })
  }

  function getImageURL(code) {
    return `${BUCKET_URL}/flags/normal/${code}.png`
  }

  function getMiniImageURL(code) {
    return `${BUCKET_URL}/flags/mini/${code}.png`
  }

  function reveal() {
    $('#h-text').html(data[index].Name)
    $('#btn-reveal').addClass('hidden')
    $('#btn-continue').removeClass('hidden')
    $('#progress').css('width', ((index+1) / data.length * 100) + "%")
  }

  function score() {
    $.ajax({
      url: `study/${data[index].Code}/${UID}`,
      type: 'GET',
      error: function(data) {
        $('#alert-error').removeClass('hidden').html('<span>Uh oh, something went wrong.</span><br /><span>Are you online?</span>')
        $('#img-flag').addClass('hidden');
        $('#h-text').addClass('hidden');
        $('#progress').addClass('hidden');
        $('#btn-continue').addClass('hidden');
        $('#btn-done').removeClass('hidden');
        console.log(data)
      }
    })
    index++;
    if (index == data.length) {
      end();
    } else {
      next();
    }
  }

  function next() {
    $('#img-flag').attr('src', '');
    $('#btn-continue').addClass('hidden')
    $('#h-text').html('Loading...')
    $('#img-flag').one("load", () => {
      $('#h-text').html('What country is this?')
      $('#btn-reveal').removeClass('hidden')
    }).attr('src', getImageURL(data[index].Code))
  }

  function end() {
    $('#h-text').html('Good job!')
    $('#table-results').removeClass('hidden')
    $('#img-flag').addClass('hidden')
    $('#img-flag').attr('src', '')
    $('#progress').addClass('hidden')
    $('#btn-continue').addClass('hidden')
    $('#btn-done').removeClass('hidden')

    for (let i = 0; i < data.length; i++) {
      $('#table-results > tbody').append(`<tr>
          <td class="text-left"><img src="${getMiniImageURL(data[i].Code)}" /></td>
          <td class="text-right">${data[i].Name}</td>
        </tr>`)
    }
  }

  $('#btn-reveal').click(reveal)
  $('#btn-continue').click(() => { score(); })
  $('#btn-done').click(() => { window.location.replace('http://s3394330-flashcards.appspot.com/'); })
  init();
})

const UID = {{.ID}}
const ORDER = {{.Order}}
const BUCKET_URL = {{.BucketURL}}
let data = []
let index = 0;

$(document).ready(() => {

  function init() {
    $('#h-text').html('Loading...')
    $.getJSON(`due/${UID}`, json => {
      if (json === null) {
        $('#h-text').html('No cards to study!')
        $('#img-flag').addClass('hidden');
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
    $('#btn-answer').removeClass('hidden')
    $('#progress').css('width', ((index+1) / data.length * 100) + "%")

  }

  function score(ratio, score) {
    data[index].Score = score;
    $.ajax({
      url: `review/${data[index].Code}/${ratio}/${UID}`,
      type: 'GET',
      error: function(data) {
        $('#alert-error').removeClass('hidden').html('<span>Uh oh, something went wrong.</span><br /><span>Are you online?</span>');
        $('#img-flag').addClass('hidden');
        $('#h-text').addClass('hidden');
        $('#progress').addClass('hidden');
        $('#btn-continue').addClass('hidden');
        $('#btn-done').removeClass('hidden');
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
    $('#btn-answer').addClass('hidden')
    $('#h-text').html('Loading...')
    $('#img-flag').one("load", () => {
      $('#h-text').html('What country is this?')
      $('#btn-reveal').removeClass('hidden')
    }).attr('src', getImageURL(data[index].Code))
  }

  function end() {
    $('#h-text').html('Good job!')
    $('#table-results').removeClass('hidden');
    $('#img-flag').addClass('hidden')
    $('#img-flag').attr('src', '')
    $('#progress').addClass('hidden')
    $('#btn-answer').addClass('hidden')
    $('#btn-done').removeClass('hidden')

    for (let i = 0; i < data.length; i++) {
      if (data[i].Score === 'Forgot' || data[i].Score === 'Hard') {
        $('#table-results > tbody').append(`<tr>
          <td class="text-left"><img src="${getMiniImageURL(data[i].Code)}" /></td>
          <td class="text-left">${data[i].Name}</td>
          <td class="text-right">${data[i].Score}</td>
        </tr>`)
      }
    }
  }

  $('#btn-reveal').click(reveal)
  $('#btn-repeat').click(() => { score(0.1, "Forgot"); })
  $('#btn-hard').click(() => { score(0.75, "Hard"); })
  $('#btn-normal').click(() => { score(1.25, "Normal"); })
  $('#btn-easy').click(() => { score(2.5, "Easy"); })
  $('#btn-done').click(() => { window.location.replace('http://s3394330-flashcards.appspot.com/'); })
  init();
})

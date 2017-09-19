// let data = [
// 	{
// 		image: "flags/au.png",
// 		answer: "Australia"
// 	}, {
// 		image: "flags/my.png",
// 		answer: "Malaysia"
// 	}, {
// 		image: "flags/no.png",
// 		answer: "Norway"
// 	}, {
// 		image: "flags/cy.png",
// 		answer: "Cyprus"
// 	}, {
// 		image: "flags/de.png",
// 		answer: "Germany"
// 	}, {
// 		image: "flags/kp.png",
// 		answer: "Best Korea"
// 	}, {
// 		image: "flags/mz.png",
// 		answer: "Mozambique"
// 	}, {
// 		image: "flags/za.png",
// 		answer: "South Africa"
// 	}
// ]

const UID = ""
let data = []
let index = 0;


$(document).ready(() => {

	function init() {
		$('#h-text').html('Loading...')
		$('#btn-answer').hide()
		$('#btn-done').hide()
		$.getJSON("https://s3394330-flashcards.appspot.com/due/123456", json => {
			data = json;
			console.log(json);
			$('#h-text').html('What country is this?')
			$('#img-flag').attr('src', getImageURL(data[index].Code));
			$('#btn-answer').hide()
			$('#btn-done').hide()
		})
	}

	function getImageURL(code) {
		let string = `https://s3-ap-southeast-2.amazonaws.com/s3394330-flashcards/flags/normal/${code}.png`
		console.log(string)
		return string
	}

	function reveal() {
		$('#h-text').html(data[index].Name)
		$('#btn-reveal').hide()
		$('#btn-answer').show()
		index++;
		$('#progress').css('width', (index / data.length * 100) + "%")

	}

	function score(ratio) {
		console.log(ratio);
		if (index == data.length) {
			end();
		} else {
			next();
		}
	}

	function next() {
		$('#img-flag').attr('src', '');
		$('#btn-answer').hide()
		$('#h-text').html('Loading...')
		$('#img-flag').one("load", () => {
      $('#h-text').html('What country is this?')
      $('#btn-reveal').show()
    }).attr('src', getImageURL(data[index].Code))
	}

	function end() {
		$('#h-text').html('Good job!')
		$('#img-flag').attr('src', '')
		$('#progress').hide()
		$('#btn-answer').hide()
		$('#btn-done').show()
	}

	$('#btn-reveal').click(reveal)
	$('#btn-repeat').click(() => { score(-1); })
	$('#btn-hard').click(() => { score(0.5); })
	$('#btn-normal').click(() => { score(2); })
	$('#btn-easy').click(() => { score(4); })
	$('#btn-done').click(() => { window.location.replace('http://s3394330-flashcards.appspot.com/'); })
	init();
})

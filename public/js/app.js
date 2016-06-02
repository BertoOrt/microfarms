$(function () {
  $('.button-collapse').sideNav();
  $('.dropdown-button').dropdown();
  $('.parallax').parallax();
  $('.logout').click(function () {
    logOut();
  })
  $('.find-our-produce').click(function () {
    window.location = "/find-our-produce"
  })
  var options = [
    {selector: 'article img', offset: 200, callback: function() {
      Materialize.fadeInImage("article img");
    } }
  ];
  Materialize.scrollFire(options);
  if(localStorage.token) {
    $.ajaxSetup({
      headers: {
           Authorization: 'Bearer ' + localStorage.token
      }
    });
    $.get('/auth')
      .done(function (response) {
        if (!response.authorized) {
          console.info(response.error);
          logOut();
        } else {
          $('.login').css('display', 'none');
          $('.logout').css('display', 'block');
          var user = getUser();
          display(user);
        }
      })
      .fail(function (error) {
        console.error(error)
      })
  } else {
    $('.login').css('display', 'block');
  }
})

function logOut() {
  localStorage.token = '';
  window.location = '/';
}

function getUser() {
  if(localStorage.token) {
    return JSON.parse(atob(localStorage.token.split('.')[1])).user;
  }
}

function display(user) {
  $('main').append('<h1>' + user.name + '</h1>')
  $('main').append('<img class="picture" src="' + user.picture + '">')
}

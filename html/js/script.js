
function getTimeStamp() {
  // https://stackoverflow.com/questions/3552461/how-to-format-a-javascript-date
  // https://stackoverflow.com/questions/5631384/remove-everything-after-a-certain-character
  return new Date().toISOString().replace("T", " ").replace("Z", "");
} // getTimeStamp

// Put a timestamp to eliza's first message
$("#first-timestamp").html(getTimeStamp);

$("li").slideDown( 300 );

// Submit the form when the key enter is pressed
$("input").on("keyup", function (e) {
  if (e.which == 13) {
    var text = $(this).val();
    if (text !== "") {
      $("form").submit(e)
    }
  }
}); // $("#mytext").on

// What to do when the form is submitted
$("form").submit(function(event){
  event.preventDefault();
  // console.log( "form submitted" );

  var userInput = $('input[name=user-input]').val();

  var formData = JSON.stringify({
    'UserMessage' : userInput
  });

  // console.log(formData);

  // https://stackoverflow.com/questions/8701812/clear-form-after-submission-with-jquery
  $( "form" ).each(function(){
    this.reset();
  });

  
  // console.log(getTimeStamp()));

  // https://stackoverflow.com/questions/1145208/jquery-how-to-add-li-in-an-existing-ul
  $('<li>')
    .append($('<div>')
      .append($('<div>')
        .append($('<p>')
          .append(userInput))
          .append($('<p>')
            .append($('<small>')
              .append(getTimeStamp())
            )
          )
        .addClass("text text-r"))
      .append($('<div>')
        .append($('<img>')
          .addClass("img-circle")
          .css("width", "100%")
          .attr("src", "/images/user-avatar.png"))
        .addClass("avatar")
        .css("padding", "0px 10px 0px 10px"))
      .addClass("msj-rta macro"))
    .css("width", "100%")
    .appendTo( $("ul") ); // $('<li>')

    $("li").slideDown( 300 );

    $('<li>')
    .append($('<div>')
      .append($('<div>')
        .append($('<img>')
          .addClass("img-circle")
          .css("width", "100%")
          .attr("src", "/images/eliza-avatar.png"))
        .addClass("avatar")
      .css("padding", "0px 10px 0px 10px"))
      .append($('<div>')
        .append($('<p>')
          .append("Typing")
          .append($('<span>')
            .addClass("dots"))
          .addClass("server-response"))
          .append($('<p>')
            .append($('<small>')
            )
          )
        .addClass("text text-l"))
      .addClass("msj macro"))
    .css("width", "100%")
    .appendTo( $("ul") ); // $('<li>')

    $("li").slideDown( 400 );

    setInterval(function() {
      var th = $('.dots');
      if(th.text().length < 5){
          th.text(th.text()+".");
      }else{
          th.text("");
      }
    }, 300);

  // https://scotch.io/tutorials/submitting-ajax-forms-with-jquery
  $.ajax({
  url: '/ajax',
  type: "post",
  contentType: 'application/json; charset=utf-8',
  data: formData,
  dataType: 'json',
  encode: true 
  })
  .done(function(response) {
    // console.log(response)

    // $('<li>')
    //   .append($('<div>')
    //     .append($('<div>')
    //       .append($('<img>')
    //         .addClass("img-circle")
    //         .css("width", "100%")
    //         .attr("src", "/images/eliza-avatar.png"))
    //       .addClass("avatar")
    //     .css("padding", "0px 10px 0px 10px"))
    //     .append($('<div>')
    //       .append($('<p>')
    //         .append(response.ServerMessage))
    //         .append($('<p>')
    //           .append($('<small>')
    //             .append(getTimeStamp())
    //           )
    //         )
    //       .addClass("text text-l"))
    //     .addClass("msj macro"))
    //   .css("width", "100%")
    //   .appendTo( $("ul") ); // $('<li>')

      // $("li").slideDown( 300 );

      $(".server-response").last().html(response.ServerMessage);
      $("small").last().html(getTimeStamp());

  }) // $.ajax().done()

  .fail(function(data) {
    // show any errors
    console.log(data);
  }) // $.ajax().fail()
  .always(function(data) {

  }) // $.ajax().always
  ; // $.ajax

}); // $("#input-form").submit

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

  // https://stackoverflow.com/questions/3552461/how-to-format-a-javascript-date
  // https://stackoverflow.com/questions/5631384/remove-everything-after-a-certain-character
  // var timestamp = new Date().toISOString().replace("T", " ").replace(/\..*/, "")
  var timestamp = new Date().toISOString().replace("T", " ").replace("Z", "")
  
  // console.log(timestamp);

  // https://stackoverflow.com/questions/1145208/jquery-how-to-add-li-in-an-existing-ul
  $('<li>')
    .append($('<div>')
      .append($('<div>')
        .append($('<p>')
          .append(userInput))
          .append($('<p>')
            .append($('<small>')
              .append(timestamp)
            )
          )
        .addClass("text text-r"))
      .addClass("msj-rta macro"))
    .css("width", "100%")
    .appendTo( $("ul") ); // $('<li>')

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
    // $('#response').append(response.ElizaMessage);

  $('<li>')
    .append($('<div>')
      .append($('<div>')
        .append($('<p>')
          .append(response.ServerMessage))
          .append($('<p>')
            .append($('<small>')
              .append(timestamp)
            )
          )
        .addClass("text text-l"))
      .addClass("msj-rta macro"))
    .css("width", "100%")
    .appendTo( $("ul") ); // $('<li>')

  }) // $.ajax().done()

  .fail(function(data) {
    // show any errors
    console.log(data);
  }) // $.ajax().fail()
  ; // $.ajax

}); // $("#input-form").submit
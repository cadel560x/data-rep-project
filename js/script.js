$(document).ready(function (){
    // console.log( "document loaded" );

    $("form").submit(function(event){
      event.preventDefault();
      // console.log( "form submitted" );

      var formData = JSON.stringify({
        'UserInput' : $('input[name=user-input]').val()
      });

      // console.log(formData);

      $.ajax({
        url: '/ajax',
        type: "post",
        contentType: 'application/json; charset=utf-8',
        // data: JSON.stringify({ UserInput: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit.'}),
        data: formData,
        dataType: 'json',
        encode: true
      })
      .done(function(r) {
          $('#response').append(r.UserInput);
      })
      .fail(function(data) {
        // show any errors
        // best to remove for production
        console.log(data);
      })
      ; // $.ajax

      // $.post("/ajax", function(response){
      // $.post("/ajax", JSON.stringify({ UserInput: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit.'}), function(response){
      // $.post("/ajax", JSON.stringify({ UserInput: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit.'})
      //   // $("#mylist").append('<li>' + data.user + ' <i>' + data.joined + '</i></li>');
      //   // $('#response').append(response.UserInput);
      //   // console.log( "ajax submitted" );
      //   // console.log( "response", response );
      // // }, 'json'); // $.post
      // // })
      // , 'json')
      // .done(function(response){
      //   $('#response').append(response.UserInput);
      //   console.log( "ajax submitted" );
      //   console.log( "response", response );
      // })
      // .fail(function(response){
      //   // $("#mylist").append('<li>' + data.user + ' <i>' + data.joined + '</i></li>');
      //   $('#response').append(response.UserInput);
      //   console.log( "ajax submitted" );
      //   console.log( "response", response );
      // // }, 'json'); // $.post
      // }); // $.post

    }); // $("#input-form").submit
  }); // $( document ).ready
(function($) {

  $("#imgsrc").change(function() {
    $("#cover-preview").attr("src",$("#imgsrc").val())
  })

  $("#btnGetData").click(function(e) {
    // clear alert
    $("#status-alert").hide()
    $("#status-alert").removeClass("alert-info")
    $("#status-alert").removeClass("alert-primary")
    $("#status-alert").removeClass("alert-success")
    $("#status-alert").removeClass("alert-warning")
    $("#status-alert").removeClass("alert-danger")

    // get $(this) button
    var btn = $(this)

    // disable button
    btn.attr("disabled", "disabled")

    // show spinner
    $("#spinner").show()

    // get data
    var isbn = $("#qisbn").val()
    var url = "/api/v1/admin/client/isbn/" + isbn

    // send request
    $.ajax({
      url: url
    }).done(function(data, textStatus, jqXHR) {
      var status = jqXHR.status

      if (status == 204) {
        $("#status-alert").addClass("alert-info").text("No data found :(")
        return
      }

      console.log(data)
      $("#isbn").val(data.Isbn);
      $("#isbn10").val(data.Isbn10);
      $("#title").val(data.Title);
      var authors = ""
      for(i=0; i<data.Authors.length; i++){
        authors += data.Authors[i].Name
        if (i != data.Authors.length - 1) {
          authors += "; "
        }
      }
      $("#authors").val(authors)
      $("#desc").val(data.Desc)
      $("#lang").val(data.Lang)
      $("#imgsrc").val(data.ImgSrc)
      $("#pages").val(data.Pages)
      $("#goodreadsid").val(data.GoodReadsId)
      $("#googlebooksid").val(data.GoogleBooksId)
      $("#openlibraryid").val(data.OpenLibraryId)

      // show book cover
      $("#cover-preview").attr("src",data.ImgSrc)

      // update alert
      $("#status-alert").addClass("alert-primary").text("Successfully Found Book")
    }).fail(function( jqXHR, textStatus, errorThrown ) {
      // log errors
      console.log(jqXHR, textStatus, errorThrown)

      // update alert
      $("#status-alert").addClass("alert-danger").text("Error: Invalid ISBN")
    }).always(function() {
      // re-enable button
      btn.removeAttr("disabled")

      // hide spinner
      $("#spinner").hide()

      // show alert
      $("#status-alert").show()
    })
  })
})($)

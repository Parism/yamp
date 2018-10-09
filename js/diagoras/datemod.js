$( function() {
    $(".datepicker").datepicker({
      dateFormat: 'D dd/mm/yy',
      dayNamesMin: ['Κυρ', 'Δευ', 'Τρι', 'Τετ', 'Πεμ', 'Παρ', 'Σατ'],
      monthNames: ['Ιανουάριος',
      'Φεβρουάριος',
      'Μάρτιος',
      'Απρίλιος',
      'Μάιος',
      'Ιούνιος',
      'Ιούλιος',
      'Αύγουστος',
      'Σεπτέμβριος',
      'Οκτώβριος',
      'Νοέμβριος',
      'Δεκέμβριος',
    ],
    dayNamesShort:[
      'Κυρ',
      'Δευ',
      'Τρι',
      'Τετ',
      'Πεμ',
      'Παρ',
      'Σαβ',
    ],
    monthNamesShort:[
      'Ιανουαρίου',
      'Φεβρουαρίου',
      'Μαρτίου',
      'Απριλίου',
      'Μαΐου',
      'Ιουνίου',
      'Ιουλίου',
      'Αυγούστου',
      'Σεπτεμβρίου',
      'Οκτωβρίου',
      'Νοεμβρίου',
      'Δεκεμβρίου',
    ],
    onSelect: function(dateText, evt) {
      var id = $(this).attr("id");
      var date = dateText.split(" ")[1]
      $('#'+id+"hidden").val(date);
    }  
  });
  });
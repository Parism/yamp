var constdata;
var state = 0;
$(document).ready(function() {
    $("#dynform").submit(function(e){
        e.preventDefault(e);
        $("#result").html("Παρακαλώ περιμένετε");
        var date = $("#datehidden").val();
        var label = $("#label").val();
        $.ajax({
            url: "/getdyn",
            data: {
                date:date,
                label:label,
            },
            accepts: {
                text: "application/json; charset=utf-8"
            },
            error: function(xhr, status, error) {
                $("#result").html(xhr.responseText);
              },
            success: function(data){
                constdata = data
                parseDataMin(constdata)
            },
            timeout: 10000 // sets timeout to 10 seconds
        });
    });
});

function parseDataFull(data){
    $("#result").html("")
    $("#result").append("<button id=\"datatoggle\">Μερική</button>")
    for (var rank in data.rankmap) {
        if (data.rankmap[rank]){
            $("#result").append("<h3>"+rank+" "+data.rankmap[rank].length+"</h3>")
            $("#result").append("<ul>")
            if (data.rankmap.hasOwnProperty(rank)) {           
                for (var person in data.rankmap[rank]){
                    $("#result").append("<li>"+data.rankmap[rank][person].surname+" "+data.rankmap[rank][person].name+"</li>")        
                }
            }
            $("#result").append("</ul>")
        }
    }
    $("#datatoggle").click(function(e){
        if (state == 0){
            state = 1
            parseDataFull(constdata)
        }else{
            state = 0
            parseDataMin(constdata)
        }
    });
};

function parseDataMin(data){
    $("#result").html("")
    $("#result").append("<button id=\"datatoggle\">Αναλυτική</button>")
    for (var rank in data.rankmap) {
        if (data.rankmap[rank]){
            $("#result").append("<h3>"+rank+" "+data.rankmap[rank].length+"</h3>")
        }
    }
    $("#datatoggle").click(function(e){
        if (state == 0){
            state = 1
            parseDataFull(constdata)
        }else{
            state = 0
            parseDataMin(constdata)
        }
    });
}
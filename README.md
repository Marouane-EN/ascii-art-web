index.html style.css home page

result.html style1.css result page

error.html error.css error page bhal 404 page not found rani 3ad na5dam 3liha

-------main.go--------

/static/ : bash servi css dyalak

/ : home page

/ascii-art : result page

/ascii-art1 : hta hi result page, dartha bhal haka 3la hsab error ila user dar chi invalid character o caf fl home page, rah user dar wahad request l servar bash yamshi /ascii-art, brit user ila dar bhal had error yb9a ri fl home page, bash na3raf bili user rah tlab had request o howa rah f home page knchof fal back bili path == "/ascii-art" ya3ni rah user 3ad f home page, kanmchi ana refresher home page o n5orajlah waha error bhal notification,
ila kan f result page katjini fal back path == "/ascii-art1" ya3ni rah user f result page o rah dar wahad request man result page,  refresher result page o n5orajlah waha error bhal notification

-------handlers.go--------

hna fin kan parse o nhandli les errors o kan executer les templetes

--------ascii.go----------

hna fin kan nrigal ascii art nta3 text li kaya3tihni user

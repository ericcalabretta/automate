package publisher

import (
	"testing"

	"github.com/chef/automate/api/interservice/cfgmgmt/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCookbookLocks(t *testing.T) {
	data := `{"cookbook_locks":{"apt":{"cache_key":"apt-6.1.4-supermarket.chef.io","dotted_decimal_identifier":"69097958685636418.2160382938276345.8702633799664","identifier":"f57c3a3229774207acdb599d61f907ea3d656ff0","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/apt/versions/6.1.4/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/apt/versions/6.1.4/download","version":"6.1.4"},"version":"6.1.4"},"audit":{"cache_key":"audit-6.0.1-supermarket.chef.io","dotted_decimal_identifier":"62574602498329957.61579351903261380.74868249563275","identifier":"de4f44f9dfcd65dac61823b152c444179f6b908b","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/audit/versions/6.0.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/audit/versions/6.0.1/download","version":"6.0.1"},"version":"6.0.1"},"build-essential":{"cache_key":"build-essential-8.0.4-supermarket.chef.io","dotted_decimal_identifier":"71878984119296372.44510415351338051.249511320474559","identifier":"ff5d8e233f21749e21fca7d25443e2ede19fd3bf","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/build-essential/versions/8.0.4/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/build-essential/versions/8.0.4/download","version":"8.0.4"},"version":"8.0.4"},"chef-client":{"cache_key":"chef-client-10.0.0-supermarket.chef.io","dotted_decimal_identifier":"42728119401384903.37464563401823896.28181711285205","identifier":"97ccff8ad70bc78519d2693b169819a190e3ebd5","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/chef-client/versions/10.0.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/chef-client/versions/10.0.0/download","version":"10.0.0"},"version":"10.0.0"},"chef-ingredient":{"cache_key":"chef-ingredient-2.3.0-supermarket.chef.io","dotted_decimal_identifier":"34436963413247013.21862618200169114.32982718774375","identifier":"7a583c5b63c4254dabef77ae5a9a1dff633c7c67","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/chef-ingredient/versions/2.3.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/chef-ingredient/versions/2.3.0/download","version":"2.3.0"},"version":"2.3.0"},"chef_client_updater":{"cache_key":"chef_client_updater-3.2.5-supermarket.chef.io","dotted_decimal_identifier":"50255799815537903.26883258556648673.124266379849325","identifier":"b28b6237e884ef5f822e64ac8ce17105056dae6d","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/chef_client_updater/versions/3.2.5/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/chef_client_updater/versions/3.2.5/download","version":"3.2.5"},"version":"3.2.5"},"chef_nginx":{"cache_key":"chef_nginx-4.0.2-supermarket.chef.io","dotted_decimal_identifier":"48818168699738677.17228525256677068.82230024656689","identifier":"ad6fdda3eeda353d35409fed96cc4ac9ab68d731","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/chef_nginx/versions/4.0.2/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/chef_nginx/versions/4.0.2/download","version":"4.0.2"},"version":"4.0.2"},"chefops-base":{"cache_key":"chefops-base-0.1.43-ops-supermarket.chef.co","dotted_decimal_identifier":"21234761322130282.25487767504191170.53970546739771","identifier":"4b70e7257eeb6a5a8cfd50347ec23115ff444a3b","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-base/versions/0.1.43/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-base/versions/0.1.43/download","version":"0.1.43"},"version":"0.1.43"},"chefops-devsec-baseline":{"cache_key":"chefops-devsec-baseline-0.1.5-ops-supermarket.chef.co","dotted_decimal_identifier":"32022649319907848.20357232083785930.12862213560956","identifier":"71c46e0d7ae2084852cb72f6d8ca0bb2b7ab9a7c","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-devsec-baseline/versions/0.1.5/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-devsec-baseline/versions/0.1.5/download","version":"0.1.5"},"version":"0.1.5"},"chefops-ldap":{"cache_key":"chefops-ldap-0.2.6-ops-supermarket.chef.co","dotted_decimal_identifier":"41657021068837052.55345156477931790.59750728888216","identifier":"93fed7073580bcc4a02067e7d10e3657ccf86f98","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-ldap/versions/0.2.6/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-ldap/versions/0.2.6/download","version":"0.2.6"},"version":"0.2.6"},"chefops-logging":{"cache_key":"chefops-logging-0.3.2-ops-supermarket.chef.co","dotted_decimal_identifier":"47677922048960846.33739361680770018.140812666001430","identifier":"a962d1475b694e77ddc54bee43e2801180efd816","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-logging/versions/0.3.2/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-logging/versions/0.3.2/download","version":"0.3.2"},"version":"0.3.2"},"chefops-motd":{"cache_key":"chefops-motd-0.1.3-ops-supermarket.chef.co","dotted_decimal_identifier":"8375418052805586.63845634239033550.234323168765881","identifier":"1dc165f9e203d2e2d3441d46f0ced51d9d3ca7b9","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-motd/versions/0.1.3/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-motd/versions/0.1.3/download","version":"0.1.3"},"version":"0.1.3"},"chefops-postfix":{"cache_key":"chefops-postfix-0.6.5-ops-supermarket.chef.co","dotted_decimal_identifier":"18097928305846226.12319470652330970.194769142402726","identifier":"404bf84bd727d22bc47e480d93dab124398bcea6","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-postfix/versions/0.6.5/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-postfix/versions/0.6.5/download","version":"0.6.5"},"version":"0.6.5"},"chefops-sensu":{"cache_key":"chefops-sensu-0.0.1-ops-supermarket.chef.co","dotted_decimal_identifier":"63301604126402386.34955800949112925.188262970755570","identifier":"e0e4793baaa7527c301da269405dab39637f09f2","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-sensu/versions/0.0.1/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefops-sensu/versions/0.0.1/download","version":"0.0.1"},"version":"0.0.1"},"chefopslibs":{"cache_key":"chefopslibs-0.1.20-ops-supermarket.chef.co","dotted_decimal_identifier":"201103598030642.30334468701397139.24895006798592","identifier":"00b6e714860f326bc509df65a09316a452097300","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefopslibs/versions/0.1.20/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/chefopslibs/versions/0.1.20/download","version":"0.1.20"},"version":"0.1.20"},"citadel":{"cache_key":"citadel-1.1.0-supermarket.chef.io","dotted_decimal_identifier":"5008081871689305.449880559491805.231887803404955","identifier":"11cad2ecf7d659019929faf09addd2e69626529b","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/citadel/versions/1.1.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/citadel/versions/1.1.0/download","version":"1.1.0"},"version":"1.1.0"},"compat_resource":{"cache_key":"compat_resource-12.19.1-supermarket.chef.io","dotted_decimal_identifier":"65364324767964307.20312867944317086.187033032973711","identifier":"e83881cec71c93482a721e30c09eaa1b057ddd8f","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/compat_resource/versions/12.19.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/compat_resource/versions/12.19.1/download","version":"12.19.1"},"version":"12.19.1"},"cron":{"cache_key":"cron-5.0.1-supermarket.chef.io","dotted_decimal_identifier":"53840031923642605.43228932742692117.238587404517629","identifier":"bf4739498120ed99947c3c25bd15d8fe757e80fd","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/cron/versions/5.0.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/cron/versions/5.0.1/download","version":"5.0.1"},"version":"5.0.1"},"dpkg_autostart":{"cache_key":"dpkg_autostart-0.2.0-supermarket.chef.io","dotted_decimal_identifier":"57422750648126506.12880729680138903.53751544930405","identifier":"cc01affe684c2a2dc2f49508529730e301bde065","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/dpkg_autostart/versions/0.2.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/dpkg_autostart/versions/0.2.0/download","version":"0.2.0"},"version":"0.2.0"},"erlang":{"cache_key":"erlang-6.0.0-supermarket.chef.io","dotted_decimal_identifier":"18098613459346336.31013177028552669.142022060639452","identifier":"404c97d22be3a06e2e51f824a7dd812b167984dc","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/erlang/versions/6.0.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/erlang/versions/6.0.0/download","version":"6.0.0"},"version":"6.0.0"},"homebrew":{"cache_key":"homebrew-5.0.1-supermarket.chef.io","dotted_decimal_identifier":"63320667854559465.26442073397699001.20988711375986","identifier":"e0f5cfda6958e95df0ecf7fea1b91316d0af9472","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/homebrew/versions/5.0.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/homebrew/versions/5.0.1/download","version":"5.0.1"},"version":"5.0.1"},"iptables":{"cache_key":"iptables-4.3.4-supermarket.chef.io","dotted_decimal_identifier":"68165896295589806.51866186821908848.256232877493232","identifier":"f22c85827ea7aeb84405a95fb970e90adda48bf0","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/iptables/versions/4.3.4/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/iptables/versions/4.3.4/download","version":"4.3.4"},"version":"4.3.4"},"java":{"cache_key":"java-1.50.0-supermarket.chef.io","dotted_decimal_identifier":"43657182654405366.5330763620806707.173709676380713","identifier":"9b19f9f0029af612f04d1ffff0339dfcef8e0229","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/java/versions/1.50.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/java/versions/1.50.0/download","version":"1.50.0"},"version":"1.50.0"},"line":{"cache_key":"line-1.0.5-supermarket.chef.io","dotted_decimal_identifier":"64224323416265828.51485683490319986.28918033111178","identifier":"e42bae8f2a9464b6e9f4d3756e721a4d0117588a","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/line/versions/1.0.5/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/line/versions/1.0.5/download","version":"1.0.5"},"version":"1.0.5"},"logrotate":{"cache_key":"logrotate-2.2.0-supermarket.chef.io","dotted_decimal_identifier":"23609341620057916.54394244012692197.8094668946088","identifier":"53e09234a4f73cc13f46d833d2e5075cafddfaa8","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/logrotate/versions/2.2.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/logrotate/versions/2.2.0/download","version":"2.2.0"},"version":"2.2.0"},"mingw":{"cache_key":"mingw-2.0.2-supermarket.chef.io","dotted_decimal_identifier":"38241076554570451.38145974032463170.277638454423332","identifier":"87dc0e77dd1ad387858fafd0f542fc82bd73af24","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/mingw/versions/2.0.2/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/mingw/versions/2.0.2/download","version":"2.0.2"},"version":"2.0.2"},"ms_dotnet":{"cache_key":"ms_dotnet-4.1.0-supermarket.chef.io","dotted_decimal_identifier":"22363768846886370.58776947537195998.5437473017386","identifier":"4f73bab1f525e2d0d15286efb3de04f202a5ce2a","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/ms_dotnet/versions/4.1.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/ms_dotnet/versions/4.1.0/download","version":"4.1.0"},"version":"4.1.0"},"ntp":{"cache_key":"ntp-3.5.6-supermarket.chef.io","dotted_decimal_identifier":"45717130017822675.11385230976933654.279531648691867","identifier":"a26b7ccedcd7d32872cea810df16fe3b889fae9b","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/ntp/versions/3.5.6/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/ntp/versions/3.5.6/download","version":"3.5.6"},"version":"3.5.6"},"oc-sensu-hec":{"cache_key":"oc-sensu-hec-1.0.87-ops-supermarket.chef.co","dotted_decimal_identifier":"34297714342683056.25985210668573722.95539334885490","identifier":"79d996e800e1b05c51695350641a56e47c1e3072","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/oc-sensu-hec/versions/1.0.87/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/oc-sensu-hec/versions/1.0.87/download","version":"1.0.87"},"version":"1.0.87"},"oc-telegraf":{"cache_key":"oc-telegraf-0.2.18-ops-supermarket.chef.co","dotted_decimal_identifier":"5396536841630756.32734395054433263.118062388425780","identifier":"132c1f25767424744bc2452823ef6b608ab93434","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/oc-telegraf/versions/0.2.18/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/oc-telegraf/versions/0.2.18/download","version":"0.2.18"},"version":"0.2.18"},"ohai":{"cache_key":"ohai-5.2.2-supermarket.chef.io","dotted_decimal_identifier":"3919640474238658.27826495916038513.90942258701595","identifier":"0dece46a1d3ec262dc0cf46b597152b62506651b","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/ohai/versions/5.2.2/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/ohai/versions/5.2.2/download","version":"5.2.2"},"version":"5.2.2"},"openssh":{"cache_key":"openssh-2.6.2-supermarket.chef.io","dotted_decimal_identifier":"64429255145829449.31966919441946334.128615528523907","identifier":"e4e610f21c34497191be6de696de74f9a2e3e883","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/openssh/versions/2.6.2/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/openssh/versions/2.6.2/download","version":"2.6.2"},"version":"2.6.2"},"packagecloud":{"cache_key":"packagecloud-0.3.0-supermarket.chef.io","dotted_decimal_identifier":"69606366194050037.39239743830649352.118909262803009","identifier":"f74a9f0b1a23f58b6856d01a6a086c25b85eb041","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/packagecloud/versions/0.3.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/packagecloud/versions/0.3.0/download","version":"0.3.0"},"version":"0.3.0"},"packages":{"cache_key":"packages-1.1.0-supermarket.chef.io","dotted_decimal_identifier":"34710627810270842.60365598982574374.53723485642327","identifier":"7b5121d08fde7ad676314402912630dc79471e57","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/packages/versions/1.1.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/packages/versions/1.1.0/download","version":"1.1.0"},"version":"1.1.0"},"postfix":{"cache_key":"postfix-5.2.1-supermarket.chef.io","dotted_decimal_identifier":"66838256885556788.17230599240666867.61900922364460","identifier":"ed750a5f11e6343d37238301cef3384c6e7d222c","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/postfix/versions/5.2.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/postfix/versions/5.2.1/download","version":"5.2.1"},"version":"5.2.1"},"push-jobs":{"cache_key":"push-jobs-5.1.2-supermarket.chef.io","dotted_decimal_identifier":"4964240664551452.69090138606413156.254708762964013","identifier":"11a2f35963e01cf5751d711e2564e7a80156702d","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/push-jobs/versions/5.1.2/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/push-jobs/versions/5.1.2/download","version":"5.1.2"},"version":"5.1.2"},"rabbitmq":{"cache_key":"rabbitmq-5.6.1-supermarket.chef.io","dotted_decimal_identifier":"66786054077618627.1768344100421229.130812752374486","identifier":"ed458ff53f19c306484cb1f23e6d76f937a32ad6","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/rabbitmq/versions/5.6.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/rabbitmq/versions/5.6.1/download","version":"5.6.1"},"version":"5.6.1"},"redisio":{"cache_key":"redisio-2.6.1-supermarket.chef.io","dotted_decimal_identifier":"22595729632169036.40775668727664874.46420099991679","identifier":"5046b2462a8c4c90dd4131fdd0ea2a380592087f","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/redisio/versions/2.6.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/redisio/versions/2.6.1/download","version":"2.6.1"},"version":"2.6.1"},"runit":{"cache_key":"runit-4.0.4-supermarket.chef.io","dotted_decimal_identifier":"38185793591614484.22180228398939498.218386101567757","identifier":"87a9c6e67e9c144eccccdac17d6ac69efa02010d","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/runit/versions/4.0.4/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/runit/versions/4.0.4/download","version":"4.0.4"},"version":"4.0.4"},"selinux_policy":{"cache_key":"selinux_policy-2.0.1-supermarket.chef.io","dotted_decimal_identifier":"33041654676373487.24703448856833533.42576899202574","identifier":"756335b9c6bfef57c3a7e5d6adfd26b934f3de0e","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/selinux_policy/versions/2.0.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/selinux_policy/versions/2.0.1/download","version":"2.0.1"},"version":"2.0.1"},"sensu":{"cache_key":"sensu-4.2.1-supermarket.chef.io","dotted_decimal_identifier":"5807310243073541.61136982847355363.39103737607811","identifier":"14a1b7c8238e05d933c3114179e323908c63b683","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/sensu/versions/4.2.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/sensu/versions/4.2.1/download","version":"4.2.1"},"version":"4.2.1"},"seven_zip":{"cache_key":"seven_zip-2.0.2-supermarket.chef.io","dotted_decimal_identifier":"40083337488693902.2896155754419667.104122173006413","identifier":"8e6795446cfa8e0a4a09cef4d5d35eb2d522a24d","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/seven_zip/versions/2.0.2/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/seven_zip/versions/2.0.2/download","version":"2.0.2"},"version":"2.0.2"},"sudo":{"cache_key":"sudo-4.0.1-supermarket.chef.io","dotted_decimal_identifier":"37533232695154660.68939905941550724.259772996848950","identifier":"855846ba1117e4f4ec7aac7eb684ec431d330d36","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/sudo/versions/4.0.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/sudo/versions/4.0.1/download","version":"4.0.1"},"version":"4.0.1"},"sysctl":{"cache_key":"sysctl-1.0.2-supermarket.chef.io","dotted_decimal_identifier":"43339278091677459.69572922010126019.182638964141830","identifier":"99f8d80379ef13f72c34366646c3a61bf2946f06","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/sysctl/versions/1.0.2/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/sysctl/versions/1.0.2/download","version":"1.0.2"},"version":"1.0.2"},"uchiwa":{"cache_key":"uchiwa-1.4.0-supermarket.chef.io","dotted_decimal_identifier":"7408943055752505.4849356687497769.273720442503259","identifier":"1a5264faffcd39113a76d6bf5a29f8f281be405b","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/uchiwa/versions/1.4.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/uchiwa/versions/1.4.0/download","version":"1.4.0"},"version":"1.4.0"},"ulimit":{"cache_key":"ulimit-1.0.0-supermarket.chef.io","dotted_decimal_identifier":"18593180655244934.15550619869400039.39488407025277","identifier":"420e6638335686373f34d55a4be723ea1c79367d","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/ulimit/versions/1.0.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/ulimit/versions/1.0.0/download","version":"1.0.0"},"version":"1.0.0"},"users":{"cache_key":"users-5.3.1-supermarket.chef.io","dotted_decimal_identifier":"18510001311982710.52844924323313198.113415656458183","identifier":"41c2bf84fe2476bbbe2dbe79362e6726a3ea47c7","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/users/versions/5.3.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/users/versions/5.3.1/download","version":"5.3.1"},"version":"5.3.1"},"vector":{"cache_key":"vector-0.1.8-ops-supermarket.chef.co","dotted_decimal_identifier":"349788557239165.69236474149017072.167415818247887","identifier":"013e217eb3637df5fa34d72715f0984388605ecf","origin":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/vector/versions/0.1.8/download","source_options":{"artifactserver":"https://ops-supermarket.chef.co:443/api/v1/cookbooks/vector/versions/0.1.8/download","version":"0.1.8"},"version":"0.1.8"},"windows":{"cache_key":"windows-4.0.1-supermarket.chef.io","dotted_decimal_identifier":"60570365749581177.23554430975858859.197368287308218","identifier":"d7306d3e6b557953aea15333d8abb38162a7a9ba","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/windows/versions/4.0.1/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/windows/versions/4.0.1/download","version":"4.0.1"},"version":"4.0.1"},"yum":{"cache_key":"yum-5.1.0-supermarket.chef.io","dotted_decimal_identifier":"31624909165547594.72046938166162316.279984579186306","identifier":"705aaff4ddd04afff64efc8c738cfea4fd628e82","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/yum/versions/5.1.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/yum/versions/5.1.0/download","version":"5.1.0"},"version":"5.1.0"},"yum-epel":{"cache_key":"yum-epel-3.1.0-supermarket.chef.io","dotted_decimal_identifier":"61825876588221950.7210614985696931.222844717230431","identifier":"dba64ea3f7d5fe199e0420d5a6a3caad1435e95f","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/yum-epel/versions/3.1.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/yum-epel/versions/3.1.0/download","version":"3.1.0"},"version":"3.1.0"},"yum-erlang_solutions":{"cache_key":"yum-erlang_solutions-2.0.0-supermarket.chef.io","dotted_decimal_identifier":"25480554343062123.14403248363945810.236905800156222","identifier":"5a866dde4c5a6b332bad96583f52d776ee10443e","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/yum-erlang_solutions/versions/2.0.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/yum-erlang_solutions/versions/2.0.0/download","version":"2.0.0"},"version":"2.0.0"},"zypper":{"cache_key":"zypper-0.4.0-supermarket.chef.io","dotted_decimal_identifier":"11302516343739339.12053819794478172.138287915714268","identifier":"28279427bc93cb2ad2e29f09545c7dc5aa17c6dc","origin":"https://supermarket.chef.io:443/api/v1/cookbooks/zypper/versions/0.4.0/download","source_options":{"artifactserver":"https://supermarket.chef.io:443/api/v1/cookbooks/zypper/versions/0.4.0/download","version":"0.4.0"},"version":"0.4.0"}},"default_attributes":{"audit":{"fetcher":"chef-server-automate","profiles":[{"compliance":"builder/cd-infrastructure-base","name":"cd-infrastructure-base"}],"reporter":"chef-server-automate"},"authorization":{"sudo":{"include_sudoers_d":true,"passwordless":true,"users":["centos","ec2-user","fallback","ubuntu"]}},"chef-apt-docker":{"uri":"http://artifactory.chef.co/docker-apt-remote/ubuntu"},"chef_client":{"bin":"/opt/chef/bin/chef-client","config":{"chef_server_url":"https://chef-server.chef.co/organizations/chef"}},"chef_client_updater":{"channel":"current","post_install_action":"exec","prevent_downgrade":true},"chefops-base":{"enable_chef_client_updater":true,"enable_email":true,"enable_monitoring":true,"enable_push_jobs":false,"use_ldap":true},"chefops-ldap":{"user_filter":"(|(memberOf=CN=Entire Company,OU=Okta Security Groups,OU=Domain Users,DC=opscodecorp,DC=com)(cn=Chef CI))"},"citadel":{"bucket":"chef-cd-citadel","region":"us-west-2"},"sensu":{"apt_repo_url":"http://artifactory.chef.co/sensu-apt-remote/","client":{"pager_team":"acc"},"subscriptions":["acc","chef-cd"]},"telegraf":{"apt_repo_url":"http://artifactory.chef.co/influxdata-apt-remote/ubuntu/"},"vector":{"pcp":{"apt_repo_url":"http://artifactory.chef.co/performance-co-pilot-apt-remote/xenial/"}}},"name":"cd-infrastructure-base","override_attributes":{},"revision_id":"e775abb4fa02004c5d0cbb7f73f082e09ebf77f1a613a6c2cbde646735b68a4c","run_list":["recipe[chefops-base::default]","recipe[audit::default]"],"solution_dependencies":{"Policyfile":[["chefops-base","= 0.1.43"],["audit","= 6.0.1"],["apt","= 6.1.4"],["build-essential","= 8.0.4"],["chef-client","= 10.0.0"],["chefops-ldap","= 0.2.6"],["chefops-logging","= 0.3.2"],["chefops-motd","= 0.1.3"],["chefops-postfix","= 0.6.5"],["chefopslibs","= 0.1.20"],["chef_client_updater","= 3.2.5"],["ntp","= 3.5.6"],["oc-sensu-hec","= 1.0.87"],["oc-telegraf","= 0.2.18"],["packages","= 1.1.0"],["push-jobs","= 5.1.2"],["sudo","= 4.0.1"],["users","= 5.3.1"],["vector","= 0.1.8"],["chefops-devsec-baseline","= 0.1.5"],["chefops-sensu","= 0.0.1"],["sensu","= 4.2.1"],["uchiwa","= 1.4.0"],["sysctl","= 1.0.2"],["postfix","= 5.2.1"],["citadel","= 1.1.0"],["yum","= 5.1.0"],["line","= 1.0.5"],["openssh","= 2.6.2"],["runit","= 4.0.4"],["chef-ingredient","= 2.3.0"],["iptables","= 4.3.4"],["ohai","= 5.2.2"],["seven_zip","= 2.0.2"],["mingw","= 2.0.2"],["windows","= 4.0.1"],["compat_resource","= 12.19.1"],["packagecloud","= 0.3.0"],["yum-epel","= 3.1.0"],["ms_dotnet","= 4.1.0"],["rabbitmq","= 5.6.1"],["redisio","= 2.6.1"],["ulimit","= 1.0.0"],["selinux_policy","= 2.0.1"],["erlang","= 6.0.0"],["yum-erlang_solutions","= 2.0.0"],["dpkg_autostart","= 0.2.0"],["logrotate","= 2.2.0"],["chef_nginx","= 4.0.2"],["java","= 1.50.0"],["zypper","= 0.4.0"],["homebrew","= 5.0.1"],["cron","= 5.0.1"]],"dependencies":{"apt (6.1.4)":[],"audit (6.0.1)":[["compat_resource","\u003e= 0.0.0"]],"build-essential (8.0.4)":[["seven_zip","\u003e= 0.0.0"],["mingw","\u003e= 1.1.0"]],"chef-client (10.0.0)":[["cron","\u003e= 4.2.0"],["logrotate","\u003e= 1.9.0"],["windows","\u003e= 2.0.0"]],"chef-ingredient (2.3.0)":[],"chef_client_updater (3.2.5)":[],"chef_nginx (4.0.2)":[["build-essential","\u003e= 0.0.0"],["ohai","\u003e= 4.1.0"],["yum-epel","\u003e= 0.0.0"],["runit","\u003e= 1.6.0"],["compat_resource","\u003e= 12.14.6"],["zypper","\u003e= 0.0.0"]],"chefops-base (0.1.43)":[["apt","\u003e= 0.0.0"],["build-essential","\u003e= 0.0.0"],["chef-client","\u003e= 0.0.0"],["chefops-ldap","\u003e= 0.0.0"],["chefops-logging","\u003e= 0.0.0"],["chefops-motd","\u003e= 0.0.0"],["chefops-postfix","\u003e= 0.0.0"],["chefopslibs","\u003e= 0.0.0"],["chef_client_updater","\u003e= 0.0.0"],["ntp","\u003e= 0.0.0"],["oc-sensu-hec","\u003e= 0.0.0"],["oc-telegraf","\u003e= 0.0.0"],["packages","\u003e= 0.0.0"],["push-jobs","\u003e= 0.0.0"],["sudo","\u003e= 0.0.0"],["users","\u003e= 0.0.0"],["vector","\u003e= 0.0.0"],["audit","\u003e= 0.0.0"],["chefops-devsec-baseline","\u003e= 0.0.0"],["chefops-sensu","\u003e= 0.0.0"]],"chefops-devsec-baseline (0.1.5)":[["sysctl","\u003e= 0.0.0"]],"chefops-ldap (0.2.6)":[["line","\u003e= 0.0.0"],["openssh","\u003e= 0.0.0"],["citadel","\u003e= 0.0.0"]],"chefops-logging (0.3.2)":[["apt","\u003e= 0.0.0"],["yum","\u003e= 0.0.0"],["chef_nginx","\u003c 5.0.0"],["java","\u003e= 0.0.0"],["citadel","\u003e= 0.0.0"],["chefopslibs","\u003e= 0.0.0"]],"chefops-motd (0.1.3)":[],"chefops-postfix (0.6.5)":[["postfix","\u003e= 0.0.0"],["citadel","\u003e= 0.0.0"],["chefopslibs","\u003e= 0.0.0"]],"chefops-sensu (0.0.1)":[["sensu","\u003e= 0.0.0"],["uchiwa","\u003e= 0.0.0"],["chefopslibs","\u003e= 0.0.0"],["build-essential","\u003e= 0.0.0"]],"chefopslibs (0.1.20)":[["citadel","\u003e= 0.0.0"]],"citadel (1.1.0)":[],"compat_resource (12.19.1)":[],"cron (5.0.1)":[],"dpkg_autostart (0.2.0)":[],"erlang (6.0.0)":[["yum-epel","\u003e= 0.0.0"],["yum-erlang_solutions","\u003e= 1.0.3"],["build-essential","\u003e= 0.0.0"]],"homebrew (5.0.1)":[],"iptables (4.3.4)":[],"java (1.50.0)":[["apt","\u003e= 0.0.0"],["windows","\u003e= 0.0.0"],["homebrew","\u003e= 0.0.0"]],"line (1.0.5)":[],"logrotate (2.2.0)":[],"mingw (2.0.2)":[["seven_zip","\u003e= 0.0.0"]],"ms_dotnet (4.1.0)":[["windows","\u003e= 2.1.0"]],"ntp (3.5.6)":[],"oc-sensu-hec (1.0.87)":[["sensu","\u003e= 0.0.0"],["uchiwa","\u003e= 0.0.0"],["chefopslibs","\u003e= 0.0.0"],["build-essential","\u003e= 0.0.0"]],"oc-telegraf (0.2.18)":[["apt","\u003e= 0.0.0"],["yum","\u003e= 0.0.0"],["chefopslibs","\u003e= 0.0.0"]],"ohai (5.2.2)":[],"openssh (2.6.2)":[["iptables","\u003e= 1.0.0"]],"packagecloud (0.3.0)":[],"packages (1.1.0)":[],"postfix (5.2.1)":[],"push-jobs (5.1.2)":[["runit","\u003e= 1.2.0"],["chef-ingredient","\u003e= 0.0.0"]],"rabbitmq (5.6.1)":[["erlang","\u003e= 0.0.0"],["yum-epel","\u003e= 0.0.0"],["yum-erlang_solutions","\u003e= 0.0.0"],["dpkg_autostart","\u003e= 0.0.0"],["logrotate","\u003e= 0.0.0"]],"redisio (2.6.1)":[["ulimit","\u003e= 0.1.2"],["build-essential","\u003e= 0.0.0"],["selinux_policy","\u003e= 0.0.0"]],"runit (4.0.4)":[["packagecloud","\u003e= 0.0.0"],["yum-epel","\u003e= 0.0.0"]],"selinux_policy (2.0.1)":[["compat_resource","\u003e= 12.16.3"]],"sensu (4.2.1)":[["apt","\u003e= 2.0.0"],["yum","\u003e= 3.0.0"],["windows","\u003e= 1.36.0"],["ms_dotnet","\u003e= 2.6.1"],["rabbitmq","\u003e= 2.0.0"],["redisio","\u003e= 1.7.0"]],"seven_zip (2.0.2)":[["windows","\u003e= 1.2.2"]],"sudo (4.0.1)":[],"sysctl (1.0.2)":[["ohai","\u003e= 5.0.0"]],"uchiwa (1.4.0)":[["yum","\u003e= 0.0.0"],["apt","\u003e= 0.0.0"]],"ulimit (1.0.0)":[],"users (5.3.1)":[],"vector (0.1.8)":[["build-essential","\u003e= 0.0.0"],["apt","\u003e= 0.0.0"],["yum","\u003e= 0.0.0"]],"windows (4.0.1)":[],"yum (5.1.0)":[],"yum-epel (3.1.0)":[],"yum-erlang_solutions (2.0.0)":[["yum-epel","\u003e= 0.0.0"]],"zypper (0.4.0)":[]}}}`

	expectedCookbooks := []*request.PolicyCookbookLock{
		{CookbookName: "chefopslibs", PolicyId: "00b6e714860f326bc509df65a09316a452097300"},
		{CookbookName: "citadel", PolicyId: "11cad2ecf7d659019929faf09addd2e69626529b"},
		{CookbookName: "packagecloud", PolicyId: "f74a9f0b1a23f58b6856d01a6a086c25b85eb041"},
		{CookbookName: "push-jobs", PolicyId: "11a2f35963e01cf5751d711e2564e7a80156702d"},
		{CookbookName: "sensu", PolicyId: "14a1b7c8238e05d933c3114179e323908c63b683"},
		{CookbookName: "compat_resource", PolicyId: "e83881cec71c93482a721e30c09eaa1b057ddd8f"},
		{CookbookName: "java", PolicyId: "9b19f9f0029af612f04d1ffff0339dfcef8e0229"},
		{CookbookName: "yum-epel", PolicyId: "dba64ea3f7d5fe199e0420d5a6a3caad1435e95f"},
		{CookbookName: "apt", PolicyId: "f57c3a3229774207acdb599d61f907ea3d656ff0"},
		{CookbookName: "build-essential", PolicyId: "ff5d8e233f21749e21fca7d25443e2ede19fd3bf"},
		{CookbookName: "logrotate", PolicyId: "53e09234a4f73cc13f46d833d2e5075cafddfaa8"},
		{CookbookName: "yum-erlang_solutions", PolicyId: "5a866dde4c5a6b332bad96583f52d776ee10443e"},
		{CookbookName: "postfix", PolicyId: "ed750a5f11e6343d37238301cef3384c6e7d222c"},
		{CookbookName: "selinux_policy", PolicyId: "756335b9c6bfef57c3a7e5d6adfd26b934f3de0e"},
		{CookbookName: "chefops-devsec-baseline", PolicyId: "71c46e0d7ae2084852cb72f6d8ca0bb2b7ab9a7c"},
		{CookbookName: "erlang", PolicyId: "404c97d22be3a06e2e51f824a7dd812b167984dc"},
		{CookbookName: "runit", PolicyId: "87a9c6e67e9c144eccccdac17d6ac69efa02010d"},
		{CookbookName: "mingw", PolicyId: "87dc0e77dd1ad387858fafd0f542fc82bd73af24"},
		{CookbookName: "packages", PolicyId: "7b5121d08fde7ad676314402912630dc79471e57"},
		{CookbookName: "chefops-base", PolicyId: "4b70e7257eeb6a5a8cfd50347ec23115ff444a3b"},
		{CookbookName: "chefops-motd", PolicyId: "1dc165f9e203d2e2d3441d46f0ced51d9d3ca7b9"},
		{CookbookName: "iptables", PolicyId: "f22c85827ea7aeb84405a95fb970e90adda48bf0"},
		{CookbookName: "chef_client_updater", PolicyId: "b28b6237e884ef5f822e64ac8ce17105056dae6d"},
		{CookbookName: "chef_nginx", PolicyId: "ad6fdda3eeda353d35409fed96cc4ac9ab68d731"},
		{CookbookName: "cron", PolicyId: "bf4739498120ed99947c3c25bd15d8fe757e80fd"},
		{CookbookName: "oc-telegraf", PolicyId: "132c1f25767424744bc2452823ef6b608ab93434"},
		{CookbookName: "users", PolicyId: "41c2bf84fe2476bbbe2dbe79362e6726a3ea47c7"},
		{CookbookName: "chef-ingredient", PolicyId: "7a583c5b63c4254dabef77ae5a9a1dff633c7c67"},
		{CookbookName: "chefops-sensu", PolicyId: "e0e4793baaa7527c301da269405dab39637f09f2"},
		{CookbookName: "ms_dotnet", PolicyId: "4f73bab1f525e2d0d15286efb3de04f202a5ce2a"},
		{CookbookName: "ulimit", PolicyId: "420e6638335686373f34d55a4be723ea1c79367d"},
		{CookbookName: "zypper", PolicyId: "28279427bc93cb2ad2e29f09545c7dc5aa17c6dc"},
		{CookbookName: "redisio", PolicyId: "5046b2462a8c4c90dd4131fdd0ea2a380592087f"},
		{CookbookName: "seven_zip", PolicyId: "8e6795446cfa8e0a4a09cef4d5d35eb2d522a24d"},
		{CookbookName: "chefops-ldap", PolicyId: "93fed7073580bcc4a02067e7d10e3657ccf86f98"},
		{CookbookName: "ntp", PolicyId: "a26b7ccedcd7d32872cea810df16fe3b889fae9b"},
		{CookbookName: "oc-sensu-hec", PolicyId: "79d996e800e1b05c51695350641a56e47c1e3072"},
		{CookbookName: "sysctl", PolicyId: "99f8d80379ef13f72c34366646c3a61bf2946f06"},
		{CookbookName: "audit", PolicyId: "de4f44f9dfcd65dac61823b152c444179f6b908b"},
		{CookbookName: "chef-client", PolicyId: "97ccff8ad70bc78519d2693b169819a190e3ebd5"},
		{CookbookName: "dpkg_autostart", PolicyId: "cc01affe684c2a2dc2f49508529730e301bde065"},
		{CookbookName: "rabbitmq", PolicyId: "ed458ff53f19c306484cb1f23e6d76f937a32ad6"},
		{CookbookName: "homebrew", PolicyId: "e0f5cfda6958e95df0ecf7fea1b91316d0af9472"},
		{CookbookName: "yum", PolicyId: "705aaff4ddd04afff64efc8c738cfea4fd628e82"},
		{CookbookName: "sudo", PolicyId: "855846ba1117e4f4ec7aac7eb684ec431d330d36"},
		{CookbookName: "vector", PolicyId: "013e217eb3637df5fa34d72715f0984388605ecf"},
		{CookbookName: "chefops-logging", PolicyId: "a962d1475b694e77ddc54bee43e2801180efd816"},
		{CookbookName: "chefops-postfix", PolicyId: "404bf84bd727d22bc47e480d93dab124398bcea6"},
		{CookbookName: "ohai", PolicyId: "0dece46a1d3ec262dc0cf46b597152b62506651b"},
		{CookbookName: "openssh", PolicyId: "e4e610f21c34497191be6de696de74f9a2e3e883"},
		{CookbookName: "line", PolicyId: "e42bae8f2a9464b6e9f4d3756e721a4d0117588a"},
		{CookbookName: "windows", PolicyId: "d7306d3e6b557953aea15333d8abb38162a7a9ba"},
		{CookbookName: "uchiwa", PolicyId: "1a5264faffcd39113a76d6bf5a29f8f281be405b"},
	}
	policyCookbooks, err := getCookbookLocks(data)
	require.NoError(t, err)

	assert.ElementsMatch(t, expectedCookbooks, policyCookbooks)
}

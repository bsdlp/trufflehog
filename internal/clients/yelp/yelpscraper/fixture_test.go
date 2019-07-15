package yelpscraper

const peterLugerYelpPage = `<!DOCTYPE HTML>

<!--[if lt IE 7 ]> <html xmlns:fb="http://www.facebook.com/2008/fbml" class="ie6 ie ltie9 ltie8 no-js" lang="en"> <![endif]-->
<!--[if IE 7 ]>    <html xmlns:fb="http://www.facebook.com/2008/fbml" class="ie7 ie ltie9 ltie8 no-js" lang="en"> <![endif]-->
<!--[if IE 8 ]>    <html xmlns:fb="http://www.facebook.com/2008/fbml" class="ie8 ie ltie9 no-js" lang="en"> <![endif]-->
<!--[if IE 9 ]>    <html xmlns:fb="http://www.facebook.com/2008/fbml" class="ie9 ie no-js" lang="en"> <![endif]-->
<!--[if (gt IE 9)|!(IE)]><!--> <html xmlns:fb="http://www.facebook.com/2008/fbml" class="no-js" lang="en"> <!--<![endif]-->
    <head>
        <script>
            (function() {
                var main = null;

                var main=function(){window.onerror=function(k,a,c,i,f){var j=(document.getElementsByTagName("html")[0].getAttribute("webdriver")==="true"||navigator.userAgent==="selenium");var h=f&&(f.name==="ServerSideRenderingError"||f.name==="CSRFallbackError");if(j&&!h){document.body.innerHTML="<h1>Javascript Error Detected</h1>";var g=document.createElement("div");g.setAttribute("id","pastebinTraceback");var d=document.createElement("code");var e={message:k,source:a,lineno:c,colno:i,error:f};var b=document.createTextNode("JS ERROR DETECTED\n"+window.JSON.stringify(e,null,2));
d.appendChild(b);g.appendChild(d);document.getElementsByTagName("body")[0].appendChild(g)}}};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main();
            })();
    </script>


    <script>            window.yPageStart = new Date().getTime();
</script>

    <script>            var initialVisibilityState = document.webkitVisibilityState;

                yPerfTimings = [];

                ySitRepParams = {"clientIP": "136.25.5.36", "datacenter": "us-west-1", "is_internal_ip": false, "edgeStartTime": 1563162272772975, "cfRayID": null, "site": "main", "serverStartTime": 1563162272789, "action": null, "edgestageSubdomain": "_disabled", "disableBeacon": false, "b3Sampled": "0", "uniqueRequestID": "5e84cbeca1806274", "isLoggedIn": false, "zipkinTraceID": "accfaea8d58a079b", "servlet": "biz_details", "yuv_record": "uUzVE3xrF4-y5Xh09NWX72vDZq0Y5FEIzIpDChwe-KAfRBHEunJxhN313wzLMns9t0smHRLWWB6mvgf1k_xN-NWoQ-sL8AA2"};
                window.ySitRepParams['initialVisibilityState'] = initialVisibilityState;
                window.ySitRepParams['seoCohorts'] = [["traffic_pagination_seo_experiment", "indexed_pagination_1"], ["biz_site_new_business_user_onboarding", "status_quo"], ["traffic_desc_search_title_last_updated_experiment", "with_last_updated_1"], ["traffic_search_url_normalization_experiment_v3", "active_01"], ["biz_site_self_serve_cpc_re_v2", "00_status_quo"], ["contributions.biz_owner_respond_to_new_review_email", "enabled"], ["traffic_nearme_optimization_experiment", "treatment"], ["biz_conditional_raq_title_experiment", "raq_ruleset2_01"], ["traffic_internal_linking_biz_details_to_raq_searches", "enabled"], ["traffic_main_biz_experiment_toggle", null], ["www_boi_modal_callouts_v2", "claim_biz_for_free"], ["traffic_search_internal_linking_service_offerings", "linked_0"], ["traffic_search_internal_linking_top_searches", "linked_1"], ["biz_site_yes_it_works_rollout", "yiw_impression_and_clicks_enabled"], ["biz_growth.biz_site_verify_with_work_email_v2", "enabled"], ["biz_growth.category_specific_claim_callouts", "status_quo"], ["biz_growth.boi_liveramp_holdout_experiment", "ads_disabled"], ["traffic_search_internal_linking_experiment_v2", "linked_0"], ["traffic_index_new_pages_cross_validation", "cv_group"], ["traffic_search_internal_linking_home_cleaning", "linked_1"], ["traffic_top_searches_for_sitemap", "status_quo_1"], ["www_biz_owner_identification", "v4"], ["traffic_critical_css_on_msite", "aggressive_1"], ["traffic_search_url_normalization", "enabled"], ["traffic_remove_cdn_sharding_exp", "remove_cdn_sharding_0"], ["yelp.www.search.traffic_bunsen_calibration_experiment", "NA_cohort_allocation_comes_from_bunsen"], ["traffic_nearme_smart_banner_experiment", "smart_banner"], ["traffic_top_searches_through_cordite", "disabled"], ["amp_biz_details", "status_quo"], ["www_boi_holdout_experiment_2019_q2", "prompts_enabled"], ["traffic_category_city_category_browse_experiment", "category_city_0"], ["traffic_remove_canonical_links_on_paginated_pages", "no_canonical_2"], ["traffic_nearme_redesign_experiment", "redesign_0"], ["traffic_canonicalize_category_in_find_desc_to_cflt_experiment", "status_quo_01"], ["traffic_biz_details_title_v6_transactions_v3", "conditions_1"]];


            (function(H){H.className=H.className.replace(/\bno-js\b/,'js');})(document.documentElement);
</script>

    <script>            pageSpeedCustomTimestamps = {};
</script>

            <script>            (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
</script>



        <script>
            (function() {
                var main = null;

                !function(e){function t(n){if(i[n])return i[n].exports;var r=i[n]={i:n,l:!1,exports:{}};return e[n].call(r.exports,r,r.exports,t),r.l=!0,r.exports}var i={};t.m=e,t.c=i,t.d=function(e,i,n){t.o(e,i)||Object.defineProperty(e,i,{configurable:!1,enumerable:!0,get:n})},t.n=function(e){var i=e&&e.__esModule?function(){return e.default}:function(){return e};return t.d(i,"a",i),i},t.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},t.p="",t(t.s=5)}([function(e,t,i){"use strict";var n=function(e,t){for(var i=0;i<e.length;i+=1)if(e[i]===t)return i;return-1};i.d(t,"c",function(){return r}),i.d(t,"b",function(){return o}),i.d(t,"a",function(){return s});var r="global",o="biz",s="biz_gtag",a=[r,"m","www",o,"webview","api","admin","error_pages",s],c=function(e,t){if(n(a,t)>=0)return[t,e].join(".");throw new Error("google analytics attempted to set "+e+" to an unrecognized tracker alias: "+t)};t.d=c},function(e,t,i){"use strict";function n(e){for(var t=1;t<arguments.length;t++){var i=null!=arguments[t]?arguments[t]:{},n=Object.keys(i);"function"==typeof Object.getOwnPropertySymbols&&(n=n.concat(Object.getOwnPropertySymbols(i).filter(function(e){return Object.getOwnPropertyDescriptor(i,e).enumerable}))),n.forEach(function(t){r(e,t,i[t])})}return e}function r(e,t,i){return t in e?Object.defineProperty(e,t,{value:i,enumerable:!0,configurable:!0,writable:!0}):e[t]=i,e}var o=i(2),s=i(3),a=i.n(s),c=function(e){var t="240+";return e<=0?t="5":e<=30?t=(5*Math.ceil(e/5)).toString():e<=120?t=(15*Math.ceil(e/15)).toString():e<=240&&(t=(30*Math.ceil(e/30)).toString()),t},u=c,l=function(){return parseInt(((new Date).getTime()-window.ygaPageStartTime)/1e3,10)},f=l,h=i(0),d=function(e,t){window.ga(Object(h.d)("send",e),t)},m=d;Object.keys||a.a.shim(),Object.entries||(Object.entries=function(e){for(var t=Object.keys(e),i=new Array(t.length),n=0;n<t.length;n+=1)i[n]=[t[n],e[t[n]]];return i});var g=function(){function e(e){r(this,"clientID",void 0),r(this,"dimensions",void 0),r(this,"domain",void 0),r(this,"friendlyMap",void 0),r(this,"jsDimensions",void 0),r(this,"metrics",void 0),r(this,"temporaryDimensions",void 0),r(this,"temporaryMetrics",void 0),r(this,"trackers",void 0),r(this,"userID",void 0),r(this,"enableHighVolumeEvents",void 0),window.ga||(window.ga=function(){}),window.gtag||(window.gtag=function(){}),this.reload(e)}var t=e.prototype;return t.setupTrackers=function(){var e=this;this.friendlyMap={},this.temporaryDimensions={},this.temporaryMetrics={},Object.entries(this.trackers).forEach(function(t){var i=t[0],n=t[1];e.setupTracker(String(i),String(n))})},t.setupTracker=function(e,t){var i=this,r=n({cookieDomain:this.domain,name:e,clientId:this.clientID},this.userID?{userId:this.userID}:{});window.ga("create",t,r),this.friendlyMap[e]={},this.temporaryDimensions[e]={},this.temporaryMetrics[e]={},this.addDimensionsToFriendlyMap(e,this.dimensions),this.addDimensionsToFriendlyMap(e,this.jsDimensions),this.dimensions[e]&&Object.keys(this.dimensions[e]).forEach(function(t){i.setDimension(String(t),i.dimensions[e][t][1])}),this.metrics[e]&&Object.keys(this.metrics[e]).forEach(function(t){i.setMetric(i.metrics[e][t][0],i.metrics[e][t][1],e)}),window.ga(Object(h.d)("set",e),"anonymizeIp",!0)},t.addDimensionsToFriendlyMap=function(e,t){var i=this;t&&t[e]&&Object.keys(t[e]).forEach(function(n){i.friendlyMap[e][n]=t[e][n][0]})},t.firePageviews=function(e){var t=this;Object.keys(this.trackers).forEach(function(i){t.firePageview(String(i),e)})},t.firePageviewsWithGlobalTrackerSampled=function(e){var t=this;Object.keys(this.trackers).forEach(function(i){(i!==h.c||t.enableHighVolumeEvents)&&t.firePageview(String(i),e)})},t.firePageview=function(e,t){if(e===h.a){var i=this.trackers[e],r=n({},this.dimensions[e],this.jsDimensions[e]),o=Object.values(r).reduce(function(e,t){var i,r=t[0],o=t[1];return n({},e,(i={},i["dimension"+r]=null===o?o:o.toString(),i))},{}),s=Object.values(this.metrics[e]||{}).reduce(function(e,t){var i,r=t[0],o=t[1];return n({},e,(i={},i["metric"+r]=o,i))},{});window.gtag("event","page_view",n({send_to:i,anonymize_ip:!0,cookie_domain:this.domain,client_id:this.clientID},o,s,this.userID?{user_id:this.userID}:{},t?{page_path:t}:{}))}else t&&window.ga(Object(h.d)("set",e),"page",t),window.ga(Object(h.d)("send",e),"pageview")},t.setDimensionOneEvent=function(e,t){var i=this;Object.keys(this.friendlyMap).forEach(function(n){var r=i.friendlyMap[n][e];void 0!==r&&null!==t&&(i.temporaryDimensions[n]["dimension"+r]=t.toString())})},t.setDimension=function(e,t){var i=this;Object.keys(this.friendlyMap).forEach(function(n){var r=i.friendlyMap[n][e];void 0!==r&&null!==t&&window.ga(Object(h.d)("set",n),"dimension"+r,t.toString())})},t.setDimensions=function(e){var t=this;Object.keys(e).forEach(function(i){t.setDimension(i,e[i])})},t.setMetric=function(e,t,i){window.ga(Object(h.d)("set",i),"metric"+e,t)},t.setMetricOneEvent=function(e,t,i){this.temporaryMetrics[i]["metric"+e]=t},t.trackEvent=function(e,t,i,r,o){var s=this;Object.keys(this.trackers).forEach(function(a){if(a===h.a){var c=s.trackers[a];window.gtag("event",t,n({send_to:c,event_category:e,non_interaction:!0},s.temporaryDimensions[a],s.temporaryMetrics[a],i&&{event_label:i},r&&{value:r},o&&{event_callback:o}))}else if(a!==h.c){var u=n({hitType:"event",eventCategory:e,eventAction:t,eventLabel:i,eventValue:r},o&&{hitCallback:o},{nonInteraction:!0},s.temporaryDimensions[a],s.temporaryMetrics[a]);m(String(a),u)}s.temporaryDimensions[a]={},s.temporaryMetrics[a]={}})},t.trackTiming=function(e,t,i,n){var r={hitType:"timing",timingCategory:e,timingVar:t,timingValue:i,timingLabel:n};Object.keys(this.trackers).forEach(function(e){m(String(e),r)})},t.trackEventHighVolume=function(e,t,i,n,r){if(this.enableHighVolumeEvents){var o=e+" / 10";this.trackEvent(o,t,i,n,r)}},t.trackEventWithTime=function(e,t,i){this.trackEvent(e,t,i,f())},t.initTimeOnPageEvent=function(){var e=this;Object(o.a)(window).on("beforeunload",function(){var t=f(),i=u(t);e.trackEventHighVolume("time on page","unload",i,t)})},t.initDwellTimeEvent=function(){var e=this;setTimeout(function(){e.trackEvent("dwell time","dwell","30 seconds")},3e4)},t.applyConfig=function(e){this.trackers=e.trackers,this.domain=e.domain,this.clientID=e.clientID,this.userID=e.user_id,this.dimensions=e.dimensions,this.metrics=e.metrics,this.jsDimensions=e.js_dimensions,this.enableHighVolumeEvents=e.enable_high_volume_events},t.reload=function(e){this.applyConfig(e),this.setupTrackers()},t.getGaConfig=function(){return{trackers:this.trackers,domain:this.domain,clientID:this.clientID,user_id:this.userID,dimensions:this.dimensions,metrics:this.metrics,js_dimensions:this.jsDimensions,enable_high_volume_events:this.enableHighVolumeEvents}},e.init=function(t){e.instance=new e(t)},e.getInstance=function(){var t=e.instance;if(null==t)throw Error("yelp_google_analytics.GoogleAnalytics not initialized. Call a site-specific GA init!");return t},e}();r(g,"instance",void 0);t.a=g},function(e,t,i){"use strict";t.a=function(){var e;return(e=window).jQuery.apply(e,arguments)}},function(e,t,i){"use strict";!function(){Object.keys||(Object.keys=function(){var e=Object.prototype.hasOwnProperty,t=!{toString:null}.propertyIsEnumerable("toString"),i=["toString","toLocaleString","valueOf","hasOwnProperty","isPrototypeOf","propertyIsEnumerable","constructor"],n=i.length;return function(r){if("object"!=typeof r&&("function"!=typeof r||null===r))throw new TypeError("Object.keys called on non-object");var o,s,a=[];for(o in r)e.call(r,o)&&a.push(o);if(t)for(s=0;s<n;s++)e.call(r,i[s])&&a.push(i[s]);return a}}())}()},,function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var n=i(1),r=function(e,t){n.a.init(e),n.a.getInstance().firePageviews(t),n.a.getInstance().initDwellTimeEvent()},o=r;window.yelp_google_analytics={google_analytics:n.a},window.ygaPageStartTime=(new Date).getTime(),main=o}]);

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main({"metrics": {"www": {"mtb-reply-rate-shown": [4, null], "mtb-response-time-shown": [1, null]}, "global": {}}, "domain": "yelp.com", "user_id": null, "dimensions": {"www": {"www_second_page_pitch": [111, "status_quo"], "rating": [114, 4.0], "www_education_banner": [190, "enabled"], "mtb_mapbox_link": [168, "False"], "messaging.www.yelp_guaranteed": [160, "status_quo"], "traffic_m_biz_parity_mobile_first": [31, "inactive_disabled"], "traffic_critical_css_on_msite": [184, "aggressive_1"], "show_holiday_banner": [105, "NO_HOLIDAY"], "page_depth": [41, 0], "messaging.www.composer_notifications_spam_and_hover_state": [177, "status_quo"], "biz_num_photos": [9, "100+"], "txn.www.checkout_page_changes": [135, "enabled"], "business_id": [141, "4yPqqJDJOQX69gC66YUDkA"], "biz_has_menu": [67, "Menu Data, Menu Url"], "is_indexable": [42, "True"], "messaging.www.raq_cards_on_search": [61, "status_quo"], "lsmoney.www.yg_raq_cards_on_search": [152, "status_quo"], "traffic_optimize_loading_js_on_mobile_biz_pages": [16, "no_cohort"], "price_range": [113, 4], "referrer": [64, "none"], "contributions.www.war_compose_signup": [148, "enabled"], "mtb_weekly_growth_1": [85, "search_vs_related_bizs_for_multibiz_raq-search_results_only-www_messaging_online_now-status_quo-not_request_a_quote-not_homeservices-en_US-not_online_now"], "mtb_weekly_growth_2": [101, "ad.request_a_quote_advertisement_desktop-True"], "biz_claimed": [7, "True"], "contributions.www.remove_review_draft_modules": [183, "status_quo"], "yr_diner.www.direct_checkout_for_exact_match": [163, "disabled"], "styleguide_buttons": [13, "status_quo"], "is_biz_user": [129, "False"], "is_paying_business": [117, "False"], "integration": [14, ""], "contributions.www.war_compose_redesign": [6, "enabled"], "maternity_data_baseline": [86, "no_cohort"], "searchux.www.services_serp_card_new_layout_v0_0": [151, "status_quo"], "nav_show_message_count": [59, "show_message_count"], "www_biz_details_raq_sticky": [84, "enabled"], "ad.web_carousel_bottom_of_biz": [78, "status_quo"], "biz_review_count": [10, "100+"], "www_signup_redesign": [92, "status_quo"], "category_paths_to_root": [94, "[['steak', 'restaurants']]"], "decrypted_yuv_id": [3, "DE1CA55AE6AF7DF3"], "biz_closed": [8, "False"], "readerx.web.biz_for_services": [116, "biz_status_quo"], "ytp_eat24_yelp_style_to_iframe": [136, "status_quo"], "distil": [53, null], "internal_ip": [27, "False"], "traffic_in_main_experiment_container": [74, false], "traffic_remove_canonical_links_on_paginated_pages": [130, "no_canonical_2"], "mtb_biz_widget": [169, "False"], "account_level": [1, "anon"], "service": [107, "yelp-main"], "content_country": [15, "US"], "full_url": [34, "/biz/peter-luger-brooklyn-2"], "distil_js_enabled": [138, null], "yelp.www.biz_details.traffic_plah_biz_details_react_migration_experiment": [187, "no_cohort"], "is_business_RAQ_enabled": [69, "disabled"], "has_canonical": [36, "False"], "messaging.www.show_city_in_multibiz": [77, "status_quo"], "readerx.web.popular_dishes": [198, "highlights_popular_dishes"], "platform_pickup_filter": [20, "enabled"], "payment.ux.www": [197, "enabled"], "platform_verticals": [38, "none"], "biz_tt_last_review": [12, "0"], "pagelet_mode_www_biz_details": [24, "allow_async"], "biz_city_state": [192, "Brooklyn, NY"], "www_search_snippets_in_sync_with_ads": [112, "status_quo_8"], "viewport_tracking": [29, null], "traffic_mobile_index_disable_app_pitch_exp": [118, "inactive_not_a_cohort"], "styleguide_typography": [159, "status_quo"], "traffic_bunsen_driven_title_biz_details": [57, "status_quo"], "ytp_order_confirmation_page": [143, "enabled"], "messaging.www.serp_raq_cta_coachmark": [44, "coachmark_3"], "remote_ip": [4, "136.25.5.0"], "second_level_categories": [110, "steak"], "meta_description_case": [79, "review"], "lsm.www.unpakt_cta_change": [108, "None"], "yelp.www.biz_details.traffic_rfn_biz_details_react_migration_experiment": [73, "default_value"], "lsat.www.dropdown_header": [54, "enabled"], "contributions.www.war_compose_recent_photos": [102, "enabled"], "nowait_restaurant.www.no_wait_message": [106, "show_waitlist_instructions"], "review_actions_dropdown": [2, "disabled"], "lower_promoted_delivery_threshold": [58, "reduced_to_fifteen"], "biz_chain": [75, "none"], "readerx.web.biz_details_in_react": [25, "status_quo"], "www_current_location_suggestion": [185, "enabled"], "self_promotion_cohort": [22, "lsg-bunsen"], "ytp_delivery_landing": [125, "platform_pages"], "biz_user_is_owner_of_biz": [128, "False"], "top_level_categories": [11, "restaurants"], "contributions.www.war_attach_photos": [99, "status_quo"], "readerx.web.right_rail_island_refresh": [56, "white_right_rail"], "readerx.www.biz_page_for_restaurants": [62, "status_quo"], "eat24_free_delivery_banner": [45, "disabled"]}, "global": {"rating": [27, 4.0], "biz_closed": [4, "False"], "distil": [12, null], "internal_ip": [18, "False"], "biz_tt_last_review": [8, "0"], "account_level": [1, "anon"], "is_paying_business": [20, "False"], "content_country": [11, "US"], "integration": [17, ""], "biz_num_photos": [5, "100+"], "full_url": [15, "/biz/peter-luger-brooklyn-2"], "distil_js_enabled": [13, null], "meta_description_case": [21, "review"], "has_canonical": [16, "False"], "biz_review_count": [6, "100+"], "is_indexable": [19, "True"], "platform_verticals": [22, "none"], "price_range": [23, 4], "biz_chain": [2, "none"], "referrer": [28, "none"], "top_level_categories": [7, "restaurants"], "category_paths_to_root": [10, "[['steak', 'restaurants']]"], "biz_claimed": [3, "True"]}}, "enable_high_volume_events": false, "trackers": {"www": "UA-30501-24", "global": "UA-30501-1"}, "js_dimensions": {"www": {"platform_order_type": [127, null], "js_vertical_search_type": [39, null]}, "global": {}}, "ga_enabled": true, "clientID": "DE1CA55AE6AF7DF3"},null);
            })();
    </script>

        


            <script>        (function (d, w) {
            if (('ontouchstart' in w) || w.DocumentTouch && d instanceof w.DocumentTouch){
                var html = d.getElementsByTagName("html")[0];
                html.classList.add("touch");
            }
        }(document, window));
</script>


                <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta http-equiv="Content-Language" content="en" />


                    <meta name="description" content="5437 reviews of Peter Luger &#34;It had been about 10 years since the last time we went. Nothing has changed which I love it!! The. Best. Steak. Ever.   My now teens used to come when they were toddlers and they didn&#39;t use to eat, or like steaks much,…"/>


                <meta name="keywords" content="Yelp,recommendation,San Francisco, bay area, local,business,review,friend,restaurant,dentist,doctor,salon,spa,shopping,store,share,community,massage,sushi,pizza,nails,New York,Los Angeles">


                    <meta name="viewport" content="width=1020">

            
    <meta name="application-name" content="Yelp">

    <meta name="msapplication-TileImage" content="https://s3-media3.fl.yelpcdn.com/assets/2/www/img/b8f6d9d556d5/ico/win8-tile.png">
    <meta name="msapplication-TileColor" content="#c41200">

    <meta name="msapplication-starturl" content="https://www.yelp.com/">
    <meta name="msapplication-navbutton-color" content="#c41200">
    <meta name="msapplication-window" content="width=1024;height=768">
    <meta name="msapplication-tooltip" content="Go to Yelp.com">

    <meta name="msapplication-task" content="name=Find Reviews; action-uri=/; icon-uri=https://s3-media1.fl.yelpcdn.com/assets/2/www/img/a6bbc59c7458/ico/favicon-16x16.ico">

    <meta name="msapplication-task" content="name=Hot New Businesses; action-uri=/openings; icon-uri=https://s3-media1.fl.yelpcdn.com/assets/2/www/img/a6bbc59c7458/ico/favicon-16x16.ico">

    <meta name="msapplication-task" content="name=Yelp Deals; action-uri=/search?find_desc=deals; icon-uri=https://s3-media1.fl.yelpcdn.com/assets/2/www/img/a6bbc59c7458/ico/favicon-16x16.ico">


                                <meta name="apple-itunes-app" content="app-id=284910350, app-argument=yelp:///biz/4yPqqJDJOQX69gC66YUDkA?utm_campaign=default&amp;utm_source=www">


                            <link href="android-app://com.yelp.android/yelp-app-indexing/biz/4yPqqJDJOQX69gC66YUDkA?utm_campaign=biz_details&amp;utm_medium=organic&amp;utm_source=google" rel="alternate" />
        <link href="ios-app://284910350/yelp//biz/4yPqqJDJOQX69gC66YUDkA?utm_campaign=biz_details&amp;utm_medium=organic&amp;utm_source=google" rel="alternate" />


                            <meta property="al:ios:app_name" content="Yelp">
        <meta property="al:ios:app_store_id" content="284910350">
        <meta property="al:ios:url" content="https://www.yelp.com/biz/peter-luger-brooklyn-2?utm_campaign=biz_details&amp;utm_medium=organic&amp;utm_source=apple">


                <link rel="mask-icon" sizes="any" href="https://s3-media1.fl.yelpcdn.com/assets/srv0/yelp_styleguide/4374c8fd03d1/assets/img/logos/yelp_burst.svg" content="#c41200">
    <link rel="shortcut icon" href="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/118ff475a341/assets/img/logos/favicon.ico">


                <link rel="search" type="application/opensearchdescription+xml" href="/opensearch" title="Yelp" />



            
        <link rel="canonical" href="https://www.yelp.com/biz/peter-luger-brooklyn-2">
            <link href="https://en.yelp.be/biz/peter-luger-brooklyn-2" hreflang="en-be" rel="alternate">
            <link href="https://www.yelp.fr/biz/peter-luger-brooklyn-2" hreflang="fr-fr" rel="alternate">
            <link href="https://en.yelp.com.ph/biz/peter-luger-brooklyn-2" hreflang="en-ph" rel="alternate">
            <link href="https://fi.yelp.fi/biz/peter-luger-brooklyn-2" hreflang="fi-fi" rel="alternate">
            <link href="https://www.yelp.pt/biz/peter-luger-brooklyn-2" hreflang="pt-pt" rel="alternate">
            <link href="https://en.yelp.my/biz/peter-luger-brooklyn-2" hreflang="en-my" rel="alternate">
            <link href="https://www.yelp.it/biz/peter-luger-brooklyn-2" hreflang="it-it" rel="alternate">
            <link href="https://nl.yelp.be/biz/peter-luger-brooklyn-2" hreflang="nl-be" rel="alternate">
            <link href="https://www.yelp.co.nz/biz/peter-luger-brooklyn-2" hreflang="en-nz" rel="alternate">
            <link href="https://www.yelp.nl/biz/peter-luger-brooklyn-2" hreflang="nl-nl" rel="alternate">
            <link href="https://www.yelp.com/biz/peter-luger-brooklyn-2" hreflang="en-us" rel="alternate">
            <link href="https://www.yelp.com.sg/biz/peter-luger-brooklyn-2" hreflang="en-sg" rel="alternate">
            <link href="https://fr.yelp.be/biz/peter-luger-brooklyn-2" hreflang="fr-be" rel="alternate">
            <link href="https://www.yelp.ie/biz/peter-luger-brooklyn-2" hreflang="en-ie" rel="alternate">
            <link href="https://www.yelp.com.au/biz/peter-luger-brooklyn-2" hreflang="en-au" rel="alternate">
            <link href="https://en.yelp.com.hk/biz/peter-luger-brooklyn-2" hreflang="en-hk" rel="alternate">
            <link href="https://de.yelp.ch/biz/peter-luger-brooklyn-2" hreflang="de-ch" rel="alternate">
            <link href="https://www.yelp.cl/biz/peter-luger-brooklyn-2" hreflang="es-cl" rel="alternate">
            <link href="https://www.yelp.co.jp/biz/peter-luger-brooklyn-2" hreflang="ja-jp" rel="alternate">
            <link href="https://fil.yelp.com.ph/biz/peter-luger-brooklyn-2" hreflang="fil-ph" rel="alternate">
            <link href="https://www.yelp.com.ar/biz/peter-luger-brooklyn-2" hreflang="es-ar" rel="alternate">
            <link href="https://www.yelp.es/biz/peter-luger-brooklyn-2" hreflang="es-es" rel="alternate">
            <link href="https://www.yelp.co.uk/biz/peter-luger-brooklyn-2" hreflang="en-gb" rel="alternate">
            <link href="https://www.yelp.ca/biz/peter-luger-brooklyn-2" hreflang="en-ca" rel="alternate">
            <link href="https://zh.yelp.com.hk/biz/peter-luger-brooklyn-2" hreflang="zh-hk" rel="alternate">
            <link href="https://www.yelp.at/biz/peter-luger-brooklyn-2" hreflang="de-at" rel="alternate">
            <link href="https://www.yelp.com.br/biz/peter-luger-brooklyn-2" hreflang="pt-br" rel="alternate">
            <link href="https://en.yelp.ch/biz/peter-luger-brooklyn-2" hreflang="en-ch" rel="alternate">
            <link href="https://www.yelp.dk/biz/peter-luger-brooklyn-2" hreflang="da-dk" rel="alternate">
            <link href="https://www.yelp.com.tw/biz/peter-luger-brooklyn-2" hreflang="zh-tw" rel="alternate">
            <link href="https://www.yelp.cz/biz/peter-luger-brooklyn-2" hreflang="cs-cz" rel="alternate">
            <link href="https://ms.yelp.my/biz/peter-luger-brooklyn-2" hreflang="ms-my" rel="alternate">
            <link href="https://www.yelp.com.tr/biz/peter-luger-brooklyn-2" hreflang="tr-tr" rel="alternate">
            <link href="https://fr.yelp.ch/biz/peter-luger-brooklyn-2" hreflang="fr-ch" rel="alternate">
            <link href="https://www.yelp.de/biz/peter-luger-brooklyn-2" hreflang="de-de" rel="alternate">
            <link href="https://www.yelp.com.mx/biz/peter-luger-brooklyn-2" hreflang="es-mx" rel="alternate">
            <link href="https://sv.yelp.fi/biz/peter-luger-brooklyn-2" hreflang="sv-fi" rel="alternate">
            <link href="https://fr.yelp.ca/biz/peter-luger-brooklyn-2" hreflang="fr-ca" rel="alternate">
            <link href="https://www.yelp.se/biz/peter-luger-brooklyn-2" hreflang="sv-se" rel="alternate">
            <link href="https://it.yelp.ch/biz/peter-luger-brooklyn-2" hreflang="it-ch" rel="alternate">
            <link href="https://www.yelp.pl/biz/peter-luger-brooklyn-2" hreflang="pl-pl" rel="alternate">
            <link href="https://www.yelp.no/biz/peter-luger-brooklyn-2" hreflang="nb-no" rel="alternate">


                    <link rel="alternate" media="only screen and (max-width: 640px)" href="https://m.yelp.com/biz/peter-luger-brooklyn-2" />



    
            <link rel="next" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=20" />

    
    <script type="application/ld+json">            {"@context": "http://schema.org", "itemListElement": [{"position": 1, "@type": "ListItem", "item": {"@id": "/c/brooklyn/restaurants", "name": "Restaurants"}}, {"position": 2, "@type": "ListItem", "item": {"@id": "/c/brooklyn/steak", "name": "Steakhouses"}}], "@type": "BreadcrumbList"}
</script>


    <script type="application/ld+json">        {"aggregateRating": {"reviewCount": 5437, "@type": "AggregateRating", "ratingValue": 4.0}, "review": [{"reviewRating": {"ratingValue": 5}, "datePublished": "2019-07-08", "description": "It had been about 10 years since the last time we went. Nothing has changed which I love it!! The. Best. Steak. Ever. \n\nMy now teens used to come when they were toddlers and they didn't use to eat, or like steaks much, but now... They killed it. \n\nI've been on diet and don't eat dinner much for the past 2 years, but I did it. I ate some steaks and even dessert, a cheesecake with a whole a lot of cream on it. I shared it with my younger daughter, but we are it all!\n\nLove the atmosphere, food and service. This is a MUST place when we visit NYC.", "author": "Jordyn S."}, {"reviewRating": {"ratingValue": 3}, "datePublished": "2019-07-05", "description": "The cuts of meat are delicious, the service is good, the prices are above average, its location is somewhat removed from Manhattan, I suggest a remodeling to give it a more fresh and innovative look.\n\nLos cortes de carne son  deliciosos, el servicio es bueno, los precios por arriba del promedio, su ubicacion es algo retirada de Manhattan, suguiero una remodelacion para darle un look mas fresco e inovativo", "author": "Guadalupe L."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-07-04", "description": "This is the place where beef becomes steak and steak becomes manna. After all the hoopla surrounding Peter Luger, it was without question that I had to pay a visit and digest some of their world-renowned steaks. I had made a reservation through their website a month in advance for dinner on a Thursday at 7:45 p.m. to celebrate my beloved sister's birthday. The online reservation system was very user-friendly and convenient. \n\nWe arrived 15 minutes early and were told to wait a few minutes for our table. We didn't have to wait long and were seated next to one of the large windows overlooking Broadway Street. Our waiter was jovial, knowledgeable and efficient. He had a dry sense of humor which we found amusing. \n\nWe started off with the Caesar salad which managed to be amazing with crunchy, fresh lettuce and grated Pecorino Romano. We ordered the Steak for Two (which is the Porterhouse steak) and the Rib Steak which has more marbling and fat. We also chose two sides, the creamed spinach and the Luger's Special German fried potatoes. Both sides are good for 2 people. \n\nWhen the food arrived, we took one look at our overladen table and predicted that leftovers would definitely result due to our ambitious ordering strategy. But against all odds and to everyone's amazement, no steak, potato or spinach was left standing. We looked with pride and a sense of accomplishment at the bare bones left on our otherwise empty plates. My sister and I both concluded that we had just experienced meat nirvana. My sister preferred the Porterhouse and I liked the Rib-eye more. Our waiter also recommended eating the steaks sans any gravy or sauce in order to fully enjoy the seasoning and core flavors of the meats.\n\nPayment choices include cash, personal check with ID and debit card. We also got a pile of milk chocolate Peter Luger coins at the end of our incredible meal. We were too stuffed to order any desserts.\n\nThey have a tiny nook of a gift shop in the restaurant which sells Peter Luger merchandise including the Luger sauce, but one can only buy meats online. They ship all over the United States and they accept credit cards for online purchases.\n\nI encourage people to enter the world of Peter Luger where great food and German beer-hallesque ambience abound.", "author": "Wendy S."}, {"reviewRating": {"ratingValue": 4}, "datePublished": "2019-06-28", "description": "Visiting from Toronto, knew we had to stop by here to try out the famous Peter Luger. Came as a party of 6, with a reservation. Even with the reservation expect to chill a bit at the front/bar while they wait and prep your table.\n\nBetween the 6 of us, we ordered the thick-cut bacon, Caesar Salad for starters, Steak for Four, Rib Steak for mains, and the potatoes and creamed Spinach for sides. Also on the table is a gravy boat of their house sauce, which to me tasted similar to a cocktail sauce, definitely has some horseradish.\n\nGoing straight to the steaks, ordering-wise there really wasn't much options so don't worry about not being able to decide what to get. The huge steaks come out on a hot plate, cut up already. The waiter will serve you a couple slices, and drizzle over some oil/butter from the plate. Good crust, strong beef taste. The tenderloin part of the porterhouse is very tender. A good amount of fat, and just a hint of aged flavour.\n\nService is not necessarily the best, but the waiters can be chatty if you engage with them. Price was about ~$70-80 per person with a drink.", "author": "Nathan A."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-06-24", "description": "Im still having trouble finding the words to express how happy and how glad I am after my second Peter Luger experience of the year.\n\nI get reservation for 4 to celebrate my fathers birthday, and enjoying our meeting after 2 years. I asked if we could have a nice table by the window in one of the salons downstairs giving the reason why we would have that special meal, and the nailed it. \n\nWe enjoyed an amazing meal with the best service. We started with some of their delicious breads and butter (it's hard to stop eating! They are completely addictive!) with the tender and juicy bacon and their signature sauce. And following that such an amazing bite, we had the steak for 3 and a green salad.\n\nWhat can I say besides the meat was absorbed perfect, I enjoyed so much better the second time. It's so delicious that even if you're not a meat fan you will enjoy it. You can taste the level of quality in the meat, that comes without any salt. You can add it or having it with their sauce, I enjoyed it just with a tiny bit of salt to embrace the flavors.\n\nWe didn't finish hungry at all, but we needed to have the cheesecake at least, and I ordered also the pecan pie a la mode because either my parents or my sister ever had one before. \n\nSo the service arrived, singing louder happy birthday, bringing the pie with the candle, and giving him a huge surprise (because we are not used to do the same in Spain).\n\nWe had a memorable meal, not just because of the delicious food, the place is just magical.", "author": "Celia N."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-06-22", "description": "Came here with my family and I must say this restaurant did meet my expectations. With all the hype around their steak, I've been wanting to try Peter Luger and I was pleased with the food, service, and price. \n\nMy parents and I got their iceberg Wedge salad, Steak for two, Creamed spinach, and for dessert, we got their Key lime pie with their homemade whipped cream. \n\nThe salad was decent and I especially like the bacon bits. I know they are also known for their bacon and we were going to get that as an appetizer, but I'm glad I got to try some of it from the salad. \n\nThe steak was awesome -- the meat was so juicy, extremely tender, and it literally melted in my mouth. They gave us their special steak sauce, but I mostly just ate the steak by itself with some salt because the meat itself was delicious. It went well with the creamed spinach and I loved it, but my parents thought it was a bit bland. \n\nI'm also glad I got their key lime pie for dessert even though I was so full. The pie was refreshing and tasty, and the whipped cream went well with it even though I usually don't like whipped cream. \n\nOverall, it was definitely one of the best steaks I've ever had and I think it does live up to its reputation.", "author": "Sumin C."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-07-14", "description": "48 hours later after eating at Peter Luger's I finally woke up from the food coma that ensued after feasting. I was warned that the food is so good it'll knock you out and that was true! I made reservations a month in advance and secured a time of 4:45pm on a Saturday. It wasn't busy and we were able to be seated when we arrived at 4:15pm. Our server was Danny and he was a very attentive server. We ordered a pellegrino (sparkling water helps digestion), bacon (2), porterhouse for 2, spinach and the potatoes. I have to say, the bacon was the best I've ever had. I've had bacon from all parts of the spectrum but this bacon brought me to heaven (and clogged my arteries). When the steak came, it was still sizzling in a pool of buttery goodness. My server basted the steak and then served it to me which was definitely a nice touch. The appetizers, spinach and potatoes were definitely up there as well. I've never had creamed spinach like the one at Peter Luger's, it definitely pairs well with the steak. For the steak itself... I suggest rare because it is still cooking as it's brought to the table. I got the medium rare which was still tasty. Don't forget to drench the signature steak sauce onto it! Our bill ended up to be -$172 before tip. I'll definitely return after I take a few months to unclog my arteries. I can't wait to do it all over again.", "author": "Michelle S."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-06-04", "description": "When you visit NYC, you have to check out their most famous steakhouse. We made reservation 1 month ahead and I was super excited to try it out.\n\nDinner was at 7:45pm. When we arrived, the place was packed. After getting my name called, we get seated and a very nice gentleman gave us the menu. We knew exactly what we wanted: Porterhouse for 3 with creamed spinach and fried potatoes. \n\nThe bread came first but pro tip: take the bread home because you're here for the steak. The steak was absolutely mind blowing. We ordered it med rare and it was cooked perfectly. It was very juicy and full of flavor. The creamed spinach was also very tasty.\n\nThe service was absolutely top notch. I will recommend this place to anyone looking for a great steak house experience.", "author": "Mark S."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-05-17", "description": "They just signed on with Resy which allows them to ignore the 4,000 calls per day they get (which they are highly skilled at). After wrestling with Resy and losing the match I decided Resy is only good for raising your blood pressure. \n\nI started calling direct...for months. Yes, I got a ring but couldn't get in the ring. \n\nTheir system is truly democratic. They can/will abuse everyone. They don't care how many Twitter followers you have or who you influence, or that you have staff to handle nasty things like this. It's the same steep climb to the summit. A table.\n  \nClose to yelling P.L. F. U. and throwing the phone in the pool in frustration I tried one last time and got Michael the Manager. What a Prince. \"You're all set\", he assured me after a pleasant chat. (By the time I got there months later they had me down for a table twice the same night and also for three consecutive nights). That was a small glitch and no matter, I was in the system, in the club, and finally...in the door. \n\nI walked in after a 3 mile, one hour Uber from Manhattan and it was bedlam. 300 people were packed 4 deep at an old fashioned bar. (Probably the same one they had when they opened in 1837). Even though I had a reservation that just got me on the list of waiters waiting to be waited on. \n\nAn old fashioned bar tender made me an old fashioned Dirty Martini (stirred, not shaken. Take that, Mr. Bond) and it was easily the finest Martini I've ever enjoyed. The serious pour had me staring at the floor...from close up. \n\nP.L. is not pretentious. Don't bother draggin' your Judith Lieber bag, or strappin' on the Patek Phillipe. No one there cares. Just bring a fat bank roll because they only accept cash and that they do care about. AMEX Reward points sacrificed, it's still worth it for the Bovine bounty you're about to receive.\n\nBe advised there are rules here, and they are strictly enforced. (A) You will suffer to get in. No walk-in's allowed. (B) You will pay cash (and tip generously). (C) You will have to wait even if you're early for your reservation. (D) You will be grilled (like a steak) to make sure your entire party is present before you're seated, (even if you're dining alone). \n\nSo, why, you ask, am I here? Because, (A) It's Peter Luger. (B) It will exceed all the hype. (C) You're gonna have the best hunk-a-meat you've ever had or will have. (Tinder notwithstanding.) \n\nThe Bar: Basics. No chemistry set, periodic table concoctions. Whatever you're handed will have no peer.\n\nThe Vibe: Bucket-listers, wannabe A-listers, real A-listers, Wise Guys, and those who have chosen wisely.  \n\nThe Staff: Jerek was a Big Brother, a Dutch Uncle, Father Confessor, BFF, Master of Meat...a great guy and a top tier true pro. \n\nThe Food: You can memorize the menu in one glance. \n\nHave the thick slice of house cured Bacon even if you're a devout Orthodox Kosher, Halal, Vegan. Once you eat it there can be no turning back. Close your eyes and jump, I say. \n\nThe Caesar Salad (despite the boxed croutons, for which I forgive you, P.L.) was exceptional. Creamy and Garlic-y it was so good, in fact, the ghost of Caesar Cardini shows up periodically to get his fix.\n\nThe Porterhouse for 1, 2, 3, 4, or the whole 182nd Airborne is why you came and it will be dead perfect. Period. \n\nThe steak arrives pre-cut on a platter with a splatter of Jus as large as an Olympic pool. (No, the steak was not swimming in it, it was just luxuriating in it.)\n\nThe meat comes direct from a 1,200 degree grill/oven and is so hot, your face will tan. Remember you're enrolled in steak Master Class. This is the way a Master serves it. \n\nYour server will double spoon the first few slices onto your plate, and  anoint it with Jus, (so good you'll be tempted to ask for a Go Cup of it.) You have just been served food made by God's own Chef.  \n\nAdd to that the necessary/required Creamed Spinach (deliciously more Spinach than Cream...Lawry's eat your heart out). Luger's version of this is divine. \n\nThe Onion Rings got a flour dredge and then a flash fry. Oh my, were they fabulous? Yes, they were.  \n\nNo room for dessert? No normal human would have available space for it. And, yes you ate the entire Porterhouse, and then gnawed on the bone like a Hyenna in public. \n\nNo matter that you can't manage dessert under any circumstances, you're gonna have the Apple Strudel under an Everest of Schlag anyway.You're gonna eat it, and you're gonna like it. I loved it. \n\nWalking out late required wading through 300 people 4 deep at the bar anxiously waiting for their name to be called and there's a reason for that.  \n\nYes, there's a reason this place is off the chart and unique. But, you don't need a reason to go. Just go. Go now!", "author": "Ron W."}, {"reviewRating": {"ratingValue": 4}, "datePublished": "2019-05-30", "description": "Burger review: very good, tasty, patty preparation is biased towards rare (we asked for medium rare) which I personally don't mind, but figured this is good to know. You have the option to add cheese to the burger- totally up to you. I added cheese and enjoyed it but will definitely have the burger without cheese on the next run so I can better appreciate the beef. Quality bun, not too soft and has form. Slightly on the firmer side. \n\nTried an order of the thick cut bacon- nice fat/meat ratio. Really melts in your mouth, but flavor is slightly lacking. \n\nCame on a Thursday during lunch with a friend, we had to wait about 20 minutes. Coming on a Friday or weekend expect a longer wait. Cash only.\n\nOverall a solid burger and will return for another.", "author": "Rashad A."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-05-25", "description": "I came to NYC for this. It's been a fee years since we've last had it and in my mind it's the best steak ever. \n\nThis is definitely at the top of my list for best Steakhouses ever. The perfect cut, aged the right time and the finishing touch. Its full of flavour and juices. Pair it with some houSr wine and you've got the combo ready. And dont forget to finish the bone, don't be shy. Wish I could have this every month. And dont forget to pair it with some creamed spinach and a milkshake if you still have room. \n\nService is pretty good in general. Everytime we've been the server's have been quite the jokers. This place is cash or debit only so keep that in mind. And make sure to have a reservation no matter the time . Otherwise you would be waiting in line with the tourists for a good 2 hours. It'll be a painful wait", "author": "Vince L."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-05-16", "description": "We ventured one hour to get here. It's a classic must-go steakhouse of Brooklyn / New York. When we got there, there was still a wait despite reservations. \n\nMake sure you get a reservation! \n\nWe were then seated and got their prime beef steak for two. Then also got some creamed spinach.\n\nAt this point of the trip we had already eaten a lot, so the fact that I could still taste how delicious the steak is (juicy, tender, flavourful) shows how great it is.\n\nBill came out to be $120.", "author": "Lesley Y."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-05-08", "description": "I was last here 20 years ago but I'm worth it so it was time to revisit!!\nNo more sawdust. \nNeed reservations weeks in advance \nAmazing servers who have been there for ages and know their meat!!\nJust awesome steak cooked exactly how you ask. \nSide of creamer spinach - delicious \nGerman potatoes - cooked with butter then baked to perfect crunch. \nApple strudel was good but the home made Schlag is even better. I can eat that plain", "author": "Dona F."}, {"reviewRating": {"ratingValue": 4}, "datePublished": "2019-05-12", "description": "Dinner at Peter Luger exceeded my expectations. It's quite casual for a place with a Michelin star; dressing up is unnecessary, yet you won't be out of place if you want to make an occasion of it. I had heard rumors of cold service but our experience was quite the opposite; our waiter was patient and helpful -- he suggested pairing Malbec wine with our steak and it was a wonderful choice.\n\nAlthough I usually prefer warm bread, the rolls that they brought to our table were delicious. I especially liked the roll filled with onion, and found it hard to pace myself knowing all the food to come! We started off with the thick-cut bacon and a caesar salad. I'm not even a huge bacon person but I will say that it's something you should try here. Tastes great with the steak sauce! The caesar salad was fantastic and dressed perfectly -- probably one of my favorite parts of the meal actually. The steak though was truly amazing. So incredibly tender, definitely among the best steaks I've had. Portion size was quite impressive! We also got the baked potato and onion rings as sides, both of which I would pass on next time. The potato was pretty standard (nothing you couldn't do at home), and the onion rings were thin and stringy. \n\nFor dessert, we got the pecan pie with the homemade schlag of course! The pie was good although not super memorable, but we did enjoy the mound of whipped cream! To make you feel a little better about the sum you spend on this dinner, they send you away with some chocolate coins -- a nice touch. An experience very worth having.", "author": "Melissa G."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-04-29", "description": "Cash only\n\nHave you ever eaten a piece of steak and thought to yourself, \"this is the best steak I've ever had in my life\"? Well that's what happened to me at Peter Luger. This place is generally seen as one of the best steak houses in New York and I am glad I was about to treat myself here for my birthday.\n\nI'd highly recommend making reservations. You can make them online or over the phone. I actually arrived half an hour before my reserved time and got seated within a couple minutes.\n\nI ordered the sizzling bacon ($6.95) which was smoky and fatty and delicious. It was so easy to cut through. I also got the single steak ($54.95) which was heavenly! Seriously, don't get me started on the steak... it was amazing and cooked perfectly medium rare. The crispiness and the slight char on the outside and then the juiciness and tenderness on the inside. I'm salivating just thinking about it. It comes with dipping sauce which is sweet and tastes like it's ketchup based. Both of these were just melt in your mouth and I recommend getting these. I've also heard great things about the burger but sadly they only have it for lunch so I wasn't able to try it.\n\nNote: I've heard of the troubles of making a reservation here. You can actually make reservations on their website now! I still recommend you make one in advance; I made mine three weeks in advanced or so and there were no problems. If you'd prefer to speak to someone you can call them. If you call, I recommend calling them right when they open as I only needed to wait a couple minutes until someone answered.\n\nIs this the best steak I've ever had? Yup.", "author": "Brian F."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-05-29", "description": "Came here with my family for a post-graduation meal! It was 100% worth the trek to Brooklyn and the 8:45 PM reservation time (make your reservations well in advance, because this was the earliest I could get 2 months in advance for a Wednesday evening!).\n\nThe four of us ordered the steak for two, the creamed spinach, the bacon, and the onion rings. I can't even describe just how good everything was -- that bacon was unlike anything else I'd ever put into my mouth. The steak too -- nothing is as good as the first bite of steak, and paired with the Peter Luger's sauce, was a dish I will never forget. Creamed spinach -- I thought it was good, my family thought it was amazing. Even my sister and dad who hate spinach ate it by the spoonful. The onion rings were fresh and hot and tasty as well, a little bit on the greasy side but that's what makes them taste so good the next day.\n\nOur server was so funny and attentive, he made us feel welcomed from the start. My dad left his backpack at the restaurant on accident and they went above and beyond in helping him get it back the following day. Definitely a great splurge-worthy meal and experience!", "author": "Phoebe W."}, {"reviewRating": {"ratingValue": 4}, "datePublished": "2019-05-09", "description": "I made reservation for my friends whom visited in NY from LA. One of my friend birthday coming up and I decided to took my friends at this place. I've been tried this place but I could tell you that their prime steak was the best I ever had in this place before. To honestly I like this place but My experience wasn't good in here. \n\nStopped by for dinner with friends. We ordered 3 of prime steak and side of creamy spinach, mixed green salad. 3 of cosmopolitan cocktails, one of bartender special cocktail and key lime pie with coffee, vanilla ice cream for dessert.\n\nPrime steak(4.5/5) was comes out hot also tasted great than I expected. Very juicy and tender, soft. Peter Luger steak sauce is famous. I loved their sauce and my all of friends were looks like they love steak with sauce. \n\nCreamy spinach(2/5)wasn't creamy than I expected. I've had eat better elsewhere. Not creamy much but savory. If you like less creamy style, you would like it.\n\nMixed green salad(3/5) was ok but very fresh vegetables. Nothing special, you can eat this salad at your home too. \n\nKey lime pie(3/5)was good but I've eat better elsewhere. I am a huge fan of dessert also I love key lime pie. I eaten that in my whole life but this place key lime pie wasn't the best. Home made whipped cream was the best. Not too sweet but kinda of savory with very soft. Even if you don't like sweets, you can try key lime pie and their home made whipped cream. But I would recommend you that if you want to eat key lime pie, drink coffee. \n\nVanilla ice cream (2/5) just ok. \n\nCosmopolitan(2/5) were really strong cocktails. Actually it isn't strong cocktails but the bartender made our cocktails very strong. I had to request them to extra cranberry juice. \n\nService(3/5)-our waiter was accommodating and very friendly. They load up our plate with steak and side. But there's busy and it was difficult to get for what we need. Very little of space between table to table. One waiter was keeping touch my chair or moved my chair by accidentally but nothing apologized. I understand how they're busy and not enough space for served but i personally thought he could say at least sorry or excuse me. Am I too picky? \n\nAmbience(2.5/5)- old restaurant and their interior is wood style. It's simple interior but I couldn't see any of special point of interior. \n\nThis place is cash only. If you have plan for go this place then have some cash. No dress cord here.", "author": "Jae H."}, {"reviewRating": {"ratingValue": 5}, "datePublished": "2019-05-14", "description": "Where to start....I have been live in nyc for almost a decade and this is the FIRST TIME I came to Peter Luger, I have to say I absolutely love it !!! I know this steakhouse is very famous, but I didn't expect the fact that I have to make a reservation one week ahead. However it all worth it at the end! \n\nIt was PACKED when we get here and people already waiting in the line. The interior design is old fashioned, classic steakhouse style which I love. \n\nWe ordered their famous bacon strips, porterhouse steak for two, wedge salad, cream Spanish , French fries, and pecan pie with their homemade whip cream. Everything tasted sooo delicious. Especially the thick cut bacon strip with their house steak Sauce.. it was like in heaven... hahaha.  The steak was cooked beautifully , tender and juicy, the sound of the sizzling pan make it extra mouth watering... My friend is not a dessert lover, but she was to able to finish the whole pecan pie herself...  \n\nFood price is very reasonable too, we (two people)spent around $300 + tip. Which is very normal for a good steak house in nyc.\n\nHighly recommended !!!!", "author": "Danny L."}, {"reviewRating": {"ratingValue": 3}, "datePublished": "2019-05-12", "description": "Made a reservation over the phone for my boyfriend's birthday about a month in advance. Service was great, but unfortunately we found the food to be hit or miss. Rib eye, Caesar salad, creamed spinach and apple strudel were all great, but the German potatoes were dry and the key lime pie had a weird chemical taste to it. On top of that, the dishes we enjoyed were just that - enjoyable. At a place like this, I want to be blown away when I take a bite. That wasn't really the case here, and so I think I am better off returning to places like 4 Charles and Del Frisco's to get my meat and potatoes fix. This is a classic place that I'm glad I tried once, but I'm not sure I'd recommend it if you're looking for a singular NYC steak experience. \n\n3.5 stars.", "author": "Gab G."}, {"reviewRating": {"ratingValue": 4}, "datePublished": "2019-05-08", "description": "Ordered Steak for 2. Asked for medium but got medium rare. Because the plate was so hot, we waited for the steak on the plate to baste itself more before consuming. The porterhouse steak is on the leaner side, not as marbled as we would like, but is still good. \n\nRecommend to make a reservation, lines for walk in gets long. No credit cards, cash or debit.\n\n(This review written by husband)", "author": "Yonnie C."}], "servesCuisine": "Steakhouses", "priceRange": "Above $61", "name": "Peter Luger", "address": {"addressLocality": "Brooklyn", "addressRegion": "NY", "streetAddress": "178 Broadway", "postalCode": "11211", "addressCountry": "US"}, "@context": "http://schema.org/", "image": "https://s3-media3.fl.yelpcdn.com/bphoto/jOoRBpCU9_2iS8z306CGOQ/ls.jpg", "@type": "Restaurant", "telephone": "+17183877400"}
</script>

                <meta name="critical_css_middleware">

            <meta property="fb:app_id" content="97534753161">
        <meta property="og:description" content="Steakhouses in Brooklyn, NY">
        <meta property="og:image" content="https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/o.jpg">
        <meta property="og:image:height" content="1088">
        <meta property="og:image:width" content="1600">
        <meta property="og:site_name" content="Yelp">
        <meta property="og:title" content="Peter Luger - Williamsburg - South Side - Brooklyn, NY">
        <meta property="og:type" content="yelpyelp:business">
        <meta property="og:url" content="https://www.yelp.com/biz/peter-luger-brooklyn-2">



            <meta property="twitter:card" content="summary">
        <meta property="twitter:site" content="@yelp">
        <meta property="twitter:domain" content="yelp.com">
        <meta property="twitter:app:name:iphone" content="Yelp">
        <meta property="twitter:app:name:ipad" content="Yelp">
        <meta property="twitter:app:name:googleplay" content="Yelp">
        <meta property="twitter:app:id:iphone" content="id284910350">
        <meta property="twitter:app:id:ipad" content="id284910350">
        <meta property="twitter:app:id:googleplay" content="com.yelp.android">
        <meta property="twitter:image" content="https://s3-media1.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/258s.jpg">
        <meta property="twitter:app:url:iphone" content="yelp:///biz/4yPqqJDJOQX69gC66YUDkA?utm_campaign=default&amp;utm_source=twitter-card">
        <meta property="twitter:app:url:ipad" content="yelp:///biz/4yPqqJDJOQX69gC66YUDkA?utm_campaign=default&amp;utm_source=twitter-card">
        <meta property="twitter:app:url:googleplay" content="intent://yelp.com/biz/4yPqqJDJOQX69gC66YUDkA?utm_source=twitter-card#Intent;scheme=http;package=com.yelp.android;end;">


    <meta name="yelp-biz-id" content="4yPqqJDJOQX69gC66YUDkA">

        
            <title>Peter Luger - 7586 Photos & 5437 Reviews - Steakhouses - 178 Broadway, Williamsburg - South Side, Brooklyn, NY - Restaurant Reviews - Phone Number - Menu - Yelp</title>


            <link rel="stylesheet" type="text/css" media="all" href="https://s3-media0.fl.yelpcdn.com/assets/public/yelp_main.yji-69d27c15b058ee43fee3.chunk.css"><link rel="stylesheet" type="text/css" media="all" href="https://s3-media0.fl.yelpcdn.com/assets/public/yelp_main.yji-2203196b83b9f66f79fa.chunk.css"><link rel="stylesheet" type="text/css" media="all" href="https://s3-media0.fl.yelpcdn.com/assets/public/yelp_main.yji-bacaf7492a8bf6fa0999.chunk.css">

            <!--[if (IE 9)|(lt IE 9) ]> <link rel="stylesheet" type="text/css" media="all" href="https://s3-media2.fl.yelpcdn.com/assets/2/www/css/8ff4a59129bf/www-pkg-main.css">
            <link rel="stylesheet" type="text/css" media="all" href="https://s3-media4.fl.yelpcdn.com/assets/2/www/css/e5e4fe9ba9d4/www-pkg-more.css"> <![endif]-->
            <!--[if (gt IE 9)|!(IE)]><!--> <link rel="stylesheet" type="text/css" media="all" href="https://s3-media4.fl.yelpcdn.com/assets/2/www/css/a0c2c519995c/www-pkg.css"> <!--<![endif]-->
    <link rel="stylesheet" type="text/css" media="all" href="https://s3-media1.fl.yelpcdn.com/assets/2/www/css/87b84eea08d9/biz_details-pkg.css">


        

    <script></script>


                    <link href="https://www.facebook.com/yelp" itemprop="sameAs" id="sitelink-profile-facebook">
        <link href="https://twitter.com/yelp" itemprop="sameAs" id="sitelink-profile-twitter">
        <link href="https://www.linkedin.com/company/31517" itemprop="sameAs" id="sitelink-profile-linkedin">
        <link href="https://instagram.com/yelp" itemprop="sameAs" id="sitelink-profile-instagram">

    <meta itemscope itemtype="http://schema.org/WebSite" itemref="sitename sitelink  sitelink-profile-facebook sitelink-profile-twitter sitelink-profile-linkedin sitelink-profile-instagram">
    <meta itemprop="name" content="Yelp" id="sitename">
    <link href="https://www.yelp.com/" itemprop="url" id="sitelink">


            <!-- css-middleware: insert inline css -->
            <script>            document.cookie = "xcj=1|ZC9X0zvT3oyyh9GF_A5KiD3-DA2-Wnlkjy3Dgxu-nPk; Domain=.yelp.com; expires=Sun, 13 Oct 2019 03:44:36 GMT; Max-Age=7776000; Path=/";
</script>

    </head>














    <body id="yelp_main_body" class="jquery country-us logged-out biz-details react-modal-body">

        <script>(function (d, w) {
    var supportsSVG = (
        !!d.createElementNS &&
        !!d.createElementNS(
            'http://www.w3.org/2000/svg',
            'svg'
        ).createSVGRect
    );
    var cdnPath = 'https://s3-media3.fl.yelpcdn.com/assets/srv0/svg_icons/be89abbbfc6f/assets/svg_sprite.js';
    var head = d.getElementsByTagName('head')[0];
    function fallback() {
        var link = d.createElement('link');
        link.rel = 'stylesheet';
        link.href = 'https://s3-media3.fl.yelpcdn.com/assets/srv0/svg_icons/fcf6338181b8/assets/sprite.css';
        head.appendChild(link);
    }
    if (!supportsSVG) {
        fallback();
        return;
    }
    if (!w.yelp) {
        w.yelp = {};
    }
    w.yelp.__injectSvgSpritesheet = function (svg) {
        document.body.insertAdjacentHTML('afterbegin', svg);
        delete window.yelp.__injectSvgSpritesheet;
    }
    function onError() {
        d.documentElement.className = [
            d.documentElement.className,
            'icon-svg-unavailable'
        ].join(' ');
        fallback();
        return true;
    }
    var script = d.createElement('script');
    script.async = true;
    script.onerror = onError;
    script.src = cdnPath + '?callback=window.yelp.__injectSvgSpritesheet';
    head.appendChild(script);
}(document, window));</script>
<noscript>
    <link rel="stylesheet" href="https://s3-media3.fl.yelpcdn.com/assets/srv0/svg_icons/fcf6338181b8/assets/sprite.css">
</noscript>

            <script id="yelp-js-error-reporting-init-error-reporting" type="application/json">{"enabled": true, "apiKey": "ac0c27867d3912ce5714ff131041af4f", "config": {"release": "e8b0b8210c131c668c7eb764d3e7fe7d7a4650b8", "environment": "prod", "sampleRate": 1.0, "blacklist": {"message": ["Blocked a frame with origin"], "type": [], "location": ["yl[a-z]-[a-z]*\\.js"]}, "metaData": {"projectName": "yelp_main"}}}</script>


    <script>            window.pageSpeedCustomTimestamps['body:start'] = (new Date()).getTime();
</script>

            <div id="fb-root"></div>



        
        <div id="wrap" class="lang-en">
                <div class="page-header">


                        

                    



                                        

                            <ul class="offscreen">
                <li>
                    <a href="#header_find_form" rel="nofollow">
                        Skip to Search Form
                    </a>
                </li>
                <li>
                    <a href="#header-nav" rel="nofollow">
                        Skip to Navigation
                    </a>
                </li>
                <li>
                    <a href="#page-content" rel="nofollow">
                        Skip to Page Content
                    </a>
                </li>
        </ul>







        <div class="main-header main-content-wrap js-main-header webview-hidden main-header--lsat">
            <div class="main-header_wrapper">
                <div class="content-container">

                    <div class="arrange arrange--18 arrange--middle main-header_arrange">

                            <div class="arrange_unit main-header--full_arrange_unit">
                                
    <div class="main-header_logo js-analytics-click" id="logo" data-analytics-label="logo">
        <a href="/">Yelp</a>
    </div>

                            </div>



                        <div class="arrange_unit arrange_unit--fill align-middle main-header--full_arrange_unit main-header_search-container">
                    <div class="main-header_search responsive-hidden-medium-only">
                <form method="get" action="/search" id="header_find_form" class="main-search yform u-space-b0" role="search">
        <div class="arrange arrange--stack">
            <div class="arrange_unit arrange_unit--fill">
                <div class="arrange arrange--equal arrange--stack">
                    <div class="arrange_unit main-search_search-field-arrange">
                            <div class="main-search_suggestions-field search-field-container find-decorator">
        <label class="main-search_pseudo-input main-search_pseudo-input--find pseudo-input" for="find_desc">
            <span class="pseudo-input_text">Find</span>
            <span class="pseudo-input_field-holder">
                <input maxlength="64" name="find_desc" id="find_desc" autocomplete="off" value="" placeholder="tacos, cheap dinner, Max’s" class="main-search_field pseudo-input_field" aria-autocomplete="list" tabindex="1">
            </span>
        </label>
        <div class="main-search_suggestions suggestions-list-container search-suggestions-list-container hidden">
            <ul class="suggestions-list" role="listbox" aria-label="Search results"></ul>
        </div>
    </div>

                    </div>
                    <div class="arrange_unit main-search_near-field-arrange  main-search_search-field-arrange arrange_unit--stack-12">
                            <div class="main-search_suggestions-field search-field-container near-decorator">
        <label class="main-search_pseudo-input main-search_pseudo-input--near pseudo-input">
            <span class="pseudo-input_text">Near</span>
            <span class="main-search_field-holder pseudo-input_field-holder">
                <input maxlength="80" name="find_loc" id="dropperText_Mast" autocomplete="off" value="San Francisco, CA" placeholder="address, neighborhood, city, state or zip" data-query="San Francisco, CA" class="main-search_field pseudo-input_field" aria-autocomplete="list" tabindex="2">
                <input type="hidden" name="ns" value="1">
            </span>
        </label>
        <div class="main-search_suggestions suggestions-list-container location-suggestions-list-container hidden">
            <ul class="suggestions-list" role="listbox" aria-label="Search results"></ul>
        </div>
    </div>

                    </div>
                </div>
            </div>
            <div class="arrange_unit main-search_actions arrange_unit--stack-12">
                <div class="arrange arrange--wrap arrange--6">
                    <div class="arrange_unit hidden-non-responsive-inline-block responsive-visible-small-inline-block main-search_action">

    <a class="ybtn ybtn--primary main-header_button js-search-close main-search_close">
                            Cancel
    </a>
                    </div>
                    <div class="arrange_unit main-search_action arrange_unit--stack-12">

    <button class="ybtn ybtn--primary main-search_submit main-header_button" id="header-search-submit" tabindex="3" title="Search" type="submit" value="submit">
                            <span class="main-search_action-icon-wrap js-search-icon-wrap">
                                <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-search icon--size-24 icon--inverse icon--fallback-inverted">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_search" />
    </svg>
</span>
                                <span class="u-offscreen">Search</span>
                            </span>
                                <div class="circle-spinner js-circle-spinner hidden">
        <div class="circle-spinner_segment container1">
            <div class="circle1"></div>
            <div class="circle2"></div>
            <div class="circle3"></div>
            <div class="circle4"></div>
        </div>
        <div class="circle-spinner_segment container2">
            <div class="circle1"></div>
            <div class="circle2"></div>
            <div class="circle3"></div>
            <div class="circle4"></div>
        </div>
        <div class="circle-spinner_segment container3">
            <div class="circle1"></div>
            <div class="circle2"></div>
            <div class="circle3"></div>
            <div class="circle4"></div>
        </div>
    </div>

    </button>
                    </div>
                </div>
            </div>
        </div>
    </form>

        </div>

                        </div>



                            <div class="arrange_unit main-header--full_arrange_unit">
                                <div class="arrange arrange--6">
                                    <div class="arrange_unit u-nowrap">
                                        <div class="main-header_account webview-hidden">
                                                        <ul class="header-nav">
            <li class="header-nav_item responsive-hidden-small js-analytics-click" data-analytics-label="signup">

    <a class="ybtn ybtn--primary main-header_button header-nav_button--sign-up js-sign-up-button" href="/signup" id="header-sign-up">
                    Sign Up
    </a>
            </li>
        <li class="header-nav_item">
                    <a class="header-nav_link header-nav_link--log-in js-log-in-button" href="https://www.yelp.com/login?return_url=%2Fbiz%2Fpeter-luger-brooklyn-2">
            Log In
        </a>

        </li>
    </ul>
        
    <div id="topbar-account-item" class="user-account clearfix drop-menu-origin hidden-non-responsive-block responsive-visible-small-block">

    <a class="ybtn ybtn--primary drop-menu-link user-account_button" href="javascript:;" id="topbar-account-link">
            <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-triangle-down icon--size-14 icon--inverse icon--fallback-inverted u-triangle-direction-down user-account_button-arrow responsive-visible-large-inline-block">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_triangle_down" />
    </svg>
</span>
            <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-hamburger icon--size-24 icon--inverse icon--fallback-inverted drop-menu-link_open">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_hamburger" />
    </svg>
</span>
            <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-close icon--size-24 icon--inverse icon--fallback-inverted drop-menu-link_close">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_close" />
    </svg>
</span>
    </a>
        <div id="topbar-account-wrap" class="drop-menu drop-menu-has-arrow">
            <div class="drop-menu-arrow responsive-hidden-small"></div>

                <div class="arrange arrange--6 arrange--equal drop-menu_auth-buttons">
                    <div class="arrange_unit">

    <a class="ybtn ybtn--secondary ybtn-full" href="https://www.yelp.com/login?return_url=%2Fbiz%2Fpeter-luger-brooklyn-2">
                            Login
    </a>
                    </div>
                    <div class="arrange_unit">

    <a class="ybtn ybtn--primary ybtn-full" href="/signup">
                            Sign Up
    </a>
                    </div>
                </div>

                <ul class="drop-menu-group--nav drop-menu-group">
                            <li class="drop-down-menu-link" role="none">
        <a class="js-analytics-click arrange arrange--middle arrange--6" href="/user_details" data-analytics-label="" role="menuitem" tabindex="0">
            <strong class="arrange_unit">
                    <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-profile icon--size-24 u-space-r1">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_profile" />
    </svg>
</span>About Me
            </strong>
            <span class="arrange_unit arrange_unit--fill u-text-right">
                <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-chevron-right icon--size-24 hidden-non-responsive-inline-block responsive-visible-medium-inline-block">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_chevron_right" />
    </svg>
</span>
            </span>
        </a>
    </li>

                            
                            
                            <li class="drop-down-menu-link hidden-non-responsive-block responsive-visible-medium-block" role="none">
        <a class="js-analytics-click arrange arrange--middle arrange--6" href="/talk" data-analytics-label="" role="menuitem" tabindex="0">
            <strong class="arrange_unit">
                    <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-talk icon--size-24 u-space-r1">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_talk" />
    </svg>
</span>Talk
            </strong>
            <span class="arrange_unit arrange_unit--fill u-text-right">
                <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-chevron-right icon--size-24 hidden-non-responsive-inline-block responsive-visible-medium-inline-block">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_chevron_right" />
    </svg>
</span>
            </span>
        </a>
    </li>

                </ul>

        </div>
    </div>



                                        </div>
                                    </div>
                                </div>
                            </div>

                    </div>

                </div>
            </div>
        </div>


            <div class="js-curloc-error-modal curloc-error-modal hidden">
    <div class="modal_head">
            <h2>Oops, we can't find your location</h2>
    </div>
    <div class="modal_body">
        <div class="curloc-permission-denied hidden">
            <p>
                First, try refreshing the page and clicking Current Location again. Make sure you click <b>Allow</b> or <b>Grant Permissions</b> if your browser asks for your location. If your browser doesn't ask you, try these steps:
            </p>
            <div class="chrome hidden">
                <ol class="modal_section u-bg-color">
                    <li>At the top of your Chrome window, near the web address, click <b>the green lock labeled Secure</b>.</li>
                    <li>In the window that pops up, make sure <b>Location</b> is set to <b>Ask</b> or <b>Allow</b>.</li>
                    <li>You're good to go! Reload this Yelp page and try your search again.</li>
                </ol>
                <p>
                    If you're still having trouble, check out <a href="https://support.google.com/chrome/answer/142065?co=GENIE.Platform%3DDesktop&hl=en&oco=0" target="_blank">Google's support page</a>.
                    You can also search near a city, place, or address instead.
                </p>
            </div>
            <div class="opera hidden">
                <ol class="modal_section u-bg-color">
                    <li>At the top of your Opera window, near the web address, you should see a <b>gray location pin</b>. Click it.</li>
                    <li>In the window that pops up, click <b>Clear This Setting</b></li>
                    <li>You're good to go! Reload this Yelp page and try your search again.</li>
                </ol>
                <p>
                    If you're still having trouble, check out <a target="_blank">Opera's support page</a>.
                    You can also search near a city, place, or address instead.
                </p>
            </div>
            <div class="safari hidden">
                <ol class="modal_section u-bg-color">
                    <li>Click <b>Safari</b> in the Menu Bar at the top of the screen, then <b>Preferences</b>.</li>
                    <li>Click the <b>Privacy</b> tab.</li>
                    <li>Under <b>Website use of location services</b>, click <b>Prompt for each website once each day</b> or <b>Prompt for each website one time only</b>.</li>
                    <li>MacOS may now prompt you to enable Location Services. If it does, follow its instructions to enable Location Services for Safari.</li>
                    <li>Close the Privacy menu and refresh the page. Try using Current Location search again. If it works, great! If not, read on for more instructions.</li>
                    <li>Back in the <b>Privacy</b> dialog, Click <b>Manage Website Data...</b> and type <b>yelp.com</b> into the search bar.</li>
                    <li>Click the <b>yelp.com</b> entry and click Remove.</li>
                    <li>You're good to go! Close the Settings tab, reload this Yelp page, and try your search again.</li>
                </ol>
                <p>
                    If you're still having trouble, check out <a href="https://support.apple.com/en-us/HT204690" target="_blank">Safari's support page</a>.
                    You can also search near a city, place, or address instead.
                </p>
            </div>
            <div class="firefox hidden">
                <ol class="modal_section u-bg-color">
                    <li>At the top of your Firefox window, to the left of the web address, you should see a <b>green lock</b>. Click it.</li>
                    <li>In the window that pops up, you should see <b>Blocked</b> or <b>Blocked Temporarily</b> next to <b>Access Your Location</b>. Click the <b>x</b> next to this line.</li>
                    <li>You're good to go! Refresh this Yelp page and try your search again.</li>
                </ol>
                <p>
                    If you're still having trouble, check out <a href="https://www.mozilla.org/en-US/firefox/geolocation/" target="_blank">Firefox's support page</a>.
                    You can also search near a city, place, or address instead.
                </p>
            </div>
            <div class="ie hidden">
                <ol class="modal_section u-bg-color">
                    <li>Click the <b>gear</b> in the upper-right hand corner of the window, then <b>Internet options</b>.</li>
                    <li>Click the <b>Privacy</b> tab in the new window that just appeared.</li>
                    <li>Uncheck the box labeled <b>Never allow websites to request your physical location</b> if it's already checked.</li>
                    <li>Click the button labeled <b>Clear Sites</b>.</li>
                    <li>You're good to go! Click <b>OK</b>, then refresh this Yelp page and try your search again.</li>
                </ol>
                <p>
                    You can also search near a city, place, or address instead.
                </p>
            </div>
            <div class="edge hidden">
                <ol class="modal_section u-bg-color">
                    <li>At the top-right hand corner of the window, click the <b>button with three dots on it</b>, then <b>Settings</b>.</li>
                    <li>Click <b>Choose what to clear</b> underneath <b>Clear browsing data</b>.</li>
                    <li>Click <b>Show more</b>, then make sure only the box labeled <b>Location permissions</b> is checked.</li>
                    <li>Click <b>Clear</b>.</li>
                    <li>You're good to go! Refresh this Yelp page and try your search again.</li>
                </ol>
                <p>
                    You can also search near a city, place, or address instead.
                </p>
            </div>
            <div class="default hidden">
                <p>
                    Oops! We don't recognize the web browser you're currently using. Try checking the browser's help menu, or searching the Web for instructions to turn on HTML5 Geolocation for your browser.
                    You can also search near a city, place, or address instead.
                </p>
            </div>
        </div>
        <div class="curloc-unavailable hidden">
            Something broke and we're not sure what. Try again later, or search near a city, place, or address instead.
        </div>
        <div class="curloc-timeout hidden">
            We couldn't find you quickly enough! Try again later, or search near a city, place, or address instead.
        </div>
        <div class="curloc-inaccurate hidden">
            We couldn't find an accurate position. If you're using a laptop or tablet, try moving it somewhere else and give it another go. Or, search near a city, place, or address instead.
        </div>
    </div>
</div>



    <div class="main-header_nav-wrapper js-header-nav main-content-wrap responsive-hidden-medium u-hidden">
        <div class="content-container">
            <nav class="main-header_nav u-nowrap">
                <div class="arrange">
                    <div class="arrange_unit arrange_unit--fill">
            <div id="header-nav">

                    <ul class="header-nav" role="menubar">
            <li 
                class="header-nav_item js-analytics-click"
                data-analytics-label="restaurants"
                role="none"
            >
                


    <div class="header-nav_link_container">
            <a class="header-nav_link" href="/search?cflt=restaurants&amp;find_loc=San+Francisco%2C+CA">
                <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-food icon--size-18 icon--white icon--fallback-inverted u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_food" />
    </svg>
</span>

            Restaurants

        </a>
    </div>

            </li>
            <li 
                class="header-nav_item js-analytics-click"
                data-analytics-label="nightlife"
                role="none"
            >
                


    <div class="header-nav_link_container">
            <a class="header-nav_link" href="/search?cflt=nightlife&amp;find_loc=San+Francisco%2C+CA">
                <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-nightlife icon--size-18 icon--white icon--fallback-inverted u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_nightlife" />
    </svg>
</span>

            Nightlife

        </a>
    </div>

            </li>
            <li 
                class="header-nav_item js-analytics-click"
                data-analytics-label="home services"
                role="none"
            >
                


    <div class="header-nav_link_container has-dropdown dropdown--arrow dropdown--fade fade-initial-load">
            <a class="header-nav_link js-dropdown-toggle dropdown_toggle" href="/search?cflt=homeservices&amp;find_loc=San+Francisco%2C+CA">
                <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-home-services icon--size-18 icon--white icon--fallback-inverted u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_home_services" />
    </svg>
</span>

            Home Services

                <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-triangle-down icon--size-14 icon--currentColor u-triangle-direction-down dropdown_arrow">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_triangle_down" />
    </svg>
</span>
        </a>
                <div class="dropdown_menu-container js-dropdown_menu-container">
        <div class="dropdown_menu js-dropdown-menu">
            <div class="dropdown_menu-inner">
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item js-analytics-click" data-analytics-label="air conditioning &amp; heating" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=hvac&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Air Conditioning & Heating
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="contractors" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=contractors&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Contractors
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="electricians" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=electricians&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Electricians
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="home cleaners" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=homecleaning&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Home Cleaners
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="landscapers" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=landscaping&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Landscapers
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="locksmiths" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=locksmiths&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Locksmiths
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="movers" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=movers&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Movers
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="painters" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=painters&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Painters
        </span>
    </a>

                            </li>
                            <li class="dropdown_item js-analytics-click" data-analytics-label="plumbers" role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="/search?cflt=plumbing&amp;find_loc=San+Francisco%2C+CA" role="menuitem">
        <span class="dropdown_label">
                Plumbers
        </span>
    </a>

                            </li>
                    </ul>
            </div>
        </div>
    </div>

    </div>

            </li>
    </ul>


                    <ul class="header-nav" role="menubar">
            <li id="write-review" class="header-nav_item js-analytics-click" data-analytics-label="write-review" role="none">
                    


    <div class="header-nav_link_container">
            <a class="header-nav_link" href="/writeareview">

            Write a Review

        </a>
    </div>

            </li>
            <li id="events" class="header-nav_item js-analytics-click" data-analytics-label="events" role="none">
                    


    <div class="header-nav_link_container">
            <a class="header-nav_link" href="/events">

            Events

        </a>
    </div>

            </li>
            <li id="talk" class="header-nav_item js-analytics-click" data-analytics-label="talk" role="none">
                    


    <div class="header-nav_link_container">
            <a class="header-nav_link" href="/talk">

            Talk

        </a>
    </div>

            </li>
            <li id="collections" class="header-nav_item js-analytics-click" data-analytics-label="collections" role="none">
                    


    <div class="header-nav_link_container">
            <a class="header-nav_link" href="/collections">

            Collections

        </a>
    </div>

            </li>
    </ul>

            </div>
                    </div>
                        <div class="arrange_unit js-analytics-click" id="header-log-in" data-analytics-label="login">
                                    <a class="header-nav_link header-nav_link--log-in js-log-in-button" href="https://www.yelp.com/login?return_url=%2Fbiz%2Fpeter-luger-brooklyn-2">
            Log In
        </a>

                        </div>
                </div>
            </nav>
        </div>
    </div>


                                <div data-hypernova-key="yelp_main__e8b0b8210c131c668c7eb764d3e7fe7d7a4650b8__yelp_main__LSATHeaderNav__dynamic" data-hypernova-id="d833dbd3-e6d1-44b8-b21d-d39a8dc66971"><div class="lemon--div__373c0__1mboc header-nav__373c0__1X2iy border--bottom__373c0__uPbXS border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-nav_bar__373c0__2eX9n border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc container__373c0__1BzPu border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-nav_arrange__373c0__L0hhX arrange__373c0__UHqhV border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-nav_unit--left__373c0__1ZpJ- header-nav_unit--desktop__373c0__StOIx arrange-unit__373c0__1piwO border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-link__373c0__14hDE border-color--default__373c0__2oFDT" aria-haspopup="true" aria-expanded="false"><a class="lemon--a__373c0__IEZFH header-link_anchor__373c0__1nIBl" href="/search?cflt=restaurants&amp;find_loc=San+Francisco%2C+CA" tabindex="0"><div class="lemon--div__373c0__1mboc header-link_wrapper__373c0__qNNMj u-padding-r2 u-padding-l2 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG responsive-hidden-small responsive-hidden-medium border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-food header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M17.22 22a1.78 1.78 0 0 1-1.74-2.167l1.298-4.98L14 13l1.756-9.657A1.635 1.635 0 0 1 19 3.635V20.22A1.78 1.78 0 0 1 17.22 22zm-7.138-9.156l.697 7.168a1.79 1.79 0 1 1-3.56 0l.7-7.178A3.985 3.985 0 0 1 5 9V3a1 1 0 0 1 2 0v5.5c0 .28.22.5.5.5s.5-.22.5-.5V3a1 1 0 0 1 2 0v5.5c0 .28.22.5.5.5s.5-.22.5-.5V3a1 1 0 0 1 2 0v5.83c0 1.85-1.2 3.518-2.918 4.014z"/></svg></span></span><span class="lemon--span__373c0__3997G text__373c0__2pB8f header-link_text__373c0__tRzjr text-color--inherit__373c0__w_15m text-align--left__373c0__2pnx_">Restaurants</span><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG u-padding-l1 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-chevron-down header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M18.364 9.525L16.95 8.11 12 13.06 7.05 8.11 5.636 9.526 12 15.89l6.364-6.365z"/></svg></span></span></div></a></div></div><div class="lemon--div__373c0__1mboc header-nav_unit--left__373c0__1ZpJ- header-nav_unit--desktop__373c0__StOIx arrange-unit__373c0__1piwO border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-link__373c0__14hDE border-color--default__373c0__2oFDT" aria-haspopup="true" aria-expanded="false"><a class="lemon--a__373c0__IEZFH header-link_anchor__373c0__1nIBl" href="/search?cflt=homeservices&amp;find_loc=San+Francisco%2C+CA" tabindex="0"><div class="lemon--div__373c0__1mboc header-link_wrapper__373c0__qNNMj u-padding-r2 u-padding-l2 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG responsive-hidden-small responsive-hidden-medium border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-home-services header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M22.5 21H17v-6.8a.2.2 0 0 0-.2-.2h-4.6a.2.2 0 0 0-.2.2V21H3.5a.5.5 0 0 1-.5-.5v-8.15a.5.5 0 0 1 .33-.47l7.34-2.67a.5.5 0 0 0 .33-.47V4.486a.3.3 0 0 1 .434-.268l11.29 5.645a.5.5 0 0 1 .276.447V20.5a.5.5 0 0 1-.5.5zM10 14.2a.2.2 0 0 0-.2-.2H5.2a.2.2 0 0 0-.2.2v3.6c0 .11.09.2.2.2h4.6a.2.2 0 0 0 .2-.2v-3.6zm6-10.7a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v3.906L16 5V3.5zM1.273 11.424A.2.2 0 0 1 1 11.238v-1.57a.2.2 0 0 1 .127-.187l8.6-3.37a.2.2 0 0 1 .273.186v1.57a.2.2 0 0 1-.127.187l-8.6 3.375z"/></svg></span></span><span class="lemon--span__373c0__3997G text__373c0__2pB8f header-link_text__373c0__tRzjr text-color--inherit__373c0__w_15m text-align--left__373c0__2pnx_">Home Services</span><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG u-padding-l1 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-chevron-down header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M18.364 9.525L16.95 8.11 12 13.06 7.05 8.11 5.636 9.526 12 15.89l6.364-6.365z"/></svg></span></span></div></a></div></div><div class="lemon--div__373c0__1mboc header-nav_unit--left__373c0__1ZpJ- header-nav_unit--desktop__373c0__StOIx arrange-unit__373c0__1piwO border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-link__373c0__14hDE border-color--default__373c0__2oFDT" aria-haspopup="true" aria-expanded="false"><a class="lemon--a__373c0__IEZFH header-link_anchor__373c0__1nIBl" href="/search?cflt=auto&amp;find_loc=San+Francisco%2C+CA" tabindex="0"><div class="lemon--div__373c0__1mboc header-link_wrapper__373c0__qNNMj u-padding-r2 u-padding-l2 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG responsive-hidden-small responsive-hidden-medium border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-car header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M22 10.22V9.5a.5.5 0 0 0-.5-.5H19l-1.176-3.528C17.53 4.586 16.784 4 15.954 4h-7.91c-.83 0-1.573.586-1.868 1.472L5 9H2.5a.5.5 0 0 0-.5.5v.72a1 1 0 0 0 .757.97l.743.185-.3.225A3.002 3.002 0 0 0 2 14v5a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-1h12v1a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-5c0-.944-.445-1.833-1.2-2.4l-.3-.225.743-.186a1 1 0 0 0 .757-.97zM8.045 6h7.91l1.143 4H6.902l1.143-4zm-.958 9H5.182a1.182 1.182 0 0 1 0-2.363c.48 0 1.492.82 2.242 1.493a.5.5 0 0 1-.337.87zm11.73 0h-1.904a.5.5 0 0 1-.337-.87c.75-.674 1.76-1.493 2.242-1.493a1.182 1.182 0 0 1 0 2.363z"/></svg></span></span><span class="lemon--span__373c0__3997G text__373c0__2pB8f header-link_text__373c0__tRzjr text-color--inherit__373c0__w_15m text-align--left__373c0__2pnx_">Auto Services</span><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG u-padding-l1 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-chevron-down header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M18.364 9.525L16.95 8.11 12 13.06 7.05 8.11 5.636 9.526 12 15.89l6.364-6.365z"/></svg></span></span></div></a></div></div><div class="lemon--div__373c0__1mboc header-nav_unit--left__373c0__1ZpJ- header-nav_unit--desktop__373c0__StOIx arrange-unit__373c0__1piwO border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-link__373c0__14hDE border-color--default__373c0__2oFDT" aria-haspopup="true" aria-expanded="false"><a class="lemon--a__373c0__IEZFH header-link_anchor__373c0__1nIBl default-cursor__373c0__26L8T" tabindex="0"><div class="lemon--div__373c0__1mboc header-link_wrapper__373c0__qNNMj u-padding-r2 u-padding-l2 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G text__373c0__2pB8f header-link_text__373c0__tRzjr text-color--inherit__373c0__w_15m text-align--left__373c0__2pnx_">More</span><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG u-padding-l1 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-chevron-down header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M18.364 9.525L16.95 8.11 12 13.06 7.05 8.11 5.636 9.526 12 15.89l6.364-6.365z"/></svg></span></span></div></a></div></div><div class="lemon--div__373c0__1mboc header-nav_fill__373c0__2uFN5 arrange-unit__373c0__1piwO arrange-unit-fill__373c0__17z0h border-color--default__373c0__2oFDT"></div><div class="lemon--div__373c0__1mboc header-nav_unit--right__373c0__1O-Fo header-nav_unit--desktop__373c0__StOIx arrange-unit__373c0__1piwO border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-link__373c0__14hDE border-color--default__373c0__2oFDT" aria-haspopup="true" aria-expanded="false"><a class="lemon--a__373c0__IEZFH header-link_anchor__373c0__1nIBl" href="/writeareview" tabindex="0"><div class="lemon--div__373c0__1mboc header-link_wrapper__373c0__qNNMj u-padding-r3 u-padding-l3 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG responsive-hidden-small responsive-hidden-medium border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-pencil header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M20.546 4.868l-1.414-1.414a1.994 1.994 0 0 0-1.415-.586c-.51 0-1.023.195-1.414.586L4.99 14.768l-2.122 6.364 6.364-2.122L20.546 7.697a2 2 0 0 0 0-2.83zM8.152 17.262l-2.12.707.706-2.123 8.858-8.86 1.414 1.416-8.858 8.858z"/></svg></span></span><span class="lemon--span__373c0__3997G text__373c0__2pB8f header-link_text__373c0__tRzjr text-color--inherit__373c0__w_15m text-align--left__373c0__2pnx_">Write a Review</span></div></a></div></div><div class="lemon--div__373c0__1mboc header-nav_unit--right__373c0__1O-Fo header-nav_unit--desktop__373c0__StOIx arrange-unit__373c0__1piwO border-color--default__373c0__2oFDT"><div class="lemon--div__373c0__1mboc header-link__373c0__14hDE border-color--default__373c0__2oFDT" aria-haspopup="true" aria-expanded="false"><a class="lemon--a__373c0__IEZFH header-link_anchor__373c0__1nIBl" href="/advertise/consumer_header_redirect" tabindex="0"><div class="lemon--div__373c0__1mboc header-link_wrapper__373c0__qNNMj u-padding-r3 u-padding-l3 border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G display--inline__373c0__1DbOG responsive-hidden-small responsive-hidden-medium border-color--default__373c0__2oFDT"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--24-bizhouse header-link_icon__373c0__325Pm" aria-hidden="true" style="width:24px;height:24px"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" class="icon_svg"><path d="M19.125 13.375a2.874 2.874 0 0 1-2.375-1.257 2.874 2.874 0 0 1-4.75 0 2.873 2.873 0 0 1-4.75 0A2.873 2.873 0 0 1 2 10.5c0-.054.01-.107.026-.158l2-6A.5.5 0 0 1 4.5 4h15c.215 0 .406.138.475.342l2 6c.016.05.025.104.025.158a2.88 2.88 0 0 1-2.875 2.875zm-9.5 1A3.86 3.86 0 0 0 12 13.56a3.86 3.86 0 0 0 2.375.815 3.86 3.86 0 0 0 2.375-.815 3.853 3.853 0 0 0 2.25.804V20h-5v-4h-4v4H5v-5.636a3.853 3.853 0 0 0 2.25-.804 3.86 3.86 0 0 0 2.375.815z"/></svg></span></span><span class="lemon--span__373c0__3997G text__373c0__2pB8f header-link_text__373c0__tRzjr text-color--inherit__373c0__w_15m text-align--left__373c0__2pnx_">For Businesses</span></div></a></div></div></div></div></div></div></div>
<script type="application/json" data-hypernova-key="yelp_main__e8b0b8210c131c668c7eb764d3e7fe7d7a4650b8__yelp_main__LSATHeaderNav__dynamic" data-hypernova-id="d833dbd3-e6d1-44b8-b21d-d39a8dc66971"><!--{"locale":"en_US","headerConfig":{"leftLinksData":[{"iconName":"IconFood24","platform":"desktop","promo":null,"text":"Restaurants","dropdownLinksData":[[{"text":"Delivery","href":"/search?find_desc=delivery&amp;find_loc=San+Francisco%2C+CA","iconName":"IconDelivery24","gaLabel":"delivery"},{"text":"Burgers","href":"/search?cflt=burgers&amp;find_loc=San+Francisco%2C+CA","iconName":"IconBurgers24","gaLabel":"burgers"},{"text":"Chinese","href":"/search?cflt=chinese&amp;find_loc=San+Francisco%2C+CA","iconName":"IconNoodles24","gaLabel":"chinese"},{"text":"Italian","href":"/search?cflt=italian&amp;find_loc=San+Francisco%2C+CA","iconName":"IconPasta24","gaLabel":"italian"}],[{"text":"Reservations","href":"/search?find_desc=reservations&amp;find_loc=San+Francisco%2C+CA","iconName":"IconReservation24","gaLabel":"reservations"},{"text":"Japanese","href":"/search?cflt=japanese&amp;find_loc=San+Francisco%2C+CA","iconName":"IconSushi24","gaLabel":"japanese"},{"text":"Mexican","href":"/search?cflt=mexican&amp;find_loc=San+Francisco%2C+CA","iconName":"IconTaco24","gaLabel":"mexican"},{"text":"Thai","href":"/search?cflt=thai&amp;find_loc=San+Francisco%2C+CA","iconName":"IconThai24","gaLabel":"thai"}]],"alias":"restaurants","href":"/search?cflt=restaurants&amp;find_loc=San+Francisco%2C+CA"},{"iconName":"IconHomeServices24","platform":"desktop","promo":null,"text":"Home Services","dropdownLinksData":[[{"text":"Contractors","href":"/search?cflt=contractors&amp;find_loc=San+Francisco%2C+CA","iconName":"IconContractor24","gaLabel":"contractors"},{"text":"Electricians","href":"/search?cflt=electricians&amp;find_loc=San+Francisco%2C+CA","iconName":"IconElectrician24","gaLabel":"electricians"},{"text":"Home Cleaners","href":"/search?cflt=homecleaning&amp;find_loc=San+Francisco%2C+CA","iconName":"IconHomeCleaning24","gaLabel":"homecleaning"},{"text":"HVAC","href":"/search?cflt=hvac&amp;find_loc=San+Francisco%2C+CA","iconName":"IconHeatingCooling24","gaLabel":"hvac"}],[{"text":"Landscaping","href":"/search?cflt=landscaping&amp;find_loc=San+Francisco%2C+CA","iconName":"IconLandscaping24","gaLabel":"landscaping"},{"text":"Locksmiths","href":"/search?cflt=locksmiths&amp;find_loc=San+Francisco%2C+CA","iconName":"IconLocksmith24","gaLabel":"locksmiths"},{"text":"Movers","href":"/search?cflt=movers&amp;find_loc=San+Francisco%2C+CA","iconName":"IconMoving24","gaLabel":"movers"},{"text":"Plumbers","href":"/search?cflt=plumbing&amp;find_loc=San+Francisco%2C+CA","iconName":"IconPlumbers24","gaLabel":"plumbing"}]],"alias":"homeservices","href":"/search?cflt=homeservices&amp;find_loc=San+Francisco%2C+CA"},{"iconName":"IconCar24","platform":"desktop","promo":null,"text":"Auto Services","dropdownLinksData":[[{"text":"Auto Repair","href":"/search?cflt=autorepair&amp;find_loc=San+Francisco%2C+CA","iconName":"IconAutoRepair24","gaLabel":"autorepair"},{"text":"Auto Detailing","href":"/search?cflt=auto_detailing&amp;find_loc=San+Francisco%2C+CA","iconName":"IconAutoDetailing24","gaLabel":"auto_detailing"},{"text":"Body Shops","href":"/search?cflt=bodyshops&amp;find_loc=San+Francisco%2C+CA","iconName":"IconBodyShop24","gaLabel":"bodyshops"},{"text":"Car Wash","href":"/search?cflt=carwash&amp;find_loc=San+Francisco%2C+CA","iconName":"IconCarWash24","gaLabel":"carwash"}],[{"text":"Car Dealers","href":"/search?cflt=car_dealers&amp;find_loc=San+Francisco%2C+CA","iconName":"IconCarDealer24","gaLabel":"car_dealers"},{"text":"Oil Change","href":"/search?cflt=oilchange&amp;find_loc=San+Francisco%2C+CA","iconName":"IconOilChange24","gaLabel":"oilchange"},{"text":"Parking","href":"/search?cflt=parking&amp;find_loc=San+Francisco%2C+CA","iconName":"IconParking24","gaLabel":"parking"},{"text":"Towing","href":"/search?cflt=towing&amp;find_loc=San+Francisco%2C+CA","iconName":"IconTowing24","gaLabel":"towing"}]],"alias":"auto","href":"/search?cflt=auto&amp;find_loc=San+Francisco%2C+CA"},{"iconName":"","platform":"desktop","promo":null,"text":"More","dropdownLinksData":[[{"text":"Dry Cleaning","href":"/search?cflt=dryclean&amp;find_loc=San+Francisco%2C+CA","iconName":"IconDryCleaning24","gaLabel":"drycleaning"},{"text":"Phone Repair","href":"/search?cflt=mobilephonerepair&amp;find_loc=San+Francisco%2C+CA","iconName":"IconPhoneRepair24","gaLabel":"phonerepair"},{"text":"Bars","href":"/search?cflt=bars&amp;find_loc=San+Francisco%2C+CA","iconName":"IconBars24","gaLabel":"bars"},{"text":"Nightlife","href":"/search?cflt=nightlife&amp;find_loc=San+Francisco%2C+CA","iconName":"IconNightlife24","gaLabel":"nightlife"}],[{"text":"Hair Salons","href":"/search?cflt=hair&amp;find_loc=San+Francisco%2C+CA","iconName":"IconSalon24","gaLabel":"hair_salons"},{"text":"Gyms","href":"/search?cflt=gyms&amp;find_loc=San+Francisco%2C+CA","iconName":"IconGyms24","gaLabel":"gyms"},{"text":"Massage","href":"/search?cflt=massage&amp;find_loc=San+Francisco%2C+CA","iconName":"IconMassage24","gaLabel":"massage"},{"text":"Shopping","href":"/search?cflt=shopping&amp;find_loc=San+Francisco%2C+CA","iconName":"IconShopping24","gaLabel":"shopping"}]],"alias":"more","href":null}],"rightLinksData":[{"iconName":"IconPencil24","platform":"desktop","promo":null,"text":"Write a Review","dropdownLinksData":[],"alias":"writeareview","href":"/writeareview"},{"iconName":"IconBizhouse24","platform":"desktop","promo":null,"text":"For Businesses","dropdownLinksData":[],"alias":"bizcallout","href":"/advertise/consumer_header_redirect"}]},"messages":{"Elite '%{year}":"Elite '%{year}","Automatically earn cash back at %{businessName} and thousands of selected local restaurants.":"Automatically earn cash back at %{businessName} and thousands of selected local restaurants.","%{min} - %{max} mins":"%{min} - %{max} mins","Oops, we can't find your location":"Oops, we can't find your location","More business info":"More business info","Claim your business to immediately update business information, respond to reviews, and more!":"Claim your business to immediately update business information, respond to reviews, and more!","Write an update":"Write an update","<strong&gt;%{distance}</strong&gt; away from %{businessName}":"<strong&gt;%{distance}</strong&gt; away from %{businessName}","%{smart_count} Photo||||%{smart_count} Photos":"%{smart_count} Photo||||%{smart_count} Photos","Waitlist":"Waitlist","Tell ProPublica":"Tell ProPublica","Slideshow":"Slideshow","Read less":"Read less","The Local Yelp":"The Local Yelp","Sorry, one or more emoji you are using in your search term is not supported at the moment.":"Sorry, one or more emoji you are using in your search term is not supported at the moment.","Nov":"Nov","Report comment":"Report comment","Woohoo! As good as it gets!":"Woohoo! As good as it gets!","Show your Verified License prominently on Yelp":"Show your Verified License prominently on Yelp","See All Results":"See All Results","Countries":"Countries","Dine here for %{cashBack}% Cash Back":"Dine here for %{cashBack}% Cash Back","We caught someone red-handed trying to pay someone to write, change, prevent, or remove a review for this business. We weren’t fooled, but wanted you to know because these actions not only hurt consumers, but also honest businesses who play by the rules.":"We caught someone red-handed trying to pay someone to write, change, prevent, or remove a review for this business. We weren’t fooled, but wanted you to know because these actions not only hurt consumers, but also honest businesses who play by the rules.","Claim This Business":"Claim This Business","Report":"Report","Got it, thanks!":"Got it, thanks!","response time":"response time","Get Directions":"Get Directions","About Me":"About Me","+%{smart_count} Popular Dish||||+%{smart_count} Popular Dishes":"+%{smart_count} Popular Dish||||+%{smart_count} Popular Dishes","A-OK.":"A-OK.","Claim your free business page to have your changes published immediately.":"Claim your free business page to have your changes published immediately.","Sent!":"Sent!","<span class=\"%{nonEmClass}\"&gt;Businesses within</span&gt; %{containerBusinessName}":"<span class=\"%{nonEmClass}\"&gt;Businesses within</span&gt; %{containerBusinessName}","People Also Viewed":"People Also Viewed","Is this data helpful?":"Is this data helpful?","From the business owner":"From the business owner","%{rating} star rating":"%{rating} star rating","Get %{cashBack}% credited back to your card":"Get %{cashBack}% credited back to your card","Businesses with a Verified badge typically see a %{increasedLeadPercentage}% increase in leads":"Businesses with a Verified badge typically see a %{increasedLeadPercentage}% increase in leads","Try searching in a smaller area by zooming in.":"Try searching in a smaller area by zooming in.","Portfolio from the Business":"Portfolio from the Business","Yay! I'm a fan.":"Yay! I'm a fan.","Add owner reply":"Add owner reply","Maternity Care Data":"Maternity Care Data","Swipe Card":"Swipe Card","<a href=\"%{url}\"&gt;%{name}</a&gt; at this location.":"<a href=\"%{url}\"&gt;%{name}</a&gt; at this location.","Write More":"Write More","Select an option...":"Select an option...","Whoa there, your Compliment is missing a message. Personalize it with some text.":"Whoa there, your Compliment is missing a message. Personalize it with some text.","Can you provide additional information about your situation?":"Can you provide additional information about your situation?","Check the spelling or try alternate spellings.":"Check the spelling or try alternate spellings.","Did one of our links take you here?":"Did one of our links take you here?","Get directions":"Get directions","Embed This Review":"Embed This Review","Some Data By Acxiom":"Some Data By Acxiom","Categories":"Categories","<ol&gt;\n            <li&gt;At the top-right hand corner of the window, click the <b&gt;button with three dots on it</b&gt;, then <b&gt;Settings</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Choose what to clear</b&gt; underneath <b&gt;Clear browsing data</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Show more</b&gt;, then make sure only the box labeled <b&gt;Location permissions</b&gt; is checked.</li&gt;\n            <li&gt;Click <b&gt;Clear</b&gt;.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top-right hand corner of the window, click the <b&gt;button with three dots on it</b&gt;, then <b&gt;Settings</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Choose what to clear</b&gt; underneath <b&gt;Clear browsing data</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Show more</b&gt;, then make sure only the box labeled <b&gt;Location permissions</b&gt; is checked.</li&gt;\n            <li&gt;Click <b&gt;Clear</b&gt;.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;","Add Photo":"Add Photo","Investor Relations":"Investor Relations","View question details":"View question details","None of these look right? Enter your address again.":"None of these look right? Enter your address again.","Yelp users haven’t asked any questions yet about <strong&gt;%{businessName}</strong&gt;.":"Yelp users haven’t asked any questions yet about <strong&gt;%{businessName}</strong&gt;.","Deliver to":"Deliver to","Yelp Header":"Yelp Header","Page %{current_page} of %{total_pages}":"Page %{current_page} of %{total_pages}","<a %{linkAttrs}&gt;Claim this business</a&gt; to view business statistics, receive messages from prospective customers, and respond to reviews.":"<a %{linkAttrs}&gt;Claim this business</a&gt; to view business statistics, receive messages from prospective customers, and respond to reviews.","%{username} said":"%{username} said","Less relevant categories omitted.":"Less relevant categories omitted.","Owner":"Owner","You're getting %{cashBack}% Cash Back here.":"You're getting %{cashBack}% Cash Back here.","This business recently made waves in the news, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.":"This business recently made waves in the news, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.","Automatically earn cash back at %{businessName} and thousands of selected local businesses.":"Automatically earn cash back at %{businessName} and thousands of selected local businesses.","Aug":"Aug","Browsing %{displayLocation} Businesses":"Browsing %{displayLocation} Businesses","Got a question about <strong&gt;%{businessName}</strong&gt;? Ask the Yelp community!":"Got a question about <strong&gt;%{businessName}</strong&gt;? Ask the Yelp community!","Support":"Support","A business owner paid for this ad. For more information visit <a href=%{bizUrl}&gt;Yelp for Business Owners</a&gt;.":"A business owner paid for this ad. For more information visit <a href=%{bizUrl}&gt;Yelp for Business Owners</a&gt;.","Languages":"Languages","<ol&gt;\n            <li&gt;At the top of your Firefox window, to the left of the web address, you should see a <b&gt;green lock</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, you should see <b&gt;Blocked</b&gt; or <b&gt;Blocked Temporarily</b&gt; next to <b&gt;Access Your Location</b&gt;. Click the <b&gt;x</b&gt; next to this line.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://www.mozilla.org/en-US/firefox/geolocation/\" target=\"_blank\"&gt;Firefox's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top of your Firefox window, to the left of the web address, you should see a <b&gt;green lock</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, you should see <b&gt;Blocked</b&gt; or <b&gt;Blocked Temporarily</b&gt; next to <b&gt;Access Your Location</b&gt;. Click the <b&gt;x</b&gt; next to this line.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://www.mozilla.org/en-US/firefox/geolocation/\" target=\"_blank\"&gt;Firefox's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Answer":"Answer","Check your phone to view the link now!":"Check your phone to view the link now!","Choose Your Compliment Type:":"Choose Your Compliment Type:","Show fewer filters":"Show fewer filters","Showing %{start}-%{end} of %{total}":"Showing %{start}-%{end} of %{total}","Some business data by <a target=\"_blank\" rel=\"nofollow\" href=\"http://www.yellow.com.tr\" class=\"%{classes}\"&gt;Yellow.com.tr</a&gt;":"Some business data by <a target=\"_blank\" rel=\"nofollow\" href=\"http://www.yellow.com.tr\" class=\"%{classes}\"&gt;Yellow.com.tr</a&gt;","Go to Yelp for Business Owners":"Go to Yelp for Business Owners","%{smart_count} Place||||%{smart_count} Places":"%{smart_count} Place||||%{smart_count} Places","Share video":"Share video","Verification Failed":"Verification Failed","<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. If your browser doesn't ask you, try these steps:</p&gt;":"<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. If your browser doesn't ask you, try these steps:</p&gt;","By %{userName}":"By %{userName}","Try a different location":"Try a different location","License Trade":"License Trade","Photos for %{businessName}":"Photos for %{businessName}","Add A Business":"Add A Business","Select a date":"Select a date","Shop here for %{cashBack}% Cash Back":"Shop here for %{cashBack}% Cash Back","Save to a New Collection":"Save to a New Collection","Normally":"Normally","New Collection":"New Collection","About Yelp":"About Yelp","There are more restaurants on Yelp that deliver to your address":"There are more restaurants on Yelp that deliver to your address","California Contractors State License Board":"California Contractors State License Board","Business owner information":"Business owner information","%{businessOwnerRole} of %{businessName}":"%{businessOwnerRole} of %{businessName}","Try searching with another emoji.":"Try searching with another emoji.","Response Rate":"Response Rate","Oops, something has gone wrong, please try again later.":"Oops, something has gone wrong, please try again later.","Understand how a business' rating changes month-to-month. <a %{link_attrs}&gt;Learn more</a&gt;.":"Understand how a business' rating changes month-to-month. <a %{link_attrs}&gt;Learn more</a&gt;.","Remove draft":"Remove draft","Funny":"Funny","Send Feedback":"Send Feedback","Explore Delivery Restaurants":"Explore Delivery Restaurants","Details":"Details","This field is required.":"This field is required.","Try a more general search, e.g. \"pizza\" instead of \"pepperoni\"":"Try a more general search, e.g. \"pizza\" instead of \"pepperoni\"","Expiration must be in YYYY-MM-DD format.":"Expiration must be in YYYY-MM-DD format.","Yelp for Business Owners app":"Yelp for Business Owners app","Comments:":"Comments:","Terms":"Terms",", ":", ","Unclaimed":"Unclaimed","About the Business":"About the Business","Response Time":"Response Time","If you’re here to leave a review based on a first-hand experience with the business, please check back at a later date.":"If you’re here to leave a review based on a first-hand experience with the business, please check back at a later date.","Friends":"Friends","Sorry, but we didn't understand the location you entered.":"Sorry, but we didn't understand the location you entered.","Yelp Project Cost Guides":"Yelp Project Cost Guides","Ask a Question":"Ask a Question","%{businessName} also recommends":"%{businessName} also recommends","Established in %{yearEstablished}.":"Established in %{yearEstablished}.","<p&gt;This business appears to be affiliated with a group of businesses that have engaged in efforts to manipulate their reputation on Yelp. As a result, we’ve decided not to recommend any of the reviews for this business.</p&gt;<p&gt;Some of the customers of this group have reported being pressured to write positive reviews in exchange for discounts, to re-post positive reviews for affiliated businesses in other states, and to accept refunds in exchange for removing critical reviews.</p&gt;<p&gt;Alerts like these are part of Yelp’s Consumer Protection Initiative, which is designed to empower and protect consumers.</p&gt;":"<p&gt;This business appears to be affiliated with a group of businesses that have engaged in efforts to manipulate their reputation on Yelp. As a result, we’ve decided not to recommend any of the reviews for this business.</p&gt;<p&gt;Some of the customers of this group have reported being pressured to write positive reviews in exchange for discounts, to re-post positive reviews for affiliated businesses in other states, and to accept refunds in exchange for removing critical reviews.</p&gt;<p&gt;Alerts like these are part of Yelp’s Consumer Protection Initiative, which is designed to empower and protect consumers.</p&gt;","View reservation history":"View reservation history","Log Out":"Log Out","Tell us what you expected to find:":"Tell us what you expected to find:","See license information":"See license information","Find Delivery":"Find Delivery","Thank you":"Thank you","More Collections":"More Collections","%{smart_count} More Attribute||||%{smart_count} More Attributes":"%{smart_count} More Attribute||||%{smart_count} More Attributes","%{description} <span class=\"%{queryLocationClass}\"&gt;in %{displayLocation}</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;in %{displayLocation}</span&gt;","Select your rating":"Select your rating","Why do you want to report this content?":"Why do you want to report this content?","Photo of %{user}":"Photo of %{user}","Once you’ve claimed, you can:":"Once you’ve claimed, you can:","Previous review":"Previous review","Duplicate Business":"Duplicate Business","Best of Yelp":"Best of Yelp","Reviews Have Disappeared":"Reviews Have Disappeared","<p&gt;This business has been in the news in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business has been in the news in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Add info":"Add info","Enter delivery address":"Enter delivery address","This business has been claimed by the owner or a representative. ":"This business has been claimed by the owner or a representative. ","RSS":"RSS","Rating Details":"Rating Details","Ultra High-End":"Ultra High-End","We couldn't find you quickly enough! Try again later, or search near a city, place, or address instead.":"We couldn't find you quickly enough! Try again later, or search near a city, place, or address instead.","License type":"License type","Photos and Videos":"Photos and Videos","Your Email":"Your Email","<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Jan":"Jan","Please provide specific details below:":"Please provide specific details below:","Redo Search In Map":"Redo Search In Map","Developers":"Developers","Embedding Code":"Embedding Code","Oct":"Oct","Content Guidelines":"Content Guidelines","Search pages feature Yelp Ads, and your ad can appear in either slot.":"Search pages feature Yelp Ads, and your ad can appear in either slot.","Log In":"Log In","Show more filters":"Show more filters","We accept locations in the following forms:":"We accept locations in the following forms:","Oops, something went wrong. Please try again later.":"Oops, something went wrong. Please try again later.","Thanks!":"Thanks!","Check your place in line from the text we've sent to %{phoneNumber}":"Check your place in line from the text we've sent to %{phoneNumber}","Click to claim the %{cashBack}% cash back offer at %{businessName}.":"Click to claim the %{cashBack}% cash back offer at %{businessName}.","Business hours may be different today.":"Business hours may be different today.","Free Sign-up":"Free Sign-up","Offer valid: %{schedule}.":"Offer valid: %{schedule}.","<span class=\"%{nonEmClass}\"&gt;Businesses within this business</span&gt;":"<span class=\"%{nonEmClass}\"&gt;Businesses within this business</span&gt;","Results aren't relevant":"Results aren't relevant","Find Friends":"Find Friends","eeep! something went wrong.":"eeep! something went wrong.","Special hours today:":"Special hours today:","Just a Note":"Just a Note","response rate":"response rate","Yelp Blog":"Yelp Blog","Project failed to load.":"Project failed to load.","Name of Business:":"Name of Business:","About This Provider":"About This Provider","License number is required.":"License number is required.","Business website":"Business website","Uncheck All":"Uncheck All","Please tell us.":"Please tell us.","Next":"Next","NEW":"NEW","Save to Collection":"Save to Collection","Share on Facebook":"Share on Facebook","Save":"Save","Yelp WiFi":"Yelp WiFi","<a %{attrs}&gt;Add more friends</a&gt; to see them here!":"<a %{attrs}&gt;Add more friends</a&gt; to see them here!","Sorry, that's an invalid search term.":"Sorry, that's an invalid search term.","I am a customer":"I am a customer","Careers":"Careers","Unusual Activity Alert":"Unusual Activity Alert","Sorry, but we don’t take sides in factual disputes. If a review appears to reflect a user’s personal experience and opinions, it is our policy to let the user stand behind their review.":"Sorry, but we don’t take sides in factual disputes. If a review appears to reflect a user’s personal experience and opinions, it is our policy to let the user stand behind their review.","Please enter both Business Name and Business Location.":"Please enter both Business Name and Business Location.","Cool":"Cool","Cash Back":"Cash Back","Webinars":"Webinars","Great Photos":"Great Photos","Are you sure you want to remove this license? This may hide license details from your Yelp page.":"Are you sure you want to remove this license? This may hide license details from your Yelp page.","Terms of Service":"Terms of Service","Claimed":"Claimed","<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minutes":"<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minutes","This business may have tried to abuse the legal system in an effort to stifle free speech, for example through legal threats or contractual gag clauses. As a reminder, reviewers who share their experiences have a First Amendment right to express their opinions on Yelp.":"This business may have tried to abuse the legal system in an effort to stifle free speech, for example through legal threats or contractual gag clauses. As a reminder, reviewers who share their experiences have a First Amendment right to express their opinions on Yelp.","Licensee":"Licensee","<ol&gt;\n            <li&gt;Click <b&gt;Safari</b&gt; in the Menu Bar at the top of the screen, then <b&gt;Preferences</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Websites</b&gt; tab, then click <b&gt;Location</b&gt; on the left-hand panel.</li&gt;\n            <li&gt;Next to <b&gt;yelp.com</b&gt; in the right-hand panel, change the dropdown to <b&gt;Ask</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;MacOS may now prompt you to enable Location Services. If it does, follow its instructions to enable Location Services for Safari.</li&gt;\n            <li&gt;Close the Settings dialog and refresh the page. Try using Current Location search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.apple.com/en-us/HT204690\" target=\"_blank\"&gt;Safari's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;Click <b&gt;Safari</b&gt; in the Menu Bar at the top of the screen, then <b&gt;Preferences</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Websites</b&gt; tab, then click <b&gt;Location</b&gt; on the left-hand panel.</li&gt;\n            <li&gt;Next to <b&gt;yelp.com</b&gt; in the right-hand panel, change the dropdown to <b&gt;Ask</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;MacOS may now prompt you to enable Location Services. If it does, follow its instructions to enable Location Services for Safari.</li&gt;\n            <li&gt;Close the Settings dialog and refresh the page. Try using Current Location search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.apple.com/en-us/HT204690\" target=\"_blank\"&gt;Safari's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Find a Table":"Find a Table","Start Order":"Start Order","Apr":"Apr","Navigate":"Navigate","Edit review":"Edit review","Monitor how many people view your business page":"Monitor how many people view your business page","Add a license||||Add another license":"Add a license||||Add another license","Swipe your linked card at the businesses.":"Swipe your linked card at the businesses.","Special hours today.":"Special hours today.","Get %{cashBack}% Cash Back":"Get %{cashBack}% Cash Back","Current Location":"Current Location","Buy Now":"Buy Now","Live wait time:":"Live wait time:","Plumber, locksmith, electrician":"Plumber, locksmith, electrician","Account Settings":"Account Settings","These Projects are paid for and provided by the business. Yelp has not verified these Projects.":"These Projects are paid for and provided by the business. Yelp has not verified these Projects.","*":"*","Previous Project":"Previous Project","If the business you're looking for isn't here, add it!":"If the business you're looking for isn't here, add it!","Yelper names or email addresses":"Yelper names or email addresses","Check the spelling or try alternate spellings":"Check the spelling or try alternate spellings","No one is on the list. When it gets busy, join the waitlist here or on the Yelp app.":"No one is on the list. When it gets busy, join the waitlist here or on the Yelp app.","I don’t work here":"I don’t work here","Like Your Profile":"Like Your Profile","Highlights from the Business":"Highlights from the Business","Report Review":"Report Review","Meet the %{businessOwnerRole}":"Meet the %{businessOwnerRole}","Okay":"Okay","See %{smart_count} question for %{businessName}||||See all %{smart_count} questions for %{businessName}":"See %{smart_count} question for %{businessName}||||See all %{smart_count} questions for %{businessName}","Message from Yelp. Only you can see this.":"Message from Yelp. Only you can see this.","Claim Offer":"Claim Offer","Add this review to your website by copying the code below.":"Add this review to your website by copying the code below.","(no rating)":"(no rating)","Press":"Press","Need help?":"Need help?","Upcoming Reservations":"Upcoming Reservations","Send to your Phone":"Send to your Phone","Search":"Search","Claim This Free Business Page":"Claim This Free Business Page","Helpful":"Helpful","Language":"Language","You’re Funny":"You’re Funny","Or, you can&amp;nbsp;<a href=%{url}&gt;add a business here</a&gt;.":"Or, you can&amp;nbsp;<a href=%{url}&gt;add a business here</a&gt;.","Try a more general search. e.g. \"pizza\" instead of \"pepperoni\"":"Try a more general search. e.g. \"pizza\" instead of \"pepperoni\"","History":"History","Closed now":"Closed now","Sep":"Sep","Got search feedback? ":"Got search feedback? ","Share photo":"Share photo","in %{smart_count} review||||in %{smart_count} reviews":"in %{smart_count} review||||in %{smart_count} reviews","Cancel":"Cancel","Business Support":"Business Support","Show Less":"Show Less","Join the Waitlist":"Join the Waitlist","See details":"See details","Other":"Other","Your Name":"Your Name","%{smart_count} review that is not currently recommended||||%{smart_count} reviews that are not currently recommended":"%{smart_count} review that is not currently recommended||||%{smart_count} reviews that are not currently recommended","Home Address":"Home Address","Success":"Success","Sign up and claim your %{cashBack}% Cash Back offer.":"Sign up and claim your %{cashBack}% Cash Back offer.","Yelp confirmed at least one person associated with this business has the above license. Yelp cannot verify if a license covers your specific needs or that everyone at this business has a required license. Businesses pay Yelp for license verification services.":"Yelp confirmed at least one person associated with this business has the above license. Yelp cannot verify if a license covers your specific needs or that everyone at this business has a required license. Businesses pay Yelp for license verification services.","Phone number":"Phone number","Monthly Trend":"Monthly Trend","Photos":"Photos","<strong&gt;%{businessOrOwnerName} says,</strong&gt; &amp;ldquo;%{recommendationText}&amp;rdquo;":"<strong&gt;%{businessOrOwnerName} says,</strong&gt; &amp;ldquo;%{recommendationText}&amp;rdquo;","Something broke and we're not sure what. Try again later, or search near a city, place, or address instead.":"Something broke and we're not sure what. Try again later, or search near a city, place, or address instead.","Add a note":"Add a note","Is this your business?":"Is this your business?","Pending Verification":"Pending Verification","Oops! Something went wrong. Please try again.":"Oops! Something went wrong. Please try again.","Claim this business. It’s free.":"Claim this business. It’s free.","Get Verified License":"Get Verified License","Are you sure you would like to remove this unfinished review?":"Are you sure you would like to remove this unfinished review?","Savings":"Savings","Share review":"Share review","Expiration":"Expiration","This business is eligible to be claimed by a local representative in addition to corporate.":"This business is eligible to be claimed by a local representative in addition to corporate.","Messaging":"Messaging","<p&gt;This business has <a href=\"%{evidenceUrl}\"&gt;been in the news</a&gt; in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business has <a href=\"%{evidenceUrl}\"&gt;been in the news</a&gt; in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Enter your home address":"Enter your home address","%{coverCount} person||||%{coverCount} people":"%{coverCount} person||||%{coverCount} people","Claim your business to immediately update business information, track page views, and more!":"Claim your business to immediately update business information, track page views, and more!","<p&gt;This business recently participated in a film that makes false statements about Yelp and our reviewers.</p&gt;<p&gt; While some say this is the era of fake news, we still think facts matter. Businesses can certainly say what they wish, but you should learn the truth about this business’ interaction with Yelp.</p&gt;":"<p&gt;This business recently participated in a film that makes false statements about Yelp and our reviewers.</p&gt;<p&gt; While some say this is the era of fake news, we still think facts matter. Businesses can certainly say what they wish, but you should learn the truth about this business’ interaction with Yelp.</p&gt;","This business is closed on %{date}, %{holidayName}.":"This business is closed on %{date}, %{holidayName}.","Eek! Methinks not.":"Eek! Methinks not.","Cute Pic":"Cute Pic","<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Consumer Alert: Get the Facts":"Consumer Alert: Get the Facts","Why do you want to report this photo?":"Why do you want to report this photo?","%{cashBack}% Cash Back Offer Details":"%{cashBack}% Cash Back Offer Details","Check out %{reviewUser}’s review of %{businessName} on @yelp":"Check out %{reviewUser}’s review of %{businessName} on @yelp","Error":"Error","Read more":"Read more","<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended reviews":"<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended reviews","A link to the business owner’s business":"A link to the business owner’s business","More information about the action that led to this Consumer Alert is available <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.":"More information about the action that led to this Consumer Alert is available <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.","This business is a Yelp advertiser.":"This business is a Yelp advertiser.","Public":"Public","Meh. I've experienced better.":"Meh. I've experienced better.","Clear results":"Clear results","View Full Menu":"View Full Menu","Dec":"Dec","Oops! <span class=\"%{nonEmClass}\"&gt;Please try again later.</span&gt;":"Oops! <span class=\"%{nonEmClass}\"&gt;Please try again later.</span&gt;","Please provide any additional information about your license such as DBA, registered address etc.":"Please provide any additional information about your license such as DBA, registered address etc.","Are you a human? Please complete the bot challenge below.":"Are you a human? Please complete the bot challenge below.","Browsing Businesses in %{displayLocation}":"Browsing Businesses in %{displayLocation}","<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minutes":"<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minutes","Pay at selected restaurants":"Pay at selected restaurants","Consumer Alert: Questionable Legal Threats":"Consumer Alert: Questionable Legal Threats","Notifications":"Notifications","Report Video":"Report Video","Oldest First":"Oldest First","Services":"Services","Try a different location.":"Try a different location.","Also, it's possible we don't have a listing for %{findLocation}. In that case, you should try adding a zip, or try a larger nearby city.":"Also, it's possible we don't have a listing for %{findLocation}. In that case, you should try adding a zip, or try a larger nearby city.","Serving %{serviceArea} and the Surrounding Area":"Serving %{serviceArea} and the Surrounding Area","Notes":"Notes","Get access to Yelp’s free tools":"Get access to Yelp’s free tools","View Service Area":"View Service Area","%{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;","Reviews":"Reviews","Yelp Blog for Business Owners":"Yelp Blog for Business Owners","Specialties":"Specialties","Photos for %{business}":"Photos for %{business}","View more photos":"View more photos","Watch Video":"Watch Video","Save Changes":"Save Changes","Select Country Code":"Select Country Code","Order Delivery":"Order Delivery","%{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;","Rating":"Rating","<p&gt;\n            Oops! We don't recognize the web browser you're currently using. Try checking the browser's help menu, or searching the Web for instructions to turn on HTML5 Geolocation for your browser.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<p&gt;\n            Oops! We don't recognize the web browser you're currently using. Try checking the browser's help menu, or searching the Web for instructions to turn on HTML5 Geolocation for your browser.\n            You can also search near a city, place, or address instead.\n        </p&gt;","This business is closed today.":"This business is closed today.","Business Success Stories":"Business Success Stories","Send message":"Send message","Did you mean?":"Did you mean?","Email addresses":"Email addresses","Search Results Feedback":"Search Results Feedback","Please refer to our <a %{contentGuidelinesLinkAttrs}&gt;Content Guidelines</a&gt; and <a %{termsOfServiceLinkAttrs}&gt;Terms of Service</a&gt; and let us know why you think the content you’ve reported may violate these guidelines.":"Please refer to our <a %{contentGuidelinesLinkAttrs}&gt;Content Guidelines</a&gt; and <a %{termsOfServiceLinkAttrs}&gt;Terms of Service</a&gt; and let us know why you think the content you’ve reported may violate these guidelines.","This business has not yet been claimed by the owner or a representative.":"This business has not yet been claimed by the owner or a representative.","<b&gt;Your trust is our top concern,</b&gt; so businesses can’t pay to alter or remove their reviews. <a href=\"%{reviewSupportUrl}\"&gt;Learn more about reviews.</a&gt;":"<b&gt;Your trust is our top concern,</b&gt; so businesses can’t pay to alter or remove their reviews. <a href=\"%{reviewSupportUrl}\"&gt;Learn more about reviews.</a&gt;","Mo' Map":"Mo' Map","Business listing provided by":"Business listing provided by","Finish My Review":"Finish My Review","%{smart_count} Update||||%{smart_count} Updates":"%{smart_count} Update||||%{smart_count} Updates","Send":"Send","Keep Review":"Keep Review","Collections including %{businessName}":"Collections including %{businessName}","More Topics":"More Topics","<strong&gt;%{distance}</strong&gt; away from this business":"<strong&gt;%{distance}</strong&gt; away from this business","Visit Yelp for Business Owners":"Visit Yelp for Business Owners","Special hours":"Special hours","Sponsored Results":"Sponsored Results","This business has been in the news in connection with a recent tragedy, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.":"This business has been in the news in connection with a recent tragedy, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.","Sorry, the specified search area is too large.":"Sorry, the specified search area is too large.","We caught someone offering up cash, discounts, gift certificates or other incentives in exchange for reviews about this business. We wanted you to know because buying reviews not only hurts consumers, but also honest businesses who play by the rules.":"We caught someone offering up cash, discounts, gift certificates or other incentives in exchange for reviews about this business. We wanted you to know because buying reviews not only hurts consumers, but also honest businesses who play by the rules.","Hours":"Hours","Pricey":"Pricey","Show More":"Show More","Write a Review":"Write a Review","Oops, there was a problem flagging this content. Please try again later.":"Oops, there was a problem flagging this content. Please try again later.","Info":"Info","Verified by Yelp on":"Verified by Yelp on","Hide less relevant categories.":"Hide less relevant categories.","Inexpensive":"Inexpensive","<a href=\"%{firstURL}\"&gt;%{firstName}</a&gt; and <a href=\"%{secondURL}\"&gt;%{secondName}</a&gt; at this location.":"<a href=\"%{firstURL}\"&gt;%{firstName}</a&gt; and <a href=\"%{secondURL}\"&gt;%{secondName}</a&gt; at this location.","Open now":"Open now","Join Yelp Cash Back":"Join Yelp Cash Back","%{smart_count} review||||%{smart_count} reviews":"%{smart_count} review||||%{smart_count} reviews","There’s always more than one side to a story, and business owners can address misunderstandings via their Business Account by posting a public comment or sending a direct message to the reviewer.":"There’s always more than one side to a story, and business owners can address misunderstandings via their Business Account by posting a public comment or sending a direct message to the reviewer.","Privacy Policy":"Privacy Policy","Powered by":"Powered by","Make a Reservation":"Make a Reservation","By appointment only":"By appointment only","Yes":"Yes","tacos, cheap dinner, Max’s":"tacos, cheap dinner, Max’s","No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;":"No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;","(optional)":"(optional)","Page: %{page}":"Page: %{page}","Yelp Nowait":"Yelp Nowait","Show more":"Show more","Ad Choices":"Ad Choices","Closed":"Closed","Moderate":"Moderate","Open today.":"Open today.","Events":"Events","Compliment":"Compliment","Appointment":"Appointment","Jun":"Jun","A business owner paid for this ad. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.":"A business owner paid for this ad. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.","Jul":"Jul","Discount":"Discount","Collections Including %{businessName}":"Collections Including %{businessName}","Got It":"Got It","Yelp for Business Owners":"Yelp for Business Owners","Non-Public":"Non-Public","Upcoming Special Hours":"Upcoming Special Hours","No matching friends found":"No matching friends found","Great Lists":"Great Lists","No answers yet.":"No answers yet.","Add a Business":"Add a Business","Sorry, we couldn’t recognize your address.":"Sorry, we couldn’t recognize your address.","Select a Previous Order (Optional)":"Select a Previous Order (Optional)","Advertise on Yelp":"Advertise on Yelp","Work here? ":"Work here? ","Copy Code":"Copy Code","Useful":"Useful","Please enter an address to start your order.":"Please enter an address to start your order.","View More":"View More","Started on %{date}":"Started on %{date}","<b&gt;Date of experience:</b&gt; %{date_of_experience}":"<b&gt;Date of experience:</b&gt; %{date_of_experience}","Show all results.":"Show all results.","Yelping since %{year_joined} with %{num_reviews} reviews":"Yelping since %{year_joined} with %{num_reviews} reviews","Get %{cashBack}% of your bill deposited back to your card.":"Get %{cashBack}% of your bill deposited back to your card.","Your Email Address:":"Your Email Address:","Based on data from <a %{linkAttrs}&gt;%{providerName}</a&gt;":"Based on data from <a %{linkAttrs}&gt;%{providerName}</a&gt;","License #":"License #","Suggestions for improving the results:":"Suggestions for improving the results:","Link your credit or debit cards":"Link your credit or debit cards","From the business":"From the business","Finding your location...":"Finding your location...","Yelp confirmed a business or employee license.":"Yelp confirmed a business or employee license.","You're on the waitlist":"You're on the waitlist","Thank You":"Thank You","Confirmation":"Confirmation","Learn more":"Learn more","<b&gt;%{businessOwnerDisplayName}</b&gt; sent you thanks for this review":"<b&gt;%{businessOwnerDisplayName}</b&gt; sent you thanks for this review","Claim your Business Page":"Claim your Business Page","You're getting %{cashBack}% Cash Back here!":"You're getting %{cashBack}% Cash Back here!","See More Projects":"See More Projects","address, neighborhood, city, state or zip":"address, neighborhood, city, state or zip","See %{smart_count} More Service||||See %{smart_count} More Services":"See %{smart_count} More Service||||See %{smart_count} More Services","%{currentNumberLicense} of %{totalNumberLicense}":"%{currentNumberLicense} of %{totalNumberLicense}","Reg. Price":"Reg. Price","About":"About","Consumer Alert":"Consumer Alert","Check out the evidence <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.":"Check out the evidence <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.","Read Less":"Read Less","For more information visit Yelp for Business Owners.":"For more information visit Yelp for Business Owners.","Hot Stuff":"Hot Stuff","Got it":"Got it","We couldn't find an accurate position. If you're using a laptop or tablet, try moving it somewhere else and give it another go. Or, search near a city, place, or address instead.":"We couldn't find an accurate position. If you're using a laptop or tablet, try moving it somewhere else and give it another go. Or, search near a city, place, or address instead.","Good Writer":"Good Writer","%{smart_count} Review||||%{smart_count} Reviews":"%{smart_count} Review||||%{smart_count} Reviews","The restaurant is not taking waitlist parties right now":"The restaurant is not taking waitlist parties right now","Sort by":"Sort by","<strong&gt;Go mobile </strong&gt; with the <a href=%{mobileAppLinkHref}&gt; %{mobileAppLinkText} </a&gt; for iOS and Android":"<strong&gt;Go mobile </strong&gt; with the <a href=%{mobileAppLinkHref}&gt; %{mobileAppLinkText} </a&gt; for iOS and Android","<b&gt;%{smart_count}</b&gt; review||||<b&gt;%{smart_count}</b&gt; reviews":"<b&gt;%{smart_count}</b&gt; review||||<b&gt;%{smart_count}</b&gt; reviews","Saved":"Saved","Sign up to claim your %{cashBack}% Cash Back offer.":"Sign up to claim your %{cashBack}% Cash Back offer.","Edit business info":"Edit business info","Offer Claimed":"Offer Claimed","Collections":"Collections","Report Comment":"Report Comment","See all photos from %{userName} for %{businessName}":"See all photos from %{userName} for %{businessName}","A number of positive reviews for this business originated from the same IP address. Our <a href=\"%{recommendedReviewFAQUrl}\"&gt; automated recommendation software</a&gt; has taken this into account in choosing which reviews to display, but we wanted to call this to your attention because someone may be trying to artificially inflate the rating for this business.":"A number of positive reviews for this business originated from the same IP address. Our <a href=\"%{recommendedReviewFAQUrl}\"&gt; automated recommendation software</a&gt; has taken this into account in choosing which reviews to display, but we wanted to call this to your attention because someone may be trying to artificially inflate the rating for this business.","<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;":"<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;","Next Project":"Next Project","Add business hours":"Add business hours","Response time:":"Response time:","Not Helpful":"Not Helpful","Get Cash Back":"Get Cash Back","We've found multiple locations matching your search.":"We've found multiple locations matching your search.","Yelp Reservations":"Yelp Reservations","Delivery Address":"Delivery Address","Collection Name":"Collection Name","Share business":"Share business","Comment from %{businessOwnerDisplayName} of %{businessName}":"Comment from %{businessOwnerDisplayName} of %{businessName}","See More <b&gt;%{categoryName} in %{cityName}</b&gt;":"See More <b&gt;%{categoryName} in %{cityName}</b&gt;","Swipe your linked card at the restaurant.":"Swipe your linked card at the restaurant.","Messages":"Messages","Updated review":"Updated review","<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; reviews":"<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; reviews","Services Offered":"Services Offered","Business Listings provided by Yellow Pages®":"Business Listings provided by Yellow Pages®","You’re Cool":"You’re Cool","Recommended Reviews":"Recommended Reviews","Try a larger search area":"Try a larger search area","Talk":"Talk","Issued by":"Issued by","License Number":"License Number","Search within reviews":"Search within reviews","OK":"OK","Warning: Users Report Deceptive Behavior":"Warning: Users Report Deceptive Behavior","Stop following %{displayName}":"Stop following %{displayName}","Share":"Share","Posting Temporarily Blocked":"Posting Temporarily Blocked","Warning":"Warning","Help us improve.":"Help us improve.","Find":"Find","<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;":"<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;","Browsing Businesses near %{displayLocation}":"Browsing Businesses near %{displayLocation}","Location of Business: (City, State)":"Location of Business: (City, State)","No answers yet. You can be the first!":"No answers yet. You can be the first!","%{smart_count} review mentioning “%{query}”||||%{smart_count} reviews mentioning “%{query}”":"%{smart_count} review mentioning “%{query}”||||%{smart_count} reviews mentioning “%{query}”","All results shown.":"All results shown.","Are you a customer or the owner/manager of the business you'd like to add?":"Are you a customer or the owner/manager of the business you'd like to add?","See All Photos":"See All Photos","Add photos":"Add photos","<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. Also, try <a href=\"https://support.apple.com/en-us/HT204690\"&gt;enabling Location Services for your browser</a&gt;. If your browser still doesn't ask you, try these steps:</p&gt;":"<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. Also, try <a href=\"https://support.apple.com/en-us/HT204690\"&gt;enabling Location Services for your browser</a&gt;. If your browser still doesn't ask you, try these steps:</p&gt;","Oops, something went wrong.":"Oops, something went wrong.","We need your address to show you restaurants that deliver.":"We need your address to show you restaurants that deliver.","Preview":"Preview","No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;":"No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;","Related Talk Topics":"Related Talk Topics","Is this your business? Claim it for free!":"Is this your business? Claim it for free!","Buy Gift Certificate":"Buy Gift Certificate","Pagination navigation":"Pagination navigation","Toggle Menu":"Toggle Menu","Located in <a href=\"%{businessUrl}\"&gt;%{name}</a&gt;":"Located in <a href=\"%{businessUrl}\"&gt;%{name}</a&gt;","Sponsored":"Sponsored","Remove license":"Remove license","%{description} <span class=\"%{queryLocationClass}\"&gt;near %{displayLocation}</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;near %{displayLocation}</span&gt;","Lowest Rated":"Lowest Rated","Leave Waitlist":"Leave Waitlist","Enter your address to find businesses that deliver to you":"Enter your address to find businesses that deliver to you","This is my business":"This is my business","Claim this business":"Claim this business","Near":"Near","Oops, something went wrong. Please try again.":"Oops, something went wrong. Please try again.","Embed review":"Embed review","A public Collection can be openly featured on Yelp and alerts followers when you make updates. A non-public Collection can still be visible to others if you share a link to it.":"A public Collection can be openly featured on Yelp and alerts followers when you make updates. A non-public Collection can still be visible to others if you share a link to it.","Closed today.":"Closed today.","We calculate the overall star rating using only reviews that our automated software currently recommends. <a %{link_attrs}&gt;Learn more</a&gt;.":"We calculate the overall star rating using only reviews that our automated software currently recommends. <a %{link_attrs}&gt;Learn more</a&gt;.","OR":"OR","If this is your business, claim it and use Yelp’s free tools for business owners: respond to reviews, get detailed analytics and convert visitors into customers.":"If this is your business, claim it and use Yelp’s free tools for business owners: respond to reviews, get detailed analytics and convert visitors into customers.","Message":"Message","Previous":"Previous","%{reviewCount} review||||%{reviewCount} reviews":"%{reviewCount} review||||%{reviewCount} reviews","Issuing Agency":"Issuing Agency","Pay at selected businesses":"Pay at selected businesses","<a href=\"%{childrenBusinessURL}\"&gt;See businesses</a&gt; at this location":"<a href=\"%{childrenBusinessURL}\"&gt;See businesses</a&gt; at this location","Feb":"Feb","Ads":"Ads","No":"No","Known For":"Known For","DexYP logo":"DexYP logo","Oops, looks like something’s wrong. Try again!":"Oops, looks like something’s wrong. Try again!","<ol&gt;\n            <li&gt;At the top of your Opera window, near the web address, you should see a <b&gt;gray location pin</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, click <b&gt;Clear This Setting</b&gt;</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://help.opera.com/en/presto/control-pages/#geolocation\" target=\"_blank\"&gt;Opera's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top of your Opera window, near the web address, you should see a <b&gt;gray location pin</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, click <b&gt;Clear This Setting</b&gt;</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://help.opera.com/en/presto/control-pages/#geolocation\" target=\"_blank\"&gt;Opera's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Share on Twitter":"Share on Twitter","<p&gt;Did you know that local officials inspect food service facilities to improve food safety?</p&gt;<p&gt;Following a recent inspection, this facility received a food safety rating that is in the bottom 5% locally, and is categorized by inspectors as “poor”.</p&gt;<p&gt;Being in the consumer protection business, we care a lot about your safety and will display this alert for six months or until we receive a significantly improved food safety rating for this business.</p&gt;":"<p&gt;Did you know that local officials inspect food service facilities to improve food safety?</p&gt;<p&gt;Following a recent inspection, this facility received a food safety rating that is in the bottom 5% locally, and is categorized by inspectors as “poor”.</p&gt;<p&gt;Being in the consumer protection business, we care a lot about your safety and will display this alert for six months or until we receive a significantly improved food safety rating for this business.</p&gt;","Today is a holiday!":"Today is a holiday!","Business owners paid for these ads. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.":"Business owners paid for these ads. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.","This business is open as usual on %{date}, %{holidayName}.":"This business is open as usual on %{date}, %{holidayName}.","Verified by Business":"Verified by Business","Not here? Tell us what we're missing.":"Not here? Tell us what we're missing.","Businesses are too far away":"Businesses are too far away","Help Improve Yelp":"Help Improve Yelp","First to review":"First to review","Consumer Alert: Poor Food Safety Score!":"Consumer Alert: Poor Food Safety Score!","%{smart_count} person found this helpful||||%{smart_count} people found this helpful":"%{smart_count} person found this helpful||||%{smart_count} people found this helpful","<ol&gt;\n            <li&gt;At the top of your Chrome window, near the web address, click <b&gt;the gray lock icon</b&gt;.</li&gt;\n            <li&gt;In the window that pops up, make sure <b&gt;Location</b&gt; is set to <b&gt;Ask (default)</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.google.com/chrome/answer/142065?co=GENIE.Platform%3DDesktop&amp;hl=en&amp;oco=0\" target=\"_blank\"&gt;Google's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top of your Chrome window, near the web address, click <b&gt;the gray lock icon</b&gt;.</li&gt;\n            <li&gt;In the window that pops up, make sure <b&gt;Location</b&gt; is set to <b&gt;Ask (default)</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.google.com/chrome/answer/142065?co=GENIE.Platform%3DDesktop&amp;hl=en&amp;oco=0\" target=\"_blank\"&gt;Google's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Redo search when map is moved":"Redo search when map is moved","Did you mean: ":"Did you mean: ","Save Money":"Save Money","Yelp Mobile":"Yelp Mobile","Project Duration":"Project Duration","%{smart_count} other review that is not currently recommended||||%{smart_count} other reviews that are not currently recommended":"%{smart_count} other review that is not currently recommended||||%{smart_count} other reviews that are not currently recommended","Overall Rating":"Overall Rating","<b&gt;%{smart_count}</b&gt; photo||||<b&gt;%{smart_count}</b&gt; photos":"<b&gt;%{smart_count}</b&gt; photo||||<b&gt;%{smart_count}</b&gt; photos","Less Map":"Less Map","Please arrive by %{arriveBy}":"Please arrive by %{arriveBy}","<ol&gt;\n            <li&gt;Click the <b&gt;gear</b&gt; in the upper-right hand corner of the window, then <b&gt;Internet options</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Privacy</b&gt; tab in the new window that just appeared.</li&gt;\n            <li&gt;Uncheck the box labeled <b&gt;Never allow websites to request your physical location</b&gt; if it's already checked.</li&gt;\n            <li&gt;Click the button labeled <b&gt;Clear Sites</b&gt;.</li&gt;\n            <li&gt;You're good to go! Click <b&gt;OK</b&gt;, then refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;Click the <b&gt;gear</b&gt; in the upper-right hand corner of the window, then <b&gt;Internet options</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Privacy</b&gt; tab in the new window that just appeared.</li&gt;\n            <li&gt;Uncheck the box labeled <b&gt;Never allow websites to request your physical location</b&gt; if it's already checked.</li&gt;\n            <li&gt;Click the button labeled <b&gt;Clear Sites</b&gt;.</li&gt;\n            <li&gt;You're good to go! Click <b&gt;OK</b&gt;, then refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;","Why do you want to remove this review?":"Why do you want to remove this review?","Close":"Close","<b&gt;%{smart_count}</b&gt; friend||||<b&gt;%{smart_count}</b&gt; friends":"<b&gt;%{smart_count}</b&gt; friend||||<b&gt;%{smart_count}</b&gt; friends","We're sorry, the page of results you requested is unavailable.":"We're sorry, the page of results you requested is unavailable.","Yelp Sort":"Yelp Sort","Can't find a business":"Can't find a business","Ask the Community":"Ask the Community","Please select at least one reason.":"Please select at least one reason.","Response rate:":"Response rate:","This business is open %{hours} on %{date}, %{holidayName}.":"This business is open %{hours} on %{date}, %{holidayName}.","Work at %{businessName}? Claim your business":"Work at %{businessName}? Claim your business","Discover":"Discover","Blog":"Blog","To":"To","Collapse hours":"Collapse hours","All Filters":"All Filters","Sign Up":"Sign Up","Why do you want to report this review?":"Why do you want to report this review?","Don’t see your question? Ask away!":"Don’t see your question? Ask away!","See all hours":"See all hours","Remove Review":"Remove Review","Elites":"Elites","Map":"Map","Mar":"Mar","May":"May","Get License Verified":"Get License Verified","Read More":"Read More","<b&gt;Your trust is our top concern,</b&gt; so businesses can't pay to alter or remove their reviews. <a href=\"%{reviewFeedMythsUrl}\"&gt;Learn more.</a&gt;":"<b&gt;Your trust is our top concern,</b&gt; so businesses can't pay to alter or remove their reviews. <a href=\"%{reviewFeedMythsUrl}\"&gt;Learn more.</a&gt;","FAQ":"FAQ","Report Photo":"Report Photo","Report review":"Report review","Your answers help people make good choices when they’re thinking about places to go and things to do.":"Your answers help people make good choices when they’re thinking about places to go and things to do.","Newest First":"Newest First","Location &amp; Hours":"Location &amp; Hours","Highest Rated":"Highest Rated","Edit":"Edit","Remove":"Remove","Popular Dishes":"Popular Dishes","See the customer leads your business page generates":"See the customer leads your business page generates","Follow %{displayName}":"Follow %{displayName}","Remove review":"Remove review","Can&amp;rsquo;t find reviews? Read more about this topic&amp;nbsp;<a href=\"/faq#missing_user_reviews\"&gt;here</a&gt;.":"Can&amp;rsquo;t find reviews? Read more about this topic&amp;nbsp;<a href=\"/faq#missing_user_reviews\"&gt;here</a&gt;.","Add Photos":"Add Photos","Ad":"Ad","%{smart_count} recommended review||||%{smart_count} recommended reviews":"%{smart_count} recommended review||||%{smart_count} recommended reviews","License Verified":"License Verified","Verified License":"Verified License","Private":"Private","under 5 mins":"under 5 mins","No Results? Try unchecking some of the boxes above.":"No Results? Try unchecking some of the boxes above.","About This Agent":"About This Agent","Claim %{cashBack}% Cash Back Here":"Claim %{cashBack}% Cash Back Here"},"gaConfig":{"metrics":{"www":{},"global":{}},"domain":"yelp.com","user_id":null,"dimensions":{"www":{"www_second_page_pitch":[111,"status_quo"],"www_search_snippets_in_sync_with_ads":[112,"status_quo_8"],"contributions.www.remove_review_draft_modules":[183,"status_quo"],"readerx.web.biz_for_services":[116,"biz_status_quo"],"ytp_eat24_yelp_style_to_iframe":[136,"status_quo"],"viewport_tracking":[29,null],"distil":[53,null],"internal_ip":[27,"False"],"messaging.www.yelp_guaranteed":[160,"status_quo"],"styleguide_buttons":[13,"status_quo"],"account_level":[1,"anon"],"searchux.www.services_serp_card_new_layout_v0_0":[151,"status_quo"],"service":[107,"yelp-main"],"content_country":[15,"US"],"integration":[14,""],"ytp_order_confirmation_page":[143,"enabled"],"messaging.www.composer_notifications_spam_and_hover_state":[177,"status_quo"],"lsm.www.unpakt_cta_change":[108,"unpakt_cta_enabled"],"messaging.www.serp_raq_cta_coachmark":[44,"coachmark_3"],"full_url":[34,"/biz/peter-luger-brooklyn-2"],"txn.www.checkout_page_changes":[135,"enabled"],"remote_ip":[4,"136.25.5.0"],"distil_js_enabled":[138,null],"contributions.www.war_compose_redesign":[6,"enabled"],"www_biz_details_raq_sticky":[84,"enabled"],"lsat.www.dropdown_header":[54,"enabled"],"ad.web_carousel_bottom_of_biz":[78,"status_quo"],"messaging.www.show_city_in_multibiz":[77,"status_quo"],"yr_diner.www.direct_checkout_for_exact_match":[163,"disabled"],"readerx.web.popular_dishes":[198,"highlights_popular_dishes"],"contributions.www.war_compose_recent_photos":[102,"enabled"],"messaging.www.raq_cards_on_search":[61,"status_quo"],"lsmoney.www.yg_raq_cards_on_search":[152,"status_quo"],"payment.ux.www":[197,"enabled"],"review_actions_dropdown":[2,"disabled"],"lower_promoted_delivery_threshold":[58,"reduced_to_fifteen"],"platform_pickup_filter":[20,"enabled"],"www_education_banner":[190,"enabled"],"www_current_location_suggestion":[185,"enabled"],"readerx.web.biz_details_in_react":[25,"status_quo"],"pagelet_mode_www_biz_details":[24,"allow_async"],"styleguide_typography":[159,"status_quo"],"ytp_delivery_landing":[125,"platform_pages"],"referrer":[64,"none"],"contributions.www.war_compose_signup":[148,"enabled"],"is_biz_user":[129,"False"],"contributions.www.war_attach_photos":[99,"status_quo"],"www_signup_redesign":[92,"status_quo"],"nowait_restaurant.www.no_wait_message":[106,"show_waitlist_instructions"],"readerx.web.right_rail_island_refresh":[56,"white_right_rail"],"decrypted_yuv_id":[3,"DE1CA55AE6AF7DF3"],"nav_show_message_count":[59,"show_message_count"],"eat24_free_delivery_banner":[45,"disabled"]},"global":{"distil_js_enabled":[13,null],"referrer":[28,"none"],"content_country":[11,"US"],"integration":[17,""],"distil":[12,null],"internal_ip":[18,"False"],"full_url":[15,"/biz/peter-luger-brooklyn-2"],"account_level":[1,"anon"]}},"enable_high_volume_events":false,"trackers":{"www":"UA-30501-24","global":"UA-30501-1"},"js_dimensions":{"www":{"platform_order_type":[127,null],"js_vertical_search_type":[39,null]},"global":{}},"ga_enabled":true,"clientID":"DE1CA55AE6AF7DF3"},"scrollableEnabled":false}--></script>
                                

                            

                            

                            

                        <script>
            (function() {
                var main = null;

                var main=(function(){function a(b){window.yDFP.ABP={};window.yDFP.ABP.getPixelSource=function(c){return b+"?ch="+c+"&rn="+Math.random()*11};window.yDFP.ABP.detect=function(l){var g=false;var f=2;var i=false;var h=false;function k(n,m){if(f===0||m>400){n(f===0&&g)}else{setTimeout(function(){k(n,m*2)},m*2)}}function j(){if(f!==0){return}g=!i&&h}var e=function(){f-=1;j()};var d=new Image();d.onload=e;d.onerror=function(){i=true;e()};d.src=window.yDFP.ABP.getPixelSource(1);var c=new Image();c.onload=e;
c.onerror=function(){h=true;e()};c.src=window.yDFP.ABP.getPixelSource(2);k(l,100)};window.yDFP.ABP.abpInstalledChannel="1382551876";window.yDFP.ABP.abpNotInstalledChannel="2859285073"}window.yDFP=window.yDFP||{};window.yDFP.initABPDetection=a;return window.yDFP.initABPDetection})();

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main("/px.gif");
            })();
    </script>

                        <div id="print-masthead">
        <img src="https://s3-media2.fl.yelpcdn.com/assets/2/www/img/b7e9d647188d/gfx/header_print.gif" class="print-bkg-img" alt="Yelp">
    </div>


                </div>




    <div class="biz-country-us">

            <span id="page-content" class="offscreen">&nbsp;</span>

                <div class="main-content-wrap main-content-wrap--full">
    <div class="top-shelf">
        <div class="content-container js-biz-details">
            
    <div id="alert-container">

    </div>

                <div class="hidden">
            <div itemscope itemtype="http://schema.org/LocalBusiness">
        

                <div itemprop="aggregateRating" itemscope itemtype="http://schema.org/AggregateRating">
        <meta itemprop="ratingValue" content="4.0">
        <span itemprop="reviewCount">5437</span>
    </div>


            <meta itemprop="image" content="https://s3-media3.fl.yelpcdn.com/bphoto/jOoRBpCU9_2iS8z306CGOQ/ls.jpg">

            <meta itemprop="priceRange" content="Above $61" />

            <meta itemprop="name" content="Peter Luger" />

            <address itemprop="address" itemscope itemtype="http://schema.org/PostalAddress">
        <span itemprop="streetAddress">178 Broadway</span><br><span itemprop="addressLocality">Brooklyn</span>, <span itemprop="addressRegion">NY</span> <span itemprop="postalCode">11211</span><br><meta content="US" itemprop="addressCountry">
    </address>

            <span itemprop="telephone">
        (718) 387-7400
    </span>

                    <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Jordyn S.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-07-08">
        <p itemprop="description">It had been about 10 years since the last time we went. Nothing has changed which I love it!! The. Best. Steak. Ever. 

My now teens used to come when they were toddlers and they didn&#39;t use to eat, or like steaks much, but now... They killed it. 

I&#39;ve been on diet and don&#39;t eat dinner much for the past 2 years, but I did it. I ate some steaks and even dessert, a cheesecake with a whole a lot of cream on it. I shared it with my younger daughter, but we are it all!

Love the atmosphere, food and service. This is a MUST place when we visit NYC.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Guadalupe L.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="3.0">
        </div>
        <meta itemprop="datePublished" content="2019-07-05">
        <p itemprop="description">The cuts of meat are delicious, the service is good, the prices are above average, its location is somewhat removed from Manhattan, I suggest a remodeling to give it a more fresh and innovative look.

Los cortes de carne son  deliciosos, el servicio es bueno, los precios por arriba del promedio, su ubicacion es algo retirada de Manhattan, suguiero una remodelacion para darle un look mas fresco e inovativo<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Wendy S.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-07-04">
        <p itemprop="description">This is the place where beef becomes steak and steak becomes manna. After all the hoopla surrounding Peter Luger, it was without question that I had to pay a visit and digest some of their world-renowned steaks. I had made a reservation through their website a month in advance for dinner on a Thursday at 7:45 p.m. to celebrate my beloved sister&#39;s birthday. The online reservation system was very user-friendly and convenient. 

We arrived 15 minutes early and were told to wait a few minutes for our table. We didn&#39;t have to wait long and were seated next to one of the large windows overlooking Broadway Street. Our waiter was jovial, knowledgeable and efficient. He had a dry sense of humor which we found amusing. 

We started off with the Caesar salad which managed to be amazing with crunchy, fresh lettuce and grated Pecorino Romano. We ordered the Steak for Two (which is the Porterhouse steak) and the Rib Steak which has more marbling and fat. We also chose two sides, the creamed spinach and the Luger&#39;s Special German fried potatoes. Both sides are good for 2 people. 

When the food arrived, we took one look at our overladen table and predicted that leftovers would definitely result due to our ambitious ordering strategy. But against all odds and to everyone&#39;s amazement, no steak, potato or spinach was left standing. We looked with pride and a sense of accomplishment at the bare bones left on our otherwise empty plates. My sister and I both concluded that we had just experienced meat nirvana. My sister preferred the Porterhouse and I liked the Rib-eye more. Our waiter also recommended eating the steaks sans any gravy or sauce in order to fully enjoy the seasoning and core flavors of the meats.

Payment choices include cash, personal check with ID and debit card. We also got a pile of milk chocolate Peter Luger coins at the end of our incredible meal. We were too stuffed to order any desserts.

They have a tiny nook of a gift shop in the restaurant which sells Peter Luger merchandise including the Luger sauce, but one can only buy meats online. They ship all over the United States and they accept credit cards for online purchases.

I encourage people to enter the world of Peter Luger where great food and German beer-hallesque ambience abound.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Nathan A.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="4.0">
        </div>
        <meta itemprop="datePublished" content="2019-06-28">
        <p itemprop="description">Visiting from Toronto, knew we had to stop by here to try out the famous Peter Luger. Came as a party of 6, with a reservation. Even with the reservation expect to chill a bit at the front/bar while they wait and prep your table.

Between the 6 of us, we ordered the thick-cut bacon, Caesar Salad for starters, Steak for Four, Rib Steak for mains, and the potatoes and creamed Spinach for sides. Also on the table is a gravy boat of their house sauce, which to me tasted similar to a cocktail sauce, definitely has some horseradish.

Going straight to the steaks, ordering-wise there really wasn&#39;t much options so don&#39;t worry about not being able to decide what to get. The huge steaks come out on a hot plate, cut up already. The waiter will serve you a couple slices, and drizzle over some oil/butter from the plate. Good crust, strong beef taste. The tenderloin part of the porterhouse is very tender. A good amount of fat, and just a hint of aged flavour.

Service is not necessarily the best, but the waiters can be chatty if you engage with them. Price was about ~$70-80 per person with a drink.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Celia N.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-06-24">
        <p itemprop="description">Im still having trouble finding the words to express how happy and how glad I am after my second Peter Luger experience of the year.

I get reservation for 4 to celebrate my fathers birthday, and enjoying our meeting after 2 years. I asked if we could have a nice table by the window in one of the salons downstairs giving the reason why we would have that special meal, and the nailed it. 

We enjoyed an amazing meal with the best service. We started with some of their delicious breads and butter (it&#39;s hard to stop eating! They are completely addictive!) with the tender and juicy bacon and their signature sauce. And following that such an amazing bite, we had the steak for 3 and a green salad.

What can I say besides the meat was absorbed perfect, I enjoyed so much better the second time. It&#39;s so delicious that even if you&#39;re not a meat fan you will enjoy it. You can taste the level of quality in the meat, that comes without any salt. You can add it or having it with their sauce, I enjoyed it just with a tiny bit of salt to embrace the flavors.

We didn&#39;t finish hungry at all, but we needed to have the cheesecake at least, and I ordered also the pecan pie a la mode because either my parents or my sister ever had one before. 

So the service arrived, singing louder happy birthday, bringing the pie with the candle, and giving him a huge surprise (because we are not used to do the same in Spain).

We had a memorable meal, not just because of the delicious food, the place is just magical.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Sumin C.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-06-22">
        <p itemprop="description">Came here with my family and I must say this restaurant did meet my expectations. With all the hype around their steak, I&#39;ve been wanting to try Peter Luger and I was pleased with the food, service, and price. 

My parents and I got their iceberg Wedge salad, Steak for two, Creamed spinach, and for dessert, we got their Key lime pie with their homemade whipped cream. 

The salad was decent and I especially like the bacon bits. I know they are also known for their bacon and we were going to get that as an appetizer, but I&#39;m glad I got to try some of it from the salad. 

The steak was awesome -- the meat was so juicy, extremely tender, and it literally melted in my mouth. They gave us their special steak sauce, but I mostly just ate the steak by itself with some salt because the meat itself was delicious. It went well with the creamed spinach and I loved it, but my parents thought it was a bit bland. 

I&#39;m also glad I got their key lime pie for dessert even though I was so full. The pie was refreshing and tasty, and the whipped cream went well with it even though I usually don&#39;t like whipped cream. 

Overall, it was definitely one of the best steaks I&#39;ve ever had and I think it does live up to its reputation.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Michelle S.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-07-14">
        <p itemprop="description">48 hours later after eating at Peter Luger&#39;s I finally woke up from the food coma that ensued after feasting. I was warned that the food is so good it&#39;ll knock you out and that was true! I made reservations a month in advance and secured a time of 4:45pm on a Saturday. It wasn&#39;t busy and we were able to be seated when we arrived at 4:15pm. Our server was Danny and he was a very attentive server. We ordered a pellegrino (sparkling water helps digestion), bacon (2), porterhouse for 2, spinach and the potatoes. I have to say, the bacon was the best I&#39;ve ever had. I&#39;ve had bacon from all parts of the spectrum but this bacon brought me to heaven (and clogged my arteries). When the steak came, it was still sizzling in a pool of buttery goodness. My server basted the steak and then served it to me which was definitely a nice touch. The appetizers, spinach and potatoes were definitely up there as well. I&#39;ve never had creamed spinach like the one at Peter Luger&#39;s, it definitely pairs well with the steak. For the steak itself... I suggest rare because it is still cooking as it&#39;s brought to the table. I got the medium rare which was still tasty. Don&#39;t forget to drench the signature steak sauce onto it! Our bill ended up to be -$172 before tip. I&#39;ll definitely return after I take a few months to unclog my arteries. I can&#39;t wait to do it all over again.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Mark S.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-06-04">
        <p itemprop="description">When you visit NYC, you have to check out their most famous steakhouse. We made reservation 1 month ahead and I was super excited to try it out.

Dinner was at 7:45pm. When we arrived, the place was packed. After getting my name called, we get seated and a very nice gentleman gave us the menu. We knew exactly what we wanted: Porterhouse for 3 with creamed spinach and fried potatoes. 

The bread came first but pro tip: take the bread home because you&#39;re here for the steak. The steak was absolutely mind blowing. We ordered it med rare and it was cooked perfectly. It was very juicy and full of flavor. The creamed spinach was also very tasty.

The service was absolutely top notch. I will recommend this place to anyone looking for a great steak house experience.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Ron W.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-17">
        <p itemprop="description">They just signed on with Resy which allows them to ignore the 4,000 calls per day they get (which they are highly skilled at). After wrestling with Resy and losing the match I decided Resy is only good for raising your blood pressure. 

I started calling direct...for months. Yes, I got a ring but couldn&#39;t get in the ring. 

Their system is truly democratic. They can/will abuse everyone. They don&#39;t care how many Twitter followers you have or who you influence, or that you have staff to handle nasty things like this. It&#39;s the same steep climb to the summit. A table.
  
Close to yelling P.L. F. U. and throwing the phone in the pool in frustration I tried one last time and got Michael the Manager. What a Prince. &#34;You&#39;re all set&#34;, he assured me after a pleasant chat. (By the time I got there months later they had me down for a table twice the same night and also for three consecutive nights). That was a small glitch and no matter, I was in the system, in the club, and finally...in the door. 

I walked in after a 3 mile, one hour Uber from Manhattan and it was bedlam. 300 people were packed 4 deep at an old fashioned bar. (Probably the same one they had when they opened in 1837). Even though I had a reservation that just got me on the list of waiters waiting to be waited on. 

An old fashioned bar tender made me an old fashioned Dirty Martini (stirred, not shaken. Take that, Mr. Bond) and it was easily the finest Martini I&#39;ve ever enjoyed. The serious pour had me staring at the floor...from close up. 

P.L. is not pretentious. Don&#39;t bother draggin&#39; your Judith Lieber bag, or strappin&#39; on the Patek Phillipe. No one there cares. Just bring a fat bank roll because they only accept cash and that they do care about. AMEX Reward points sacrificed, it&#39;s still worth it for the Bovine bounty you&#39;re about to receive.

Be advised there are rules here, and they are strictly enforced. (A) You will suffer to get in. No walk-in&#39;s allowed. (B) You will pay cash (and tip generously). (C) You will have to wait even if you&#39;re early for your reservation. (D) You will be grilled (like a steak) to make sure your entire party is present before you&#39;re seated, (even if you&#39;re dining alone). 

So, why, you ask, am I here? Because, (A) It&#39;s Peter Luger. (B) It will exceed all the hype. (C) You&#39;re gonna have the best hunk-a-meat you&#39;ve ever had or will have. (Tinder notwithstanding.) 

The Bar: Basics. No chemistry set, periodic table concoctions. Whatever you&#39;re handed will have no peer.

The Vibe: Bucket-listers, wannabe A-listers, real A-listers, Wise Guys, and those who have chosen wisely.  

The Staff: Jerek was a Big Brother, a Dutch Uncle, Father Confessor, BFF, Master of Meat...a great guy and a top tier true pro. 

The Food: You can memorize the menu in one glance. 

Have the thick slice of house cured Bacon even if you&#39;re a devout Orthodox Kosher, Halal, Vegan. Once you eat it there can be no turning back. Close your eyes and jump, I say. 

The Caesar Salad (despite the boxed croutons, for which I forgive you, P.L.) was exceptional. Creamy and Garlic-y it was so good, in fact, the ghost of Caesar Cardini shows up periodically to get his fix.

The Porterhouse for 1, 2, 3, 4, or the whole 182nd Airborne is why you came and it will be dead perfect. Period. 

The steak arrives pre-cut on a platter with a splatter of Jus as large as an Olympic pool. (No, the steak was not swimming in it, it was just luxuriating in it.)

The meat comes direct from a 1,200 degree grill/oven and is so hot, your face will tan. Remember you&#39;re enrolled in steak Master Class. This is the way a Master serves it. 

Your server will double spoon the first few slices onto your plate, and  anoint it with Jus, (so good you&#39;ll be tempted to ask for a Go Cup of it.) You have just been served food made by God&#39;s own Chef.  

Add to that the necessary/required Creamed Spinach (deliciously more Spinach than Cream...Lawry&#39;s eat your heart out). Luger&#39;s version of this is divine. 

The Onion Rings got a flour dredge and then a flash fry. Oh my, were they fabulous? Yes, they were.  

No room for dessert? No normal human would have available space for it. And, yes you ate the entire Porterhouse, and then gnawed on the bone like a Hyenna in public. 

No matter that you can&#39;t manage dessert under any circumstances, you&#39;re gonna have the Apple Strudel under an Everest of Schlag anyway.You&#39;re gonna eat it, and you&#39;re gonna like it. I loved it. 

Walking out late required wading through 300 people 4 deep at the bar anxiously waiting for their name to be called and there&#39;s a reason for that.  

Yes, there&#39;s a reason this place is off the chart and unique. But, you don&#39;t need a reason to go. Just go. Go now!<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Rashad A.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="4.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-30">
        <p itemprop="description">Burger review: very good, tasty, patty preparation is biased towards rare (we asked for medium rare) which I personally don&#39;t mind, but figured this is good to know. You have the option to add cheese to the burger- totally up to you. I added cheese and enjoyed it but will definitely have the burger without cheese on the next run so I can better appreciate the beef. Quality bun, not too soft and has form. Slightly on the firmer side. 

Tried an order of the thick cut bacon- nice fat/meat ratio. Really melts in your mouth, but flavor is slightly lacking. 

Came on a Thursday during lunch with a friend, we had to wait about 20 minutes. Coming on a Friday or weekend expect a longer wait. Cash only.

Overall a solid burger and will return for another.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Vince L.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-25">
        <p itemprop="description">I came to NYC for this. It&#39;s been a fee years since we&#39;ve last had it and in my mind it&#39;s the best steak ever. 

This is definitely at the top of my list for best Steakhouses ever. The perfect cut, aged the right time and the finishing touch. Its full of flavour and juices. Pair it with some houSr wine and you&#39;ve got the combo ready. And dont forget to finish the bone, don&#39;t be shy. Wish I could have this every month. And dont forget to pair it with some creamed spinach and a milkshake if you still have room. 

Service is pretty good in general. Everytime we&#39;ve been the server&#39;s have been quite the jokers. This place is cash or debit only so keep that in mind. And make sure to have a reservation no matter the time . Otherwise you would be waiting in line with the tourists for a good 2 hours. It&#39;ll be a painful wait<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Lesley Y.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-16">
        <p itemprop="description">We ventured one hour to get here. It&#39;s a classic must-go steakhouse of Brooklyn / New York. When we got there, there was still a wait despite reservations. 

Make sure you get a reservation! 

We were then seated and got their prime beef steak for two. Then also got some creamed spinach.

At this point of the trip we had already eaten a lot, so the fact that I could still taste how delicious the steak is (juicy, tender, flavourful) shows how great it is.

Bill came out to be $120.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Dona F.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-08">
        <p itemprop="description">I was last here 20 years ago but I&#39;m worth it so it was time to revisit!!
No more sawdust. 
Need reservations weeks in advance 
Amazing servers who have been there for ages and know their meat!!
Just awesome steak cooked exactly how you ask. 
Side of creamer spinach - delicious 
German potatoes - cooked with butter then baked to perfect crunch. 
Apple strudel was good but the home made Schlag is even better. I can eat that plain<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Melissa G.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="4.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-12">
        <p itemprop="description">Dinner at Peter Luger exceeded my expectations. It&#39;s quite casual for a place with a Michelin star; dressing up is unnecessary, yet you won&#39;t be out of place if you want to make an occasion of it. I had heard rumors of cold service but our experience was quite the opposite; our waiter was patient and helpful -- he suggested pairing Malbec wine with our steak and it was a wonderful choice.

Although I usually prefer warm bread, the rolls that they brought to our table were delicious. I especially liked the roll filled with onion, and found it hard to pace myself knowing all the food to come! We started off with the thick-cut bacon and a caesar salad. I&#39;m not even a huge bacon person but I will say that it&#39;s something you should try here. Tastes great with the steak sauce! The caesar salad was fantastic and dressed perfectly -- probably one of my favorite parts of the meal actually. The steak though was truly amazing. So incredibly tender, definitely among the best steaks I&#39;ve had. Portion size was quite impressive! We also got the baked potato and onion rings as sides, both of which I would pass on next time. The potato was pretty standard (nothing you couldn&#39;t do at home), and the onion rings were thin and stringy. 

For dessert, we got the pecan pie with the homemade schlag of course! The pie was good although not super memorable, but we did enjoy the mound of whipped cream! To make you feel a little better about the sum you spend on this dinner, they send you away with some chocolate coins -- a nice touch. An experience very worth having.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Brian F.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-04-29">
        <p itemprop="description">Cash only

Have you ever eaten a piece of steak and thought to yourself, &#34;this is the best steak I&#39;ve ever had in my life&#34;? Well that&#39;s what happened to me at Peter Luger. This place is generally seen as one of the best steak houses in New York and I am glad I was about to treat myself here for my birthday.

I&#39;d highly recommend making reservations. You can make them online or over the phone. I actually arrived half an hour before my reserved time and got seated within a couple minutes.

I ordered the sizzling bacon ($6.95) which was smoky and fatty and delicious. It was so easy to cut through. I also got the single steak ($54.95) which was heavenly! Seriously, don&#39;t get me started on the steak... it was amazing and cooked perfectly medium rare. The crispiness and the slight char on the outside and then the juiciness and tenderness on the inside. I&#39;m salivating just thinking about it. It comes with dipping sauce which is sweet and tastes like it&#39;s ketchup based. Both of these were just melt in your mouth and I recommend getting these. I&#39;ve also heard great things about the burger but sadly they only have it for lunch so I wasn&#39;t able to try it.

Note: I&#39;ve heard of the troubles of making a reservation here. You can actually make reservations on their website now! I still recommend you make one in advance; I made mine three weeks in advanced or so and there were no problems. If you&#39;d prefer to speak to someone you can call them. If you call, I recommend calling them right when they open as I only needed to wait a couple minutes until someone answered.

Is this the best steak I&#39;ve ever had? Yup.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Phoebe W.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-29">
        <p itemprop="description">Came here with my family for a post-graduation meal! It was 100% worth the trek to Brooklyn and the 8:45 PM reservation time (make your reservations well in advance, because this was the earliest I could get 2 months in advance for a Wednesday evening!).

The four of us ordered the steak for two, the creamed spinach, the bacon, and the onion rings. I can&#39;t even describe just how good everything was -- that bacon was unlike anything else I&#39;d ever put into my mouth. The steak too -- nothing is as good as the first bite of steak, and paired with the Peter Luger&#39;s sauce, was a dish I will never forget. Creamed spinach -- I thought it was good, my family thought it was amazing. Even my sister and dad who hate spinach ate it by the spoonful. The onion rings were fresh and hot and tasty as well, a little bit on the greasy side but that&#39;s what makes them taste so good the next day.

Our server was so funny and attentive, he made us feel welcomed from the start. My dad left his backpack at the restaurant on accident and they went above and beyond in helping him get it back the following day. Definitely a great splurge-worthy meal and experience!<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Jae H.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="4.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-09">
        <p itemprop="description">I made reservation for my friends whom visited in NY from LA. One of my friend birthday coming up and I decided to took my friends at this place. I&#39;ve been tried this place but I could tell you that their prime steak was the best I ever had in this place before. To honestly I like this place but My experience wasn&#39;t good in here. 

Stopped by for dinner with friends. We ordered 3 of prime steak and side of creamy spinach, mixed green salad. 3 of cosmopolitan cocktails, one of bartender special cocktail and key lime pie with coffee, vanilla ice cream for dessert.

Prime steak(4.5/5) was comes out hot also tasted great than I expected. Very juicy and tender, soft. Peter Luger steak sauce is famous. I loved their sauce and my all of friends were looks like they love steak with sauce. 

Creamy spinach(2/5)wasn&#39;t creamy than I expected. I&#39;ve had eat better elsewhere. Not creamy much but savory. If you like less creamy style, you would like it.

Mixed green salad(3/5) was ok but very fresh vegetables. Nothing special, you can eat this salad at your home too. 

Key lime pie(3/5)was good but I&#39;ve eat better elsewhere. I am a huge fan of dessert also I love key lime pie. I eaten that in my whole life but this place key lime pie wasn&#39;t the best. Home made whipped cream was the best. Not too sweet but kinda of savory with very soft. Even if you don&#39;t like sweets, you can try key lime pie and their home made whipped cream. But I would recommend you that if you want to eat key lime pie, drink coffee. 

Vanilla ice cream (2/5) just ok. 

Cosmopolitan(2/5) were really strong cocktails. Actually it isn&#39;t strong cocktails but the bartender made our cocktails very strong. I had to request them to extra cranberry juice. 

Service(3/5)-our waiter was accommodating and very friendly. They load up our plate with steak and side. But there&#39;s busy and it was difficult to get for what we need. Very little of space between table to table. One waiter was keeping touch my chair or moved my chair by accidentally but nothing apologized. I understand how they&#39;re busy and not enough space for served but i personally thought he could say at least sorry or excuse me. Am I too picky? 

Ambience(2.5/5)- old restaurant and their interior is wood style. It&#39;s simple interior but I couldn&#39;t see any of special point of interior. 

This place is cash only. If you have plan for go this place then have some cash. No dress cord here.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Danny L.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="5.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-14">
        <p itemprop="description">Where to start....I have been live in nyc for almost a decade and this is the FIRST TIME I came to Peter Luger, I have to say I absolutely love it !!! I know this steakhouse is very famous, but I didn&#39;t expect the fact that I have to make a reservation one week ahead. However it all worth it at the end! 

It was PACKED when we get here and people already waiting in the line. The interior design is old fashioned, classic steakhouse style which I love. 

We ordered their famous bacon strips, porterhouse steak for two, wedge salad, cream Spanish , French fries, and pecan pie with their homemade whip cream. Everything tasted sooo delicious. Especially the thick cut bacon strip with their house steak Sauce.. it was like in heaven... hahaha.  The steak was cooked beautifully , tender and juicy, the sound of the sizzling pan make it extra mouth watering... My friend is not a dessert lover, but she was to able to finish the whole pecan pie herself...  

Food price is very reasonable too, we (two people)spent around $300 + tip. Which is very normal for a good steak house in nyc.

Highly recommended !!!!<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Gab G.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="3.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-12">
        <p itemprop="description">Made a reservation over the phone for my boyfriend&#39;s birthday about a month in advance. Service was great, but unfortunately we found the food to be hit or miss. Rib eye, Caesar salad, creamed spinach and apple strudel were all great, but the German potatoes were dry and the key lime pie had a weird chemical taste to it. On top of that, the dishes we enjoyed were just that - enjoyable. At a place like this, I want to be blown away when I take a bite. That wasn&#39;t really the case here, and so I think I am better off returning to places like 4 Charles and Del Frisco&#39;s to get my meat and potatoes fix. This is a classic place that I&#39;m glad I tried once, but I&#39;m not sure I&#39;d recommend it if you&#39;re looking for a singular NYC steak experience. 

3.5 stars.<p>
    </div>

            <div itemprop="review" itemscope itemtype="http://schema.org/Review">
        <meta itemprop="author" content="Yonnie C.">
        <div itemprop="reviewRating" itemscope itemtype="http://schema.org/Rating">
            <meta itemprop="ratingValue" content="4.0">
        </div>
        <meta itemprop="datePublished" content="2019-05-08">
        <p itemprop="description">Ordered Steak for 2. Asked for medium but got medium rare. Because the plate was so hot, we waited for the steak on the plate to baste itself more before consuming. The porterhouse steak is on the leaner side, not as marbled as we would like, but is still good. 

Recommend to make a reservation, lines for walk in gets long. No credit cards, cash or debit.

(This review written by husband)<p>
    </div>


    </div>

    </div>



    

    <div class="biz-page-header clearfix">
        <div class="biz-page-header-left claim-status">

            

    


    <div class="u-space-t1">
        <h1 class="biz-page-title embossed-text-white shortenough">Peter</h1>
        <div class=u-inline-block>
            <h1 class="biz-page-title embossed-text-white shortenough">Luger</h1>
            <div class="u-inline-block">
                                            <div class="u-nowrap claim-status_teaser js-claim-status-hover">
        <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-checkmark-badged icon--size-18 icon--blue-dark claim-status_icon u-space-r1 claim-status_icon--claimed">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_checkmark_badged" />
    </svg>
</span>Claimed
    </div>


                        <div class="u-hidden js-claim-status-hover-content">
        <p class="u-space-b0">
            This business has been claimed by the owner or a representative.
                <a class="js-biz-page-claim-link" href="https://www.yelp-support.com/article/000032392?l=en_US" target="_blank">Learn more</a>
        </p>
    </div>



            </div>
        </div>
    </div>

    

    <div class="biz-main-info embossed-text-white">
            <div class="rating-info clearfix">
                        <div class="biz-rating biz-rating-very-large clearfix" >
                



    <div class="i-stars i-stars--large-4 rating-very-large" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>


                    <span class="review-count rating-qualifier">
            5437 reviews
    </span>

        </div>

                        <div id="rating-details-modal-content" class="hidden" data-monthly-ratings="{&#34;2016&#34;: [[0, 4.0], [1, 4.0], [2, 4.0], [3, 4.0], [4, 4.0], [5, 4.0], [6, 4.0], [7, 4.0], [8, 4.0], [9, 4.0], [10, 4.0], [11, 4.0]], &#34;2017&#34;: [[0, 4.0], [1, 4.0], [2, 4.0], [3, 4.0], [4, 4.0], [5, 4.5], [6, 4.0], [7, 4.0], [8, 4.0], [9, 4.0], [10, 4.0], [11, 4.5]], &#34;2018&#34;: [[0, 4.0], [1, 4.0], [2, 4.0], [3, 4.0], [4, 4.5], [5, 4.0], [6, 4.0], [7, 4.0], [8, 4.0], [9, 4.0], [10, 4.0], [11, 4.5]], &#34;2019&#34;: [[0, 4.5], [1, 4.0], [2, 4.0], [3, 4.0], [4, 4.0], [5, 4.0], [6, 4.0]], &#34;2015&#34;: [[0, 4.0], [1, 4.0], [2, 4.0], [3, 4.0], [4, 4.0], [5, 4.0], [6, 4.5], [7, 4.0], [8, 4.0], [9, 4.0], [10, 4.0], [11, 4.0]]}">
    <div class="modal_head">
            <h2>Rating Details</h2>
    </div>

    <div class="modal_body">
                <div class="rating-details-graph-block ysection">
                    <div class="arrange arrange--middle section-header">
                        <p class="arrange_unit arrange_unit--fill">
                            <strong>Monthly Trend</strong>
                        </p>
                        <div class="arrange_unit">
                                



        <div class="tab-nav-container">
            <ul class="tab-nav js-tab-nav" role="tablist">
                    
            <li class="tab-nav_item" role="presentation">
                        





    <a aria-selected="true" class="tab-link js-tab-link tab-link--nav js-tab-link--nav is-selected" data-year="2019" href="javascript:;" role="tab"
    >
                <span class="tab-link_label">2019</span>
    </a>


            </li>



            <li class="tab-nav_item" role="presentation">
                        





    <a aria-selected="false" class="tab-link js-tab-link tab-link--nav js-tab-link--nav" data-year="2018" href="javascript:;" role="tab"
    >
                <span class="tab-link_label">2018</span>
    </a>


            </li>



            <li class="tab-nav_item" role="presentation">
                        





    <a aria-selected="false" class="tab-link js-tab-link tab-link--nav js-tab-link--nav" data-year="2017" href="javascript:;" role="tab"
    >
                <span class="tab-link_label">2017</span>
    </a>


            </li>



            <li class="tab-nav_item" role="presentation">
                        





    <a aria-selected="false" class="tab-link js-tab-link tab-link--nav js-tab-link--nav" data-year="2016" href="javascript:;" role="tab"
    >
                <span class="tab-link_label">2016</span>
    </a>


            </li>



            <li class="tab-nav_item tab-nav_item--last" role="presentation">
                        





    <a aria-selected="false" class="tab-link js-tab-link tab-link--nav js-tab-link--nav" data-year="2015" href="javascript:;" role="tab"
    >
                <span class="tab-link_label">2015</span>
    </a>


            </li>



            </ul>
        </div>


                        </div>
                    </div>
                    <div class="monthly-avg-graph-container graph-container u-hide-overflow">
                        <div class="monthly-avg-graph-placeholder graph-placeholder"></div>
                    </div>
                    <small class="u-block u-space-t1">
                        Understand how a business’ rating changes month-to-month. <a href="https://www.yelp-support.com/article/How-is-the-Monthly-Trend-of-a-business-s-rating-calculated?l=en_US">Learn more</a>.
                    </small>
                </div>

                <div class="section-header"><h4>Overall Rating</h4></div>
                <p class="rating-details-ratings-info">Yelping since 2005 with 5437 reviews</p>

                    <table class="histogram histogram--alternating histogram--large">

            <tr class="histogram_row histogram_row--1">
                <th scope="row" class="histogram_label nowrap">
                    5 stars
                </th>
                <td>
                    <table>
                        <tr>
                            <td style="width: 100%;">
                                    <div class="histogram_bar"></div>
                            </td>
                            <td class="histogram_count">2725</td>
                        </tr>
                    </table>
                </td>
            </tr>

            <tr class="histogram_row histogram_row--2">
                <th scope="row" class="histogram_label nowrap">
                    4 stars
                </th>
                <td>
                    <table>
                        <tr>
                            <td style="width: 45%;">
                                    <div class="histogram_bar"></div>
                            </td>
                            <td class="histogram_count">1239</td>
                        </tr>
                    </table>
                </td>
            </tr>

            <tr class="histogram_row histogram_row--3">
                <th scope="row" class="histogram_label nowrap">
                    3 stars
                </th>
                <td>
                    <table>
                        <tr>
                            <td style="width: 28%;">
                                    <div class="histogram_bar"></div>
                            </td>
                            <td class="histogram_count">760</td>
                        </tr>
                    </table>
                </td>
            </tr>

            <tr class="histogram_row histogram_row--4">
                <th scope="row" class="histogram_label nowrap">
                    2 stars
                </th>
                <td>
                    <table>
                        <tr>
                            <td style="width: 14%;">
                                    <div class="histogram_bar"></div>
                            </td>
                            <td class="histogram_count">380</td>
                        </tr>
                    </table>
                </td>
            </tr>

            <tr class="histogram_row histogram_row--5">
                <th scope="row" class="histogram_label nowrap">
                    1 star
                </th>
                <td>
                    <table>
                        <tr>
                            <td style="width: 12%;">
                                    <div class="histogram_bar"></div>
                            </td>
                            <td class="histogram_count">333</td>
                        </tr>
                    </table>
                </td>
            </tr>
    </table>



            <small>We calculate the overall star rating using only reviews that our automated software currently recommends. <a href="https://www.yelp-support.com/article/What-is-Yelp-s-recommendation-software?l=en_US">Learn more</a>.</small>
    </div>
    </div>

                    
    <div class="rating-details">
        <a href="javascript:;" class="chiclet-link chiclet-link--with-text show-tooltip js-rating-details">
            <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-histogram icon--size-14 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_histogram" />
    </svg>
</span>         Details
        <span class="offscreen">, Opens a popup</span>

            <span class="tooltip-wrapper">
                <span class="tooltip">Rating details</span>
            </span>
        </a>
    </div>

            </div>

        <div class="price-category">
                        <span class="bullet-after">
            
        <span class="business-attribute price-range">$$$$</span>
        </span>
            <span class="category-str-list">
                    <a href="/c/brooklyn/steak">Steakhouses</a>
    </span>




                    <a class="edit-business-button chiclet-link chiclet-link--with-text" data-ro-mode-action="Edit business info" href="/biz_attribute?biz_id=4yPqqJDJOQX69gC66YUDkA">
        Edit
    </a>

        </div>
    </div>

            <script>
            (function() {
                var main = null;

                var main=function(c){var d=window.performance||window.mozPerformance||window.msPerformance||window.webkitPerformance;if(d){var a=d.mark||d.mozPerformance||d.msMark||d.webkitMark;var b=d.clearMarks||d.mozClearMarks||d.msClearMarks||d.webkitClearMarks;if(typeof a==="function"&&typeof b==="function"){b.call(d,c);a.call(d,c)}}};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main("yelp_header_title");
            })();
    </script>
        </div>
        <div class="biz-page-header-right u-relative">
                <div class="biz-page-actions nowrap">
        


    <a class="ybtn ybtn--primary war-button" href="/writeareview/biz/4yPqqJDJOQX69gC66YUDkA?return_url=%2Fbiz%2F4yPqqJDJOQX69gC66YUDkA&amp;source=biz_details_war_button">
            <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-star icon--size-24 icon--currentColor u-space-r-half icon--fallback-inverted">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_star" />
    </svg>
</span>Write a Review
    </a>




    <span class="ybtn-group clearfix js-save-to-collection u-flex" data-save-text="Save" data-saved-text="Saved" data-error-text="Oops, something went wrong.">

    <a class="ybtn ybtn--small ybtn--secondary add-photo-button" href="/biz_user_photos/4yPqqJDJOQX69gC66YUDkA/upload">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-add-photo icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_add_photo" />
    </svg>
</span>
            Add Photo
    </a>

        


        <div class="send-to-friend-wrapper">
            <div data-hypernova-key="yelp_main__e8b0b8210c131c668c7eb764d3e7fe7d7a4650b8__yelp_main__SendToFriendButton__dynamic" data-hypernova-id="42841392-24ce-4aa2-9efa-b35dd4efeaa5"><button class="button__373c0__3yl_n small__373c0__1sSnK secondary__373c0__3Q2Us action-bar-button__373c0__3VxNd" style="--mousedown-x:0px;--mousedown-y:0px;--button-width:0px" type="submit" value="submit"><span class="lemon--span__373c0__3997G text__373c0__2pB8f text-color--inherit__373c0__w_15m text-align--center__373c0__2vkq2 text-size--inherit__373c0__2gFQ3"><span class="lemon--span__373c0__3997G icon__373c0__ehCWV icon--18-share" aria-hidden="true" style="width:18px;height:18px"><svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 18 18" class="icon_svg"><path d="M17.714 6.43L13 10.356v-3.03c-1 0-5.097 1.47-6.286 3.62.274-3.08 4.286-5.5 6.286-5.5V2.5l4.714 3.93zM3 4v10h11v-2.5l1-1V15H2V3h8.5l-1 1H3z"/></svg></span> <!-- -->Share</span></button></div>
<script type="application/json" data-hypernova-key="yelp_main__e8b0b8210c131c668c7eb764d3e7fe7d7a4650b8__yelp_main__SendToFriendButton__dynamic" data-hypernova-id="42841392-24ce-4aa2-9efa-b35dd4efeaa5"><!--{"locale":"en_US","backendContext":{"context":[{"schema_id":279661,"payload_data":{"guv":"DE1CA55AE6AF7DF3"}},{"schema_id":171864,"payload_data":{"unique_request_id":"5e84cbeca1806274"}},{"schema_id":280941,"payload_data":{"internal_ip":false}},{"schema_id":280599,"payload_data":{"country":"US","language":"en"}},{"schema_id":283288,"payload_data":{"user_id":null}},{"schema_id":311701,"payload_data":{"platform":"www"}},{"schema_id":284174,"payload_data":{"normed_app_version":null}},{"schema_id":284175,"payload_data":{"normed_device_type":null}},{"schema_id":312263,"payload_data":{"product":"consumer","product_version":"re8b0b8210c-deploy-safe-safe-safe"}},{"schema_id":309536,"payload_data":{"unique_view_id":"5e84cbeca1806274","previous_unique_view_id":null}},{"schema_id":311448,"payload_data":{"biz_user_id_encid":null}},{"schema_id":311702,"payload_data":{"os_version":"x.x.x","os_name":"unknown"}},{"schema_id":306745,"payload_data":{"name":"www"}},{"schema_id":311859,"payload_data":{"page_id":"9c8b6d9f740d50f59a4c50db8f885ad31447f635"}},{"schema_id":311757,"payload_data":{"is_interactive":true}},{"schema_id":311510,"payload_data":{"user_id_encid":null}},{"schema_id":311758,"payload_data":{"interface_version":"x.x.x","interface_name":"unknown"}},{"schema_id":312281,"payload_data":{"offset":null}},{"schema_id":311756,"payload_data":{"hardware_model":null}}]},"messages":{"Elite '%{year}":"Elite '%{year}","Automatically earn cash back at %{businessName} and thousands of selected local restaurants.":"Automatically earn cash back at %{businessName} and thousands of selected local restaurants.","%{min} - %{max} mins":"%{min} - %{max} mins","Oops, we can't find your location":"Oops, we can't find your location","More business info":"More business info","Claim your business to immediately update business information, respond to reviews, and more!":"Claim your business to immediately update business information, respond to reviews, and more!","Write an update":"Write an update","<strong&gt;%{distance}</strong&gt; away from %{businessName}":"<strong&gt;%{distance}</strong&gt; away from %{businessName}","%{smart_count} Photo||||%{smart_count} Photos":"%{smart_count} Photo||||%{smart_count} Photos","Waitlist":"Waitlist","Tell ProPublica":"Tell ProPublica","Slideshow":"Slideshow","Read less":"Read less","The Local Yelp":"The Local Yelp","Sorry, one or more emoji you are using in your search term is not supported at the moment.":"Sorry, one or more emoji you are using in your search term is not supported at the moment.","Nov":"Nov","Report comment":"Report comment","Woohoo! As good as it gets!":"Woohoo! As good as it gets!","Show your Verified License prominently on Yelp":"Show your Verified License prominently on Yelp","See All Results":"See All Results","Countries":"Countries","Dine here for %{cashBack}% Cash Back":"Dine here for %{cashBack}% Cash Back","We caught someone red-handed trying to pay someone to write, change, prevent, or remove a review for this business. We weren’t fooled, but wanted you to know because these actions not only hurt consumers, but also honest businesses who play by the rules.":"We caught someone red-handed trying to pay someone to write, change, prevent, or remove a review for this business. We weren’t fooled, but wanted you to know because these actions not only hurt consumers, but also honest businesses who play by the rules.","Claim This Business":"Claim This Business","Report":"Report","Got it, thanks!":"Got it, thanks!","response time":"response time","Get Directions":"Get Directions","About Me":"About Me","+%{smart_count} Popular Dish||||+%{smart_count} Popular Dishes":"+%{smart_count} Popular Dish||||+%{smart_count} Popular Dishes","A-OK.":"A-OK.","Claim your free business page to have your changes published immediately.":"Claim your free business page to have your changes published immediately.","Sent!":"Sent!","<span class=\"%{nonEmClass}\"&gt;Businesses within</span&gt; %{containerBusinessName}":"<span class=\"%{nonEmClass}\"&gt;Businesses within</span&gt; %{containerBusinessName}","People Also Viewed":"People Also Viewed","Is this data helpful?":"Is this data helpful?","From the business owner":"From the business owner","%{rating} star rating":"%{rating} star rating","Get %{cashBack}% credited back to your card":"Get %{cashBack}% credited back to your card","Businesses with a Verified badge typically see a %{increasedLeadPercentage}% increase in leads":"Businesses with a Verified badge typically see a %{increasedLeadPercentage}% increase in leads","Try searching in a smaller area by zooming in.":"Try searching in a smaller area by zooming in.","Portfolio from the Business":"Portfolio from the Business","Yay! I'm a fan.":"Yay! I'm a fan.","Add owner reply":"Add owner reply","Maternity Care Data":"Maternity Care Data","Swipe Card":"Swipe Card","<a href=\"%{url}\"&gt;%{name}</a&gt; at this location.":"<a href=\"%{url}\"&gt;%{name}</a&gt; at this location.","Write More":"Write More","Select an option...":"Select an option...","Whoa there, your Compliment is missing a message. Personalize it with some text.":"Whoa there, your Compliment is missing a message. Personalize it with some text.","Can you provide additional information about your situation?":"Can you provide additional information about your situation?","Check the spelling or try alternate spellings.":"Check the spelling or try alternate spellings.","Did one of our links take you here?":"Did one of our links take you here?","Get directions":"Get directions","Embed This Review":"Embed This Review","Some Data By Acxiom":"Some Data By Acxiom","Categories":"Categories","<ol&gt;\n            <li&gt;At the top-right hand corner of the window, click the <b&gt;button with three dots on it</b&gt;, then <b&gt;Settings</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Choose what to clear</b&gt; underneath <b&gt;Clear browsing data</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Show more</b&gt;, then make sure only the box labeled <b&gt;Location permissions</b&gt; is checked.</li&gt;\n            <li&gt;Click <b&gt;Clear</b&gt;.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top-right hand corner of the window, click the <b&gt;button with three dots on it</b&gt;, then <b&gt;Settings</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Choose what to clear</b&gt; underneath <b&gt;Clear browsing data</b&gt;.</li&gt;\n            <li&gt;Click <b&gt;Show more</b&gt;, then make sure only the box labeled <b&gt;Location permissions</b&gt; is checked.</li&gt;\n            <li&gt;Click <b&gt;Clear</b&gt;.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;","Add Photo":"Add Photo","Investor Relations":"Investor Relations","View question details":"View question details","None of these look right? Enter your address again.":"None of these look right? Enter your address again.","Yelp users haven’t asked any questions yet about <strong&gt;%{businessName}</strong&gt;.":"Yelp users haven’t asked any questions yet about <strong&gt;%{businessName}</strong&gt;.","Deliver to":"Deliver to","Yelp Header":"Yelp Header","Page %{current_page} of %{total_pages}":"Page %{current_page} of %{total_pages}","<a %{linkAttrs}&gt;Claim this business</a&gt; to view business statistics, receive messages from prospective customers, and respond to reviews.":"<a %{linkAttrs}&gt;Claim this business</a&gt; to view business statistics, receive messages from prospective customers, and respond to reviews.","%{username} said":"%{username} said","Less relevant categories omitted.":"Less relevant categories omitted.","Owner":"Owner","You're getting %{cashBack}% Cash Back here.":"You're getting %{cashBack}% Cash Back here.","This business recently made waves in the news, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.":"This business recently made waves in the news, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.","Automatically earn cash back at %{businessName} and thousands of selected local businesses.":"Automatically earn cash back at %{businessName} and thousands of selected local businesses.","Aug":"Aug","Browsing %{displayLocation} Businesses":"Browsing %{displayLocation} Businesses","Got a question about <strong&gt;%{businessName}</strong&gt;? Ask the Yelp community!":"Got a question about <strong&gt;%{businessName}</strong&gt;? Ask the Yelp community!","Support":"Support","A business owner paid for this ad. For more information visit <a href=%{bizUrl}&gt;Yelp for Business Owners</a&gt;.":"A business owner paid for this ad. For more information visit <a href=%{bizUrl}&gt;Yelp for Business Owners</a&gt;.","Languages":"Languages","<ol&gt;\n            <li&gt;At the top of your Firefox window, to the left of the web address, you should see a <b&gt;green lock</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, you should see <b&gt;Blocked</b&gt; or <b&gt;Blocked Temporarily</b&gt; next to <b&gt;Access Your Location</b&gt;. Click the <b&gt;x</b&gt; next to this line.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://www.mozilla.org/en-US/firefox/geolocation/\" target=\"_blank\"&gt;Firefox's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top of your Firefox window, to the left of the web address, you should see a <b&gt;green lock</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, you should see <b&gt;Blocked</b&gt; or <b&gt;Blocked Temporarily</b&gt; next to <b&gt;Access Your Location</b&gt;. Click the <b&gt;x</b&gt; next to this line.</li&gt;\n            <li&gt;You're good to go! Refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://www.mozilla.org/en-US/firefox/geolocation/\" target=\"_blank\"&gt;Firefox's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Answer":"Answer","Check your phone to view the link now!":"Check your phone to view the link now!","Choose Your Compliment Type:":"Choose Your Compliment Type:","Show fewer filters":"Show fewer filters","Showing %{start}-%{end} of %{total}":"Showing %{start}-%{end} of %{total}","Some business data by <a target=\"_blank\" rel=\"nofollow\" href=\"http://www.yellow.com.tr\" class=\"%{classes}\"&gt;Yellow.com.tr</a&gt;":"Some business data by <a target=\"_blank\" rel=\"nofollow\" href=\"http://www.yellow.com.tr\" class=\"%{classes}\"&gt;Yellow.com.tr</a&gt;","Go to Yelp for Business Owners":"Go to Yelp for Business Owners","%{smart_count} Place||||%{smart_count} Places":"%{smart_count} Place||||%{smart_count} Places","Share video":"Share video","Verification Failed":"Verification Failed","<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. If your browser doesn't ask you, try these steps:</p&gt;":"<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. If your browser doesn't ask you, try these steps:</p&gt;","By %{userName}":"By %{userName}","Try a different location":"Try a different location","License Trade":"License Trade","Photos for %{businessName}":"Photos for %{businessName}","Add A Business":"Add A Business","Select a date":"Select a date","Shop here for %{cashBack}% Cash Back":"Shop here for %{cashBack}% Cash Back","Save to a New Collection":"Save to a New Collection","Normally":"Normally","New Collection":"New Collection","About Yelp":"About Yelp","There are more restaurants on Yelp that deliver to your address":"There are more restaurants on Yelp that deliver to your address","California Contractors State License Board":"California Contractors State License Board","Business owner information":"Business owner information","%{businessOwnerRole} of %{businessName}":"%{businessOwnerRole} of %{businessName}","Try searching with another emoji.":"Try searching with another emoji.","Response Rate":"Response Rate","Oops, something has gone wrong, please try again later.":"Oops, something has gone wrong, please try again later.","Understand how a business' rating changes month-to-month. <a %{link_attrs}&gt;Learn more</a&gt;.":"Understand how a business' rating changes month-to-month. <a %{link_attrs}&gt;Learn more</a&gt;.","Remove draft":"Remove draft","Funny":"Funny","Send Feedback":"Send Feedback","Explore Delivery Restaurants":"Explore Delivery Restaurants","Details":"Details","This field is required.":"This field is required.","Try a more general search, e.g. \"pizza\" instead of \"pepperoni\"":"Try a more general search, e.g. \"pizza\" instead of \"pepperoni\"","Expiration must be in YYYY-MM-DD format.":"Expiration must be in YYYY-MM-DD format.","Yelp for Business Owners app":"Yelp for Business Owners app","Comments:":"Comments:","Terms":"Terms",", ":", ","Unclaimed":"Unclaimed","About the Business":"About the Business","Response Time":"Response Time","If you’re here to leave a review based on a first-hand experience with the business, please check back at a later date.":"If you’re here to leave a review based on a first-hand experience with the business, please check back at a later date.","Friends":"Friends","Sorry, but we didn't understand the location you entered.":"Sorry, but we didn't understand the location you entered.","Yelp Project Cost Guides":"Yelp Project Cost Guides","Ask a Question":"Ask a Question","%{businessName} also recommends":"%{businessName} also recommends","Established in %{yearEstablished}.":"Established in %{yearEstablished}.","<p&gt;This business appears to be affiliated with a group of businesses that have engaged in efforts to manipulate their reputation on Yelp. As a result, we’ve decided not to recommend any of the reviews for this business.</p&gt;<p&gt;Some of the customers of this group have reported being pressured to write positive reviews in exchange for discounts, to re-post positive reviews for affiliated businesses in other states, and to accept refunds in exchange for removing critical reviews.</p&gt;<p&gt;Alerts like these are part of Yelp’s Consumer Protection Initiative, which is designed to empower and protect consumers.</p&gt;":"<p&gt;This business appears to be affiliated with a group of businesses that have engaged in efforts to manipulate their reputation on Yelp. As a result, we’ve decided not to recommend any of the reviews for this business.</p&gt;<p&gt;Some of the customers of this group have reported being pressured to write positive reviews in exchange for discounts, to re-post positive reviews for affiliated businesses in other states, and to accept refunds in exchange for removing critical reviews.</p&gt;<p&gt;Alerts like these are part of Yelp’s Consumer Protection Initiative, which is designed to empower and protect consumers.</p&gt;","View reservation history":"View reservation history","Log Out":"Log Out","Tell us what you expected to find:":"Tell us what you expected to find:","See license information":"See license information","Find Delivery":"Find Delivery","Thank you":"Thank you","More Collections":"More Collections","%{smart_count} More Attribute||||%{smart_count} More Attributes":"%{smart_count} More Attribute||||%{smart_count} More Attributes","%{description} <span class=\"%{queryLocationClass}\"&gt;in %{displayLocation}</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;in %{displayLocation}</span&gt;","Select your rating":"Select your rating","Why do you want to report this content?":"Why do you want to report this content?","Photo of %{user}":"Photo of %{user}","Once you’ve claimed, you can:":"Once you’ve claimed, you can:","Previous review":"Previous review","Duplicate Business":"Duplicate Business","Best of Yelp":"Best of Yelp","Reviews Have Disappeared":"Reviews Have Disappeared","<p&gt;This business has been in the news in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business has been in the news in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Add info":"Add info","Enter delivery address":"Enter delivery address","This business has been claimed by the owner or a representative. ":"This business has been claimed by the owner or a representative. ","RSS":"RSS","Rating Details":"Rating Details","Ultra High-End":"Ultra High-End","We couldn't find you quickly enough! Try again later, or search near a city, place, or address instead.":"We couldn't find you quickly enough! Try again later, or search near a city, place, or address instead.","License type":"License type","Photos and Videos":"Photos and Videos","Your Email":"Your Email","<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Jan":"Jan","Please provide specific details below:":"Please provide specific details below:","Redo Search In Map":"Redo Search In Map","Developers":"Developers","Embedding Code":"Embedding Code","Oct":"Oct","Content Guidelines":"Content Guidelines","Search pages feature Yelp Ads, and your ad can appear in either slot.":"Search pages feature Yelp Ads, and your ad can appear in either slot.","Log In":"Log In","Show more filters":"Show more filters","We accept locations in the following forms:":"We accept locations in the following forms:","Oops, something went wrong. Please try again later.":"Oops, something went wrong. Please try again later.","Thanks!":"Thanks!","Check your place in line from the text we've sent to %{phoneNumber}":"Check your place in line from the text we've sent to %{phoneNumber}","Click to claim the %{cashBack}% cash back offer at %{businessName}.":"Click to claim the %{cashBack}% cash back offer at %{businessName}.","Business hours may be different today.":"Business hours may be different today.","Free Sign-up":"Free Sign-up","Offer valid: %{schedule}.":"Offer valid: %{schedule}.","<span class=\"%{nonEmClass}\"&gt;Businesses within this business</span&gt;":"<span class=\"%{nonEmClass}\"&gt;Businesses within this business</span&gt;","Results aren't relevant":"Results aren't relevant","Find Friends":"Find Friends","eeep! something went wrong.":"eeep! something went wrong.","Special hours today:":"Special hours today:","Just a Note":"Just a Note","response rate":"response rate","Yelp Blog":"Yelp Blog","Project failed to load.":"Project failed to load.","Name of Business:":"Name of Business:","About This Provider":"About This Provider","License number is required.":"License number is required.","Business website":"Business website","Uncheck All":"Uncheck All","Please tell us.":"Please tell us.","Next":"Next","NEW":"NEW","Save to Collection":"Save to Collection","Share on Facebook":"Share on Facebook","Save":"Save","Yelp WiFi":"Yelp WiFi","<a %{attrs}&gt;Add more friends</a&gt; to see them here!":"<a %{attrs}&gt;Add more friends</a&gt; to see them here!","Sorry, that's an invalid search term.":"Sorry, that's an invalid search term.","I am a customer":"I am a customer","Careers":"Careers","Unusual Activity Alert":"Unusual Activity Alert","Sorry, but we don’t take sides in factual disputes. If a review appears to reflect a user’s personal experience and opinions, it is our policy to let the user stand behind their review.":"Sorry, but we don’t take sides in factual disputes. If a review appears to reflect a user’s personal experience and opinions, it is our policy to let the user stand behind their review.","Please enter both Business Name and Business Location.":"Please enter both Business Name and Business Location.","Cool":"Cool","Cash Back":"Cash Back","Webinars":"Webinars","Great Photos":"Great Photos","Are you sure you want to remove this license? This may hide license details from your Yelp page.":"Are you sure you want to remove this license? This may hide license details from your Yelp page.","Terms of Service":"Terms of Service","Claimed":"Claimed","<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minutes":"<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting delivery orders in <b&gt;%{smart_count}</b&gt; minutes","This business may have tried to abuse the legal system in an effort to stifle free speech, for example through legal threats or contractual gag clauses. As a reminder, reviewers who share their experiences have a First Amendment right to express their opinions on Yelp.":"This business may have tried to abuse the legal system in an effort to stifle free speech, for example through legal threats or contractual gag clauses. As a reminder, reviewers who share their experiences have a First Amendment right to express their opinions on Yelp.","Licensee":"Licensee","<ol&gt;\n            <li&gt;Click <b&gt;Safari</b&gt; in the Menu Bar at the top of the screen, then <b&gt;Preferences</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Websites</b&gt; tab, then click <b&gt;Location</b&gt; on the left-hand panel.</li&gt;\n            <li&gt;Next to <b&gt;yelp.com</b&gt; in the right-hand panel, change the dropdown to <b&gt;Ask</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;MacOS may now prompt you to enable Location Services. If it does, follow its instructions to enable Location Services for Safari.</li&gt;\n            <li&gt;Close the Settings dialog and refresh the page. Try using Current Location search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.apple.com/en-us/HT204690\" target=\"_blank\"&gt;Safari's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;Click <b&gt;Safari</b&gt; in the Menu Bar at the top of the screen, then <b&gt;Preferences</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Websites</b&gt; tab, then click <b&gt;Location</b&gt; on the left-hand panel.</li&gt;\n            <li&gt;Next to <b&gt;yelp.com</b&gt; in the right-hand panel, change the dropdown to <b&gt;Ask</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;MacOS may now prompt you to enable Location Services. If it does, follow its instructions to enable Location Services for Safari.</li&gt;\n            <li&gt;Close the Settings dialog and refresh the page. Try using Current Location search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.apple.com/en-us/HT204690\" target=\"_blank\"&gt;Safari's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Find a Table":"Find a Table","Start Order":"Start Order","Apr":"Apr","Navigate":"Navigate","Edit review":"Edit review","Monitor how many people view your business page":"Monitor how many people view your business page","Add a license||||Add another license":"Add a license||||Add another license","Swipe your linked card at the businesses.":"Swipe your linked card at the businesses.","Special hours today.":"Special hours today.","Get %{cashBack}% Cash Back":"Get %{cashBack}% Cash Back","Current Location":"Current Location","Buy Now":"Buy Now","Live wait time:":"Live wait time:","Plumber, locksmith, electrician":"Plumber, locksmith, electrician","Account Settings":"Account Settings","These Projects are paid for and provided by the business. Yelp has not verified these Projects.":"These Projects are paid for and provided by the business. Yelp has not verified these Projects.","*":"*","Previous Project":"Previous Project","If the business you're looking for isn't here, add it!":"If the business you're looking for isn't here, add it!","Yelper names or email addresses":"Yelper names or email addresses","Check the spelling or try alternate spellings":"Check the spelling or try alternate spellings","No one is on the list. When it gets busy, join the waitlist here or on the Yelp app.":"No one is on the list. When it gets busy, join the waitlist here or on the Yelp app.","I don’t work here":"I don’t work here","Like Your Profile":"Like Your Profile","Highlights from the Business":"Highlights from the Business","Report Review":"Report Review","Meet the %{businessOwnerRole}":"Meet the %{businessOwnerRole}","Okay":"Okay","See %{smart_count} question for %{businessName}||||See all %{smart_count} questions for %{businessName}":"See %{smart_count} question for %{businessName}||||See all %{smart_count} questions for %{businessName}","Message from Yelp. Only you can see this.":"Message from Yelp. Only you can see this.","Claim Offer":"Claim Offer","Add this review to your website by copying the code below.":"Add this review to your website by copying the code below.","(no rating)":"(no rating)","Press":"Press","Need help?":"Need help?","Upcoming Reservations":"Upcoming Reservations","Send to your Phone":"Send to your Phone","Search":"Search","Claim This Free Business Page":"Claim This Free Business Page","Helpful":"Helpful","Language":"Language","You’re Funny":"You’re Funny","Or, you can&amp;nbsp;<a href=%{url}&gt;add a business here</a&gt;.":"Or, you can&amp;nbsp;<a href=%{url}&gt;add a business here</a&gt;.","Try a more general search. e.g. \"pizza\" instead of \"pepperoni\"":"Try a more general search. e.g. \"pizza\" instead of \"pepperoni\"","History":"History","Closed now":"Closed now","Sep":"Sep","Got search feedback? ":"Got search feedback? ","Share photo":"Share photo","in %{smart_count} review||||in %{smart_count} reviews":"in %{smart_count} review||||in %{smart_count} reviews","Cancel":"Cancel","Business Support":"Business Support","Show Less":"Show Less","Join the Waitlist":"Join the Waitlist","See details":"See details","Other":"Other","Your Name":"Your Name","%{smart_count} review that is not currently recommended||||%{smart_count} reviews that are not currently recommended":"%{smart_count} review that is not currently recommended||||%{smart_count} reviews that are not currently recommended","Home Address":"Home Address","Success":"Success","Sign up and claim your %{cashBack}% Cash Back offer.":"Sign up and claim your %{cashBack}% Cash Back offer.","Yelp confirmed at least one person associated with this business has the above license. Yelp cannot verify if a license covers your specific needs or that everyone at this business has a required license. Businesses pay Yelp for license verification services.":"Yelp confirmed at least one person associated with this business has the above license. Yelp cannot verify if a license covers your specific needs or that everyone at this business has a required license. Businesses pay Yelp for license verification services.","Phone number":"Phone number","Monthly Trend":"Monthly Trend","Photos":"Photos","<strong&gt;%{businessOrOwnerName} says,</strong&gt; &amp;ldquo;%{recommendationText}&amp;rdquo;":"<strong&gt;%{businessOrOwnerName} says,</strong&gt; &amp;ldquo;%{recommendationText}&amp;rdquo;","Something broke and we're not sure what. Try again later, or search near a city, place, or address instead.":"Something broke and we're not sure what. Try again later, or search near a city, place, or address instead.","Add a note":"Add a note","Is this your business?":"Is this your business?","Pending Verification":"Pending Verification","Oops! Something went wrong. Please try again.":"Oops! Something went wrong. Please try again.","Claim this business. It’s free.":"Claim this business. It’s free.","Get Verified License":"Get Verified License","Are you sure you would like to remove this unfinished review?":"Are you sure you would like to remove this unfinished review?","Savings":"Savings","Share review":"Share review","Expiration":"Expiration","This business is eligible to be claimed by a local representative in addition to corporate.":"This business is eligible to be claimed by a local representative in addition to corporate.","Messaging":"Messaging","<p&gt;This business has <a href=\"%{evidenceUrl}\"&gt;been in the news</a&gt; in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business has <a href=\"%{evidenceUrl}\"&gt;been in the news</a&gt; in connection with a recent tragedy, which means people may be coming to this page to share their thoughts and concerns.</p&gt;<p&gt;We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events, even though we might agree with some of the viewpoints people seek to express. The best place to share your thoughts in response to this event is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;.</p&gt;<p&gt;We apply this same policy to all businesses featured in the media regardless of the business and regardless of the issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Enter your home address":"Enter your home address","%{coverCount} person||||%{coverCount} people":"%{coverCount} person||||%{coverCount} people","Claim your business to immediately update business information, track page views, and more!":"Claim your business to immediately update business information, track page views, and more!","<p&gt;This business recently participated in a film that makes false statements about Yelp and our reviewers.</p&gt;<p&gt; While some say this is the era of fake news, we still think facts matter. Businesses can certainly say what they wish, but you should learn the truth about this business’ interaction with Yelp.</p&gt;":"<p&gt;This business recently participated in a film that makes false statements about Yelp and our reviewers.</p&gt;<p&gt; While some say this is the era of fake news, we still think facts matter. Businesses can certainly say what they wish, but you should learn the truth about this business’ interaction with Yelp.</p&gt;","This business is closed on %{date}, %{holidayName}.":"This business is closed on %{date}, %{holidayName}.","Eek! Methinks not.":"Eek! Methinks not.","Cute Pic":"Cute Pic","<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;":"<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news rather than actual consumer experiences with the business.</p&gt;<p&gt;The best place to share your thoughts is on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt;. We’ve temporarily disabled the posting of content to this page as we work to verify the content you see here reflects actual consumer experiences rather than the recent events (even if that means disabling the ability for users to express points of view we might agree with).</p&gt;<p&gt;Please note that we apply this same policy regardless of the business and regardless of the topic at issue. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.</p&gt;","Consumer Alert: Get the Facts":"Consumer Alert: Get the Facts","Why do you want to report this photo?":"Why do you want to report this photo?","%{cashBack}% Cash Back Offer Details":"%{cashBack}% Cash Back Offer Details","Check out %{reviewUser}’s review of %{businessName} on @yelp":"Check out %{reviewUser}’s review of %{businessName} on @yelp","Error":"Error","Read more":"Read more","<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended reviews":"<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; recommended reviews","A link to the business owner’s business":"A link to the business owner’s business","More information about the action that led to this Consumer Alert is available <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.":"More information about the action that led to this Consumer Alert is available <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.","This business is a Yelp advertiser.":"This business is a Yelp advertiser.","Public":"Public","Meh. I've experienced better.":"Meh. I've experienced better.","Clear results":"Clear results","View Full Menu":"View Full Menu","Dec":"Dec","Oops! <span class=\"%{nonEmClass}\"&gt;Please try again later.</span&gt;":"Oops! <span class=\"%{nonEmClass}\"&gt;Please try again later.</span&gt;","Please provide any additional information about your license such as DBA, registered address etc.":"Please provide any additional information about your license such as DBA, registered address etc.","Are you a human? Please complete the bot challenge below.":"Are you a human? Please complete the bot challenge below.","Browsing Businesses in %{displayLocation}":"Browsing Businesses in %{displayLocation}","<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minutes":"<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minute||||<b&gt;Get your order in.</b&gt; This restaurant will stop accepting pickup orders in <b&gt;%{smart_count}</b&gt; minutes","Pay at selected restaurants":"Pay at selected restaurants","Consumer Alert: Questionable Legal Threats":"Consumer Alert: Questionable Legal Threats","Notifications":"Notifications","Report Video":"Report Video","Oldest First":"Oldest First","Services":"Services","Try a different location.":"Try a different location.","Also, it's possible we don't have a listing for %{findLocation}. In that case, you should try adding a zip, or try a larger nearby city.":"Also, it's possible we don't have a listing for %{findLocation}. In that case, you should try adding a zip, or try a larger nearby city.","Serving %{serviceArea} and the Surrounding Area":"Serving %{serviceArea} and the Surrounding Area","Notes":"Notes","Get access to Yelp’s free tools":"Get access to Yelp’s free tools","View Service Area":"View Service Area","%{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;","Reviews":"Reviews","Yelp Blog for Business Owners":"Yelp Blog for Business Owners","Specialties":"Specialties","Photos for %{business}":"Photos for %{business}","View more photos":"View more photos","Watch Video":"Watch Video","Save Changes":"Save Changes","Select Country Code":"Select Country Code","Order Delivery":"Order Delivery","%{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;","Rating":"Rating","<p&gt;\n            Oops! We don't recognize the web browser you're currently using. Try checking the browser's help menu, or searching the Web for instructions to turn on HTML5 Geolocation for your browser.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<p&gt;\n            Oops! We don't recognize the web browser you're currently using. Try checking the browser's help menu, or searching the Web for instructions to turn on HTML5 Geolocation for your browser.\n            You can also search near a city, place, or address instead.\n        </p&gt;","This business is closed today.":"This business is closed today.","Business Success Stories":"Business Success Stories","Send message":"Send message","Did you mean?":"Did you mean?","Email addresses":"Email addresses","Search Results Feedback":"Search Results Feedback","Please refer to our <a %{contentGuidelinesLinkAttrs}&gt;Content Guidelines</a&gt; and <a %{termsOfServiceLinkAttrs}&gt;Terms of Service</a&gt; and let us know why you think the content you’ve reported may violate these guidelines.":"Please refer to our <a %{contentGuidelinesLinkAttrs}&gt;Content Guidelines</a&gt; and <a %{termsOfServiceLinkAttrs}&gt;Terms of Service</a&gt; and let us know why you think the content you’ve reported may violate these guidelines.","This business has not yet been claimed by the owner or a representative.":"This business has not yet been claimed by the owner or a representative.","<b&gt;Your trust is our top concern,</b&gt; so businesses can’t pay to alter or remove their reviews. <a href=\"%{reviewSupportUrl}\"&gt;Learn more about reviews.</a&gt;":"<b&gt;Your trust is our top concern,</b&gt; so businesses can’t pay to alter or remove their reviews. <a href=\"%{reviewSupportUrl}\"&gt;Learn more about reviews.</a&gt;","Mo' Map":"Mo' Map","Business listing provided by":"Business listing provided by","Finish My Review":"Finish My Review","%{smart_count} Update||||%{smart_count} Updates":"%{smart_count} Update||||%{smart_count} Updates","Send":"Send","Keep Review":"Keep Review","Collections including %{businessName}":"Collections including %{businessName}","More Topics":"More Topics","<strong&gt;%{distance}</strong&gt; away from this business":"<strong&gt;%{distance}</strong&gt; away from this business","Visit Yelp for Business Owners":"Visit Yelp for Business Owners","Special hours":"Special hours","Sponsored Results":"Sponsored Results","This business has been in the news in connection with a recent tragedy, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.":"This business has been in the news in connection with a recent tragedy, which often means people come to this page to post their views on the news rather than a first-hand consumer experience. As a result, we’ve temporarily disabled the ability to post content about this business. Read more on <a href=\"%{yelpSupportLink}\"&gt;Yelp Support</a&gt;.","Sorry, the specified search area is too large.":"Sorry, the specified search area is too large.","We caught someone offering up cash, discounts, gift certificates or other incentives in exchange for reviews about this business. We wanted you to know because buying reviews not only hurts consumers, but also honest businesses who play by the rules.":"We caught someone offering up cash, discounts, gift certificates or other incentives in exchange for reviews about this business. We wanted you to know because buying reviews not only hurts consumers, but also honest businesses who play by the rules.","Hours":"Hours","Pricey":"Pricey","Show More":"Show More","Write a Review":"Write a Review","Oops, there was a problem flagging this content. Please try again later.":"Oops, there was a problem flagging this content. Please try again later.","Info":"Info","Verified by Yelp on":"Verified by Yelp on","Hide less relevant categories.":"Hide less relevant categories.","Inexpensive":"Inexpensive","<a href=\"%{firstURL}\"&gt;%{firstName}</a&gt; and <a href=\"%{secondURL}\"&gt;%{secondName}</a&gt; at this location.":"<a href=\"%{firstURL}\"&gt;%{firstName}</a&gt; and <a href=\"%{secondURL}\"&gt;%{secondName}</a&gt; at this location.","Open now":"Open now","Join Yelp Cash Back":"Join Yelp Cash Back","%{smart_count} review||||%{smart_count} reviews":"%{smart_count} review||||%{smart_count} reviews","There’s always more than one side to a story, and business owners can address misunderstandings via their Business Account by posting a public comment or sending a direct message to the reviewer.":"There’s always more than one side to a story, and business owners can address misunderstandings via their Business Account by posting a public comment or sending a direct message to the reviewer.","Privacy Policy":"Privacy Policy","Powered by":"Powered by","Make a Reservation":"Make a Reservation","By appointment only":"By appointment only","Yes":"Yes","tacos, cheap dinner, Max’s":"tacos, cheap dinner, Max’s","No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;":"No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;%{displayLocation}</span&gt;","(optional)":"(optional)","Page: %{page}":"Page: %{page}","Yelp Nowait":"Yelp Nowait","Show more":"Show more","Ad Choices":"Ad Choices","Closed":"Closed","Moderate":"Moderate","Open today.":"Open today.","Events":"Events","Compliment":"Compliment","Appointment":"Appointment","Jun":"Jun","A business owner paid for this ad. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.":"A business owner paid for this ad. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.","Jul":"Jul","Discount":"Discount","Collections Including %{businessName}":"Collections Including %{businessName}","Got It":"Got It","Yelp for Business Owners":"Yelp for Business Owners","Non-Public":"Non-Public","Upcoming Special Hours":"Upcoming Special Hours","No matching friends found":"No matching friends found","Great Lists":"Great Lists","No answers yet.":"No answers yet.","Add a Business":"Add a Business","Sorry, we couldn’t recognize your address.":"Sorry, we couldn’t recognize your address.","Select a Previous Order (Optional)":"Select a Previous Order (Optional)","Advertise on Yelp":"Advertise on Yelp","Work here? ":"Work here? ","Copy Code":"Copy Code","Useful":"Useful","Please enter an address to start your order.":"Please enter an address to start your order.","View More":"View More","Started on %{date}":"Started on %{date}","<b&gt;Date of experience:</b&gt; %{date_of_experience}":"<b&gt;Date of experience:</b&gt; %{date_of_experience}","Show all results.":"Show all results.","Yelping since %{year_joined} with %{num_reviews} reviews":"Yelping since %{year_joined} with %{num_reviews} reviews","Get %{cashBack}% of your bill deposited back to your card.":"Get %{cashBack}% of your bill deposited back to your card.","Your Email Address:":"Your Email Address:","Based on data from <a %{linkAttrs}&gt;%{providerName}</a&gt;":"Based on data from <a %{linkAttrs}&gt;%{providerName}</a&gt;","License #":"License #","Suggestions for improving the results:":"Suggestions for improving the results:","Link your credit or debit cards":"Link your credit or debit cards","From the business":"From the business","Finding your location...":"Finding your location...","Yelp confirmed a business or employee license.":"Yelp confirmed a business or employee license.","You're on the waitlist":"You're on the waitlist","Thank You":"Thank You","Confirmation":"Confirmation","Learn more":"Learn more","<b&gt;%{businessOwnerDisplayName}</b&gt; sent you thanks for this review":"<b&gt;%{businessOwnerDisplayName}</b&gt; sent you thanks for this review","Claim your Business Page":"Claim your Business Page","You're getting %{cashBack}% Cash Back here!":"You're getting %{cashBack}% Cash Back here!","See More Projects":"See More Projects","address, neighborhood, city, state or zip":"address, neighborhood, city, state or zip","See %{smart_count} More Service||||See %{smart_count} More Services":"See %{smart_count} More Service||||See %{smart_count} More Services","%{currentNumberLicense} of %{totalNumberLicense}":"%{currentNumberLicense} of %{totalNumberLicense}","Reg. Price":"Reg. Price","About":"About","Consumer Alert":"Consumer Alert","Check out the evidence <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.":"Check out the evidence <a href=\"%{evidenceUrl}\" target=\"_blank\"&gt;here</a&gt;.","Read Less":"Read Less","For more information visit Yelp for Business Owners.":"For more information visit Yelp for Business Owners.","Hot Stuff":"Hot Stuff","Got it":"Got it","We couldn't find an accurate position. If you're using a laptop or tablet, try moving it somewhere else and give it another go. Or, search near a city, place, or address instead.":"We couldn't find an accurate position. If you're using a laptop or tablet, try moving it somewhere else and give it another go. Or, search near a city, place, or address instead.","Good Writer":"Good Writer","%{smart_count} Review||||%{smart_count} Reviews":"%{smart_count} Review||||%{smart_count} Reviews","The restaurant is not taking waitlist parties right now":"The restaurant is not taking waitlist parties right now","Sort by":"Sort by","<strong&gt;Go mobile </strong&gt; with the <a href=%{mobileAppLinkHref}&gt; %{mobileAppLinkText} </a&gt; for iOS and Android":"<strong&gt;Go mobile </strong&gt; with the <a href=%{mobileAppLinkHref}&gt; %{mobileAppLinkText} </a&gt; for iOS and Android","<b&gt;%{smart_count}</b&gt; review||||<b&gt;%{smart_count}</b&gt; reviews":"<b&gt;%{smart_count}</b&gt; review||||<b&gt;%{smart_count}</b&gt; reviews","Saved":"Saved","Sign up to claim your %{cashBack}% Cash Back offer.":"Sign up to claim your %{cashBack}% Cash Back offer.","Edit business info":"Edit business info","Offer Claimed":"Offer Claimed","Collections":"Collections","Report Comment":"Report Comment","See all photos from %{userName} for %{businessName}":"See all photos from %{userName} for %{businessName}","A number of positive reviews for this business originated from the same IP address. Our <a href=\"%{recommendedReviewFAQUrl}\"&gt; automated recommendation software</a&gt; has taken this into account in choosing which reviews to display, but we wanted to call this to your attention because someone may be trying to artificially inflate the rating for this business.":"A number of positive reviews for this business originated from the same IP address. Our <a href=\"%{recommendedReviewFAQUrl}\"&gt; automated recommendation software</a&gt; has taken this into account in choosing which reviews to display, but we wanted to call this to your attention because someone may be trying to artificially inflate the rating for this business.","<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;":"<p&gt;This business recently made waves in the news, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;","Next Project":"Next Project","Add business hours":"Add business hours","Response time:":"Response time:","Not Helpful":"Not Helpful","Get Cash Back":"Get Cash Back","We've found multiple locations matching your search.":"We've found multiple locations matching your search.","Yelp Reservations":"Yelp Reservations","Delivery Address":"Delivery Address","Collection Name":"Collection Name","Share business":"Share business","Comment from %{businessOwnerDisplayName} of %{businessName}":"Comment from %{businessOwnerDisplayName} of %{businessName}","See More <b&gt;%{categoryName} in %{cityName}</b&gt;":"See More <b&gt;%{categoryName} in %{cityName}</b&gt;","Swipe your linked card at the restaurant.":"Swipe your linked card at the restaurant.","Messages":"Messages","Updated review":"Updated review","<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; reviews":"<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; review||||<span itemprop=\"reviewCount\"&gt;%{smart_count}</span&gt; reviews","Services Offered":"Services Offered","Business Listings provided by Yellow Pages®":"Business Listings provided by Yellow Pages®","You’re Cool":"You’re Cool","Recommended Reviews":"Recommended Reviews","Try a larger search area":"Try a larger search area","Talk":"Talk","Issued by":"Issued by","License Number":"License Number","Search within reviews":"Search within reviews","OK":"OK","Warning: Users Report Deceptive Behavior":"Warning: Users Report Deceptive Behavior","Stop following %{displayName}":"Stop following %{displayName}","Share":"Share","Posting Temporarily Blocked":"Posting Temporarily Blocked","Warning":"Warning","Help us improve.":"Help us improve.","Find":"Find","<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;":"<p&gt;This business recently made <a href=\"%{evidenceUrl}\"&gt;waves in the news</a&gt;, which often means people come to this page to post their views on the news.</p&gt;<p&gt;While we don’t take a stand one way or the other when it comes to this news event, we work to verify that the content you see here reflects personal consumer experiences with the business rather than the news itself. As a result, we’ve temporarily disabled the posting of content to this page.</p&gt;<p&gt;You should feel free to post your thoughts about the recent media coverage for this business on <a href=\"%{yelpTalkLink}\"&gt;Yelp Talk</a&gt; at any time.</p&gt;","Browsing Businesses near %{displayLocation}":"Browsing Businesses near %{displayLocation}","Location of Business: (City, State)":"Location of Business: (City, State)","No answers yet. You can be the first!":"No answers yet. You can be the first!","%{smart_count} review mentioning “%{query}”||||%{smart_count} reviews mentioning “%{query}”":"%{smart_count} review mentioning “%{query}”||||%{smart_count} reviews mentioning “%{query}”","All results shown.":"All results shown.","Are you a customer or the owner/manager of the business you'd like to add?":"Are you a customer or the owner/manager of the business you'd like to add?","See All Photos":"See All Photos","Add photos":"Add photos","<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. Also, try <a href=\"https://support.apple.com/en-us/HT204690\"&gt;enabling Location Services for your browser</a&gt;. If your browser still doesn't ask you, try these steps:</p&gt;":"<p&gt;First, try refreshing the page and clicking Current Location again. Make sure you click <b&gt;Allow</b&gt; or <b&gt;Grant Permissions</b&gt; if your browser asks for your location. Also, try <a href=\"https://support.apple.com/en-us/HT204690\"&gt;enabling Location Services for your browser</a&gt;. If your browser still doesn't ask you, try these steps:</p&gt;","Oops, something went wrong.":"Oops, something went wrong.","We need your address to show you restaurants that deliver.":"We need your address to show you restaurants that deliver.","Preview":"Preview","No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;":"No Results for %{description} <span class=\"%{queryLocationClass}\"&gt;near <a href=\"%{locationUrl}\"&gt;%{displayLocation}</a&gt;</span&gt;","Related Talk Topics":"Related Talk Topics","Is this your business? Claim it for free!":"Is this your business? Claim it for free!","Buy Gift Certificate":"Buy Gift Certificate","Pagination navigation":"Pagination navigation","Toggle Menu":"Toggle Menu","Located in <a href=\"%{businessUrl}\"&gt;%{name}</a&gt;":"Located in <a href=\"%{businessUrl}\"&gt;%{name}</a&gt;","Sponsored":"Sponsored","Remove license":"Remove license","%{description} <span class=\"%{queryLocationClass}\"&gt;near %{displayLocation}</span&gt;":"%{description} <span class=\"%{queryLocationClass}\"&gt;near %{displayLocation}</span&gt;","Lowest Rated":"Lowest Rated","Leave Waitlist":"Leave Waitlist","Enter your address to find businesses that deliver to you":"Enter your address to find businesses that deliver to you","This is my business":"This is my business","Claim this business":"Claim this business","Near":"Near","Oops, something went wrong. Please try again.":"Oops, something went wrong. Please try again.","Embed review":"Embed review","A public Collection can be openly featured on Yelp and alerts followers when you make updates. A non-public Collection can still be visible to others if you share a link to it.":"A public Collection can be openly featured on Yelp and alerts followers when you make updates. A non-public Collection can still be visible to others if you share a link to it.","Closed today.":"Closed today.","We calculate the overall star rating using only reviews that our automated software currently recommends. <a %{link_attrs}&gt;Learn more</a&gt;.":"We calculate the overall star rating using only reviews that our automated software currently recommends. <a %{link_attrs}&gt;Learn more</a&gt;.","OR":"OR","If this is your business, claim it and use Yelp’s free tools for business owners: respond to reviews, get detailed analytics and convert visitors into customers.":"If this is your business, claim it and use Yelp’s free tools for business owners: respond to reviews, get detailed analytics and convert visitors into customers.","Message":"Message","Previous":"Previous","%{reviewCount} review||||%{reviewCount} reviews":"%{reviewCount} review||||%{reviewCount} reviews","Issuing Agency":"Issuing Agency","Pay at selected businesses":"Pay at selected businesses","<a href=\"%{childrenBusinessURL}\"&gt;See businesses</a&gt; at this location":"<a href=\"%{childrenBusinessURL}\"&gt;See businesses</a&gt; at this location","Feb":"Feb","Ads":"Ads","No":"No","Known For":"Known For","DexYP logo":"DexYP logo","Oops, looks like something’s wrong. Try again!":"Oops, looks like something’s wrong. Try again!","<ol&gt;\n            <li&gt;At the top of your Opera window, near the web address, you should see a <b&gt;gray location pin</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, click <b&gt;Clear This Setting</b&gt;</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://help.opera.com/en/presto/control-pages/#geolocation\" target=\"_blank\"&gt;Opera's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top of your Opera window, near the web address, you should see a <b&gt;gray location pin</b&gt;. Click it.</li&gt;\n            <li&gt;In the window that pops up, click <b&gt;Clear This Setting</b&gt;</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://help.opera.com/en/presto/control-pages/#geolocation\" target=\"_blank\"&gt;Opera's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Share on Twitter":"Share on Twitter","<p&gt;Did you know that local officials inspect food service facilities to improve food safety?</p&gt;<p&gt;Following a recent inspection, this facility received a food safety rating that is in the bottom 5% locally, and is categorized by inspectors as “poor”.</p&gt;<p&gt;Being in the consumer protection business, we care a lot about your safety and will display this alert for six months or until we receive a significantly improved food safety rating for this business.</p&gt;":"<p&gt;Did you know that local officials inspect food service facilities to improve food safety?</p&gt;<p&gt;Following a recent inspection, this facility received a food safety rating that is in the bottom 5% locally, and is categorized by inspectors as “poor”.</p&gt;<p&gt;Being in the consumer protection business, we care a lot about your safety and will display this alert for six months or until we receive a significantly improved food safety rating for this business.</p&gt;","Today is a holiday!":"Today is a holiday!","Business owners paid for these ads. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.":"Business owners paid for these ads. For more information visit <a href=\"%{bizSiteUrl}\"&gt;Yelp for Business Owners</a&gt;.","This business is open as usual on %{date}, %{holidayName}.":"This business is open as usual on %{date}, %{holidayName}.","Verified by Business":"Verified by Business","Not here? Tell us what we're missing.":"Not here? Tell us what we're missing.","Businesses are too far away":"Businesses are too far away","Help Improve Yelp":"Help Improve Yelp","First to review":"First to review","Consumer Alert: Poor Food Safety Score!":"Consumer Alert: Poor Food Safety Score!","%{smart_count} person found this helpful||||%{smart_count} people found this helpful":"%{smart_count} person found this helpful||||%{smart_count} people found this helpful","<ol&gt;\n            <li&gt;At the top of your Chrome window, near the web address, click <b&gt;the gray lock icon</b&gt;.</li&gt;\n            <li&gt;In the window that pops up, make sure <b&gt;Location</b&gt; is set to <b&gt;Ask (default)</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.google.com/chrome/answer/142065?co=GENIE.Platform%3DDesktop&amp;hl=en&amp;oco=0\" target=\"_blank\"&gt;Google's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;At the top of your Chrome window, near the web address, click <b&gt;the gray lock icon</b&gt;.</li&gt;\n            <li&gt;In the window that pops up, make sure <b&gt;Location</b&gt; is set to <b&gt;Ask (default)</b&gt; or <b&gt;Allow</b&gt;.</li&gt;\n            <li&gt;You're good to go! Reload this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            If you're still having trouble, check out <a href=\"https://support.google.com/chrome/answer/142065?co=GENIE.Platform%3DDesktop&amp;hl=en&amp;oco=0\" target=\"_blank\"&gt;Google's support page</a&gt;.\n            You can also search near a city, place, or address instead.\n        </p&gt;","Redo search when map is moved":"Redo search when map is moved","Did you mean: ":"Did you mean: ","Save Money":"Save Money","Yelp Mobile":"Yelp Mobile","Project Duration":"Project Duration","%{smart_count} other review that is not currently recommended||||%{smart_count} other reviews that are not currently recommended":"%{smart_count} other review that is not currently recommended||||%{smart_count} other reviews that are not currently recommended","Overall Rating":"Overall Rating","<b&gt;%{smart_count}</b&gt; photo||||<b&gt;%{smart_count}</b&gt; photos":"<b&gt;%{smart_count}</b&gt; photo||||<b&gt;%{smart_count}</b&gt; photos","Less Map":"Less Map","Please arrive by %{arriveBy}":"Please arrive by %{arriveBy}","<ol&gt;\n            <li&gt;Click the <b&gt;gear</b&gt; in the upper-right hand corner of the window, then <b&gt;Internet options</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Privacy</b&gt; tab in the new window that just appeared.</li&gt;\n            <li&gt;Uncheck the box labeled <b&gt;Never allow websites to request your physical location</b&gt; if it's already checked.</li&gt;\n            <li&gt;Click the button labeled <b&gt;Clear Sites</b&gt;.</li&gt;\n            <li&gt;You're good to go! Click <b&gt;OK</b&gt;, then refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;":"<ol&gt;\n            <li&gt;Click the <b&gt;gear</b&gt; in the upper-right hand corner of the window, then <b&gt;Internet options</b&gt;.</li&gt;\n            <li&gt;Click the <b&gt;Privacy</b&gt; tab in the new window that just appeared.</li&gt;\n            <li&gt;Uncheck the box labeled <b&gt;Never allow websites to request your physical location</b&gt; if it's already checked.</li&gt;\n            <li&gt;Click the button labeled <b&gt;Clear Sites</b&gt;.</li&gt;\n            <li&gt;You're good to go! Click <b&gt;OK</b&gt;, then refresh this Yelp page and try your search again.</li&gt;\n        </ol&gt;\n        <p&gt;\n            You can also search near a city, place, or address instead.\n        </p&gt;","Why do you want to remove this review?":"Why do you want to remove this review?","Close":"Close","<b&gt;%{smart_count}</b&gt; friend||||<b&gt;%{smart_count}</b&gt; friends":"<b&gt;%{smart_count}</b&gt; friend||||<b&gt;%{smart_count}</b&gt; friends","We're sorry, the page of results you requested is unavailable.":"We're sorry, the page of results you requested is unavailable.","Yelp Sort":"Yelp Sort","Can't find a business":"Can't find a business","Ask the Community":"Ask the Community","Please select at least one reason.":"Please select at least one reason.","Response rate:":"Response rate:","This business is open %{hours} on %{date}, %{holidayName}.":"This business is open %{hours} on %{date}, %{holidayName}.","Work at %{businessName}? Claim your business":"Work at %{businessName}? Claim your business","Discover":"Discover","Blog":"Blog","To":"To","Collapse hours":"Collapse hours","All Filters":"All Filters","Sign Up":"Sign Up","Why do you want to report this review?":"Why do you want to report this review?","Don’t see your question? Ask away!":"Don’t see your question? Ask away!","See all hours":"See all hours","Remove Review":"Remove Review","Elites":"Elites","Map":"Map","Mar":"Mar","May":"May","Get License Verified":"Get License Verified","Read More":"Read More","<b&gt;Your trust is our top concern,</b&gt; so businesses can't pay to alter or remove their reviews. <a href=\"%{reviewFeedMythsUrl}\"&gt;Learn more.</a&gt;":"<b&gt;Your trust is our top concern,</b&gt; so businesses can't pay to alter or remove their reviews. <a href=\"%{reviewFeedMythsUrl}\"&gt;Learn more.</a&gt;","FAQ":"FAQ","Report Photo":"Report Photo","Report review":"Report review","Your answers help people make good choices when they’re thinking about places to go and things to do.":"Your answers help people make good choices when they’re thinking about places to go and things to do.","Newest First":"Newest First","Location &amp; Hours":"Location &amp; Hours","Highest Rated":"Highest Rated","Edit":"Edit","Remove":"Remove","Popular Dishes":"Popular Dishes","See the customer leads your business page generates":"See the customer leads your business page generates","Follow %{displayName}":"Follow %{displayName}","Remove review":"Remove review","Can&amp;rsquo;t find reviews? Read more about this topic&amp;nbsp;<a href=\"/faq#missing_user_reviews\"&gt;here</a&gt;.":"Can&amp;rsquo;t find reviews? Read more about this topic&amp;nbsp;<a href=\"/faq#missing_user_reviews\"&gt;here</a&gt;.","Add Photos":"Add Photos","Ad":"Ad","%{smart_count} recommended review||||%{smart_count} recommended reviews":"%{smart_count} recommended review||||%{smart_count} recommended reviews","License Verified":"License Verified","Verified License":"Verified License","Private":"Private","under 5 mins":"under 5 mins","No Results? Try unchecking some of the boxes above.":"No Results? Try unchecking some of the boxes above.","About This Agent":"About This Agent","Claim %{cashBack}% Cash Back Here":"Claim %{cashBack}% Cash Back Here"},"props":{"loggedIn":false,"businessId":"4yPqqJDJOQX69gC66YUDkA","twitterUrl":"https://twitter.com/share?text=Check+out+Peter+Luger+on+%40yelp&amp;url=https%3A%2F%2Fwww.yelp.com%2Fbiz%2Fpeter-luger-brooklyn-2%3Futm_campaign%3Dwww_business_share_popup%26utm_medium%3Dsocial%26utm_source%3Dtwitter.com","searchForFriendsCsrfToken":"f872c104f17527af11c81cafac2b5c932af80296d9006e1b35d668d83e915cd6","disabled":false,"findFriendsUrl":"/find_friends/","shareUrl":"https://www.yelp.com/biz/peter-luger-brooklyn-2?utm_campaign=www_business_share_popup&amp;utm_medium=copy_link&amp;utm_source=(direct)","facebookUrl":"http://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fwww.yelp.com%2Fbiz%2Fpeter-luger-brooklyn-2%3Futm_campaign%3Dwww_business_share_popup%26utm_medium%3Dsocial%26utm_source%3Dfacebook.com","submitCsrfToken":"72c0c02be2be54646f6ca5e38e0eb4ae1a8090623a540aeb69fc325b304e30b6","searchForFriendsUrl":"/send_to_friend/friend_type_ahead_v2","submitUrl":"/send_to_friend/business/4yPqqJDJOQX69gC66YUDkA"}}--></script>
        </div>


    <div class="ybtn ybtn--small ybtn--secondary js-show-modal" data-ga-action="click" data-ga-category="collections" data-ga-label="biz page save button" data-modal-url="/collection/user/save_to_collection?item_id=4yPqqJDJOQX69gC66YUDkA">
                <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-save icon--size-14 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_save" />
    </svg>
</span>
                <span class="save-text">
                        Save
                </span>
                <span class="offscreen">, Opens a popup</span>
    </div>

    </span>
    </div>


        </div>

        
    </div>

    <div class="biz-page-subheader">
        <div class="mapbox-container">
                <div class="mapbox" data-lightbox-page-title="Directions - Peter Luger - Williamsburg - South Side - Brooklyn, NY">
            <div class="mapbox-map">
        <a href="/map/peter-luger-brooklyn-2" class="biz-map-directions">
                    
    <img alt="Map" height="135" src="https://maps.googleapis.com/maps/api/staticmap?scale=2&amp;center=40.709945%2C-73.962478&amp;language=None&amp;zoom=15&amp;markers=scale%3A2%7Cicon%3Ahttps%3A%2F%2Fyelp-images.s3.amazonaws.com%2Fassets%2Fmap-markers%2Fannotation_64x86.png%7C40.709945%2C-73.962478&amp;client=gme-yelp&amp;sensor=false&amp;size=286x135&amp;signature=QcRIRDJTn4-OdIPf1zWN3Ks5cGc=" width="286">


        </a>
        <div
            class="lightbox-map hidden"
            data-map-state="{&#34;serviceAreas&#34;: [], &#34;moMapPossible&#34;: true, &#34;scrollwheelZoom&#34;: false, &#34;zoomControlPosition&#34;: &#34;top_right&#34;, &#34;minZoomlevel&#34;: null, &#34;isFullBleed&#34;: false, &#34;maxZoomlevel&#34;: null, &#34;zoom&#34;: 15, &#34;library&#34;: &#34;google&#34;, &#34;fitToGeobox&#34;: false, &#34;hoods&#34;: [], &#34;adPinColor&#34;: null, &#34;markers&#34;: [{&#34;location&#34;: null, &#34;key&#34;: &#34;directions_marker&#34;, &#34;icon&#34;: {&#34;name&#34;: &#34;directions&#34;, &#34;anchorOffset&#34;: [12, 32], &#34;activeOrigin&#34;: [0, 0], &#34;scaledSize&#34;: [24, 32], &#34;regularUri&#34;: &#34;https://s3-media4.fl.yelpcdn.com/assets/srv0/yelp_maps/79f63ebc20db/assets/img/directions@2x.png&#34;, &#34;size&#34;: [24, 32], &#34;activeUri&#34;: &#34;https://s3-media4.fl.yelpcdn.com/assets/srv0/yelp_maps/7249ab345ac8/assets/img/directions_highlighted@2x.png&#34;, &#34;regularOrigin&#34;: [0, 0]}}, {&#34;resourceType&#34;: &#34;business&#34;, &#34;url&#34;: &#34;/biz/peter-luger-brooklyn-2&#34;, &#34;resourceId&#34;: &#34;4yPqqJDJOQX69gC66YUDkA&#34;, &#34;shouldOpenInNewTab&#34;: false, &#34;location&#34;: {&#34;latitude&#34;: 40.709945, &#34;longitude&#34;: -73.962478}, &#34;key&#34;: &#34;starred_business&#34;, &#34;hovercardId&#34;: &#34;FM6GMdzV5FR83nQvEZJu3A&#34;, &#34;icon&#34;: {&#34;name&#34;: &#34;starred&#34;, &#34;anchorOffset&#34;: [12, 32], &#34;activeOrigin&#34;: [0, 0], &#34;scaledSize&#34;: [24, 32], &#34;regularUri&#34;: &#34;https://s3-media4.fl.yelpcdn.com/assets/srv0/yelp_maps/7ffd8e34c576/assets/img/annotation_star@2x.png&#34;, &#34;size&#34;: [24, 32], &#34;activeUri&#34;: &#34;https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_maps/63e0e17c0579/assets/img/annotation_star_highlighted@2x.png&#34;, &#34;regularOrigin&#34;: [0, 0]}}, {&#34;location&#34;: null, &#34;key&#34;: &#34;current_location&#34;, &#34;icon&#34;: {&#34;name&#34;: &#34;current_location&#34;, &#34;anchorOffset&#34;: [23, 23], &#34;activeOrigin&#34;: [0, 0], &#34;scaledSize&#34;: [46, 46], &#34;regularUri&#34;: &#34;https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_maps/56b9eee3f027/assets/img/current_location_dot@2x.png&#34;, &#34;size&#34;: [46, 46], &#34;activeUri&#34;: &#34;https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_maps/56b9eee3f027/assets/img/current_location_dot@2x.png&#34;, &#34;regularOrigin&#34;: [0, 0]}}], &#34;shouldDrawCheckbox&#34;: false, &#34;overlayWidth&#34;: 335, &#34;topBizBounds&#34;: null, &#34;market&#34;: null, &#34;center&#34;: {&#34;latitude&#34;: 40.709945, &#34;longitude&#34;: -73.962478}}"
            data-locale="&lt;yelp.core.locale.api.Locale en_US&gt;"
            data-polyglot-translations="{}"
            data-business-id="4yPqqJDJOQX69gC66YUDkA"
            data-business-address="178 Broadway, Brooklyn, NY 11211"
            data-location-dropper-uri="/locations"
            
        >
                    <div class="map-wrapper">
            <div class="side-box">
                
    <div class="get-directions-box island">
        <div class="get-directions-content">
                <h3>Get directions</h3>
    <div class="contentbox">
        <form action="/map/peter-luger-brooklyn-2" method='GET' class="clearfix yform">
            <div class="section-header u-space-preventcollapse-block">    



        <div class="tab-nav-container">
            <ul class="tab-nav js-tab-nav tab-nav--full js-tab-nav--full" role="tablist">
                    
            <li class="tab-nav_item" role="presentation">
                        





    <a aria-selected="true" class="tab-link js-tab-link tab-link--nav js-tab-link--nav is-selected" data-tab-id="driving" href="javascript:;" role="tab"
    >
            <span aria-label="Driving" style="width: 24px; height: 24px;" class="icon icon--24-car icon--size-24 icon--currentColor is-active tab-link_icon js-tab-link_icon tab-link_icon-wrap">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_car" />
    </svg>
</span>
    </a>


            </li>



            <li class="tab-nav_item" role="presentation">
                        





    <a aria-selected="false" class="tab-link js-tab-link tab-link--nav js-tab-link--nav" data-tab-id="transit" href="javascript:;" role="tab"
    >
            <span aria-label="Public Transit" style="width: 24px; height: 24px;" class="icon icon--24-transit icon--size-24 icon--currentColor tab-link_icon js-tab-link_icon tab-link_icon-wrap">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_transit" />
    </svg>
</span>
    </a>


            </li>



            <li class="tab-nav_item" role="presentation">
                        





    <a aria-selected="false" class="tab-link js-tab-link tab-link--nav js-tab-link--nav" data-tab-id="walking" href="javascript:;" role="tab"
    >
            <span aria-label="Walking" style="width: 24px; height: 24px;" class="icon icon--24-walk icon--size-24 icon--currentColor tab-link_icon js-tab-link_icon tab-link_icon-wrap">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_walk" />
    </svg>
</span>
    </a>


            </li>



            <li class="tab-nav_item tab-nav_item--last" role="presentation">
                        





    <a aria-selected="false" class="tab-link js-tab-link tab-link--nav js-tab-link--nav" data-tab-id="cycling" href="javascript:;" role="tab"
    >
            <span aria-label="Cycling" style="width: 24px; height: 24px;" class="icon icon--24-bicycle icon--size-24 icon--currentColor tab-link_icon js-tab-link_icon tab-link_icon-wrap">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_bicycle" />
    </svg>
</span>
    </a>


            </li>



            </ul>
        </div>

</div>
                <div class="js-starting-point u-space-b2">
        <a class="swapper pull-right read-more" href="javascript:;">Swap start/end points</a>
        <label>Start from</label>
        <div class="location">
            <div class="user-location nested-icon-label">
                <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-marker icon--size-18 icon--linked i">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_marker" />
    </svg>
</span>
                            <input autocomplete="off" id="js-dropper-text" name="location" type="text">


            </div>
        </div>
    </div>

                <div class="js-ending-point u-space-b2">
        <strong class="hidden">Start from</strong>
        <a class="swapper pull-right read-more hidden" href="javascript:;">Swap start/end points</a>
        <div class="location">
            <div class="business-location media-block">
                <div class="media-avatar">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-marker icon--size-18 icon--error">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_marker" />
    </svg>
</span>
                </div>
                <div class="media-story">
                    <div class="media-title">
                            


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="FM6GMdzV5FR83nQvEZJu3A" href="/biz/peter-luger-brooklyn-2"><span >Peter Luger</span></a>


                    </div>
                            <address>
        178 Broadway, Brooklyn, NY 11211
    </address>

                </div>
            </div>
        </div>
    </div>

            <button type="submit" value="submit" class="ybtn ybtn--primary ybtn--small get-directions-button"><span>Get directions</span></button>
        </form>

        <h3 class="result-header hidden"></h3>
        <div class="textual-results"></div>
    </div>

        </div>

    </div>

            </div>

        <div class="yelp-map-container directions-map">
            <div class="engine-container"></div>
        </div>
    </div>


        </div>
    </div>

            <div class="mapbox-text">
        <ul>
            <li class="u-relative">
                <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-marker icon--size-18 u-absolute u-sticky-top">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_marker" />
    </svg>
</span>
                        <a href="/biz_attribute?biz_id=4yPqqJDJOQX69gC66YUDkA" class="link-more icon-wrapper mapbox-edit">
            <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-pencil icon--size-14 icon--linked u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_pencil" />
    </svg>
</span><span>Edit</span>
        </a>
    <div class="map-box-address u-space-l4">
            <strong class="street-address">
                            <address>
        178 Broadway<br>Brooklyn, NY 11211
    </address>


            </strong>

                    <span class="cross-streets">
            b/t Driggs Ave & 6th St
        </span>
            <br>
                    <span class="neighborhood-str-list">
            Williamsburg - South Side, South Williamsburg        </span>


    </div>

            </li>

                <li class="clearfix">
                    <div>
                        <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-directions icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_directions" />
    </svg>
</span>
                            <a href="/map/peter-luger-brooklyn-2" class="biz-directions">Get Directions</a>

                    </div>
                </li>

                <li>
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-phone icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_phone" />
    </svg>
</span>
                            <span class="offscreen">Phone number</span>
        <span class="biz-phone">
            (718) 387-7400
        </span>

                </li>

                <li>
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-external-link icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_external_link" />
    </svg>
</span>    <span class="biz-website js-biz-website js-add-url-tagging">
        <span class="offscreen">Business website</span>
        <a href="/biz_redir?url=http%3A%2F%2Fwww.peterluger.com&amp;website_link_type=website&amp;src_bizid=4yPqqJDJOQX69gC66YUDkA&amp;cachebuster=1563162274&amp;s=c35a4d60dd779e850d7e1238385efddc0cf56b6d39b408176012801528cd9e90" target="_blank" rel="noopener">peterluger.com</a>
    </span>

                </li>



            <li class="clearfix">
                <div>
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-mobile icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_mobile" />
    </svg>
</span>
                        <a href="javascript:;" class="js-biz-to-phone">Send to your Phone</a>
                </div>
            </li>
        </ul>
    </div>

    </div>

            
        </div>
        <div class="showcase-container">
                

    <div class="showcase-container_inner showcase showcase-3-photo">

        <div class="top-shelf-grey"></div>

        <div class="showcase-footer-links">
            
        </div>



        <div class="lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="7692" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA" data-starting-index="0">
            <div class="showcase-photos">
                        





    <div class="js-photo photo photo-1" data-ga-label="left_photo" data-media-id="jOoRBpCU9_2iS8z306CGOQ" data-media-index="1">
        <div class="showcase-photo-box">
                    <a href="/biz_photos/peter-luger-brooklyn-2?select=jOoRBpCU9_2iS8z306CGOQ">
            
                <img alt="Photo of Peter Luger - Brooklyn, NY, United States. &#34;Oversized Drinks&#34;" class="photo-box-img" height="250" src="https://s3-media3.fl.yelpcdn.com/bphoto/jOoRBpCU9_2iS8z306CGOQ/ls.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/jOoRBpCU9_2iS8z306CGOQ/258s.jpg 1.03x,https://s3-media3.fl.yelpcdn.com/bphoto/jOoRBpCU9_2iS8z306CGOQ/348s.jpg 1.39x,https://s3-media3.fl.yelpcdn.com/bphoto/jOoRBpCU9_2iS8z306CGOQ/300s.jpg 1.20x" width="250">


        </a>

        </div>


                    <div class="photo-box-overlay js-overlay">
                <div class="media-block photo-box-overlay_caption">
                        <div class="media-avatar avatar">
                    <div class="photo-box pb-30s" data-hovercard-id="voGn4qT5ps9LEBinTy8qug">
                <a href="/user_details?userid=WU3ImaespusxxBHyTp3qmg" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Stephanie G." class="photo-box-img" height="30" src="https://s3-media2.fl.yelpcdn.com/photo/zVLKLMbG-12cnxd7wFGEvg/30s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/photo/zVLKLMbG-12cnxd7wFGEvg/60s.jpg 2.00x,https://s3-media2.fl.yelpcdn.com/photo/zVLKLMbG-12cnxd7wFGEvg/90s.jpg 3.00x,https://s3-media2.fl.yelpcdn.com/photo/zVLKLMbG-12cnxd7wFGEvg/ss.jpg 1.33x" width="30">

        </a>

    </div>



    </div>

                        <div class="media-story">
            <a class="photo-desc" href="/biz_photos/peter-luger-brooklyn-2?select=jOoRBpCU9_2iS8z306CGOQ">
                &#34;Oversized Drinks&#34;
            </a>
        <span class="author">
                by         <a class="user-display-name js-analytics-click" href="/user_details?userid=WU3ImaespusxxBHyTp3qmg" data-hovercard-id="voGn4qT5ps9LEBinTy8qug" data-analytics-label="about_me" id="dropdown_user-name">Stephanie G.</a>
        </span>
    </div>

                </div>
        </div>

    </div>

        





    <div class="js-photo photo photo-2" data-ga-label="middle_photo" data-media-id="DnReRUkXRtsmKycQEYl0pg" data-media-index="0">
        <div class="showcase-photo-box">
                    <a href="/biz_photos/peter-luger-brooklyn-2?select=DnReRUkXRtsmKycQEYl0pg">
            
                <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Special German Fried Potatoes (for 2) $12.95 as served by waiter, with 2 raw pieces of steak... I mean &#34;medium rare&#34;" class="photo-box-img" height="250" src="https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/ls.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/258s.jpg 1.03x,https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/348s.jpg 1.39x,https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/300s.jpg 1.20x" width="250">


        </a>

        </div>


                    <div class="photo-box-overlay js-overlay">
                <div class="media-block photo-box-overlay_caption">
                        <div class="media-avatar avatar">
                    <div class="photo-box pb-30s" data-hovercard-id="jhcC53NaX4oPCiMVW_V-pg">
                <a href="/user_details?userid=0H2ggT37lK8hsbH8aGR1Ug" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Dean C." class="photo-box-img" height="30" src="https://s3-media1.fl.yelpcdn.com/photo/7NI7TCJ3yavVgoncyPUS2Q/30s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/photo/7NI7TCJ3yavVgoncyPUS2Q/60s.jpg 2.00x,https://s3-media1.fl.yelpcdn.com/photo/7NI7TCJ3yavVgoncyPUS2Q/90s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/photo/7NI7TCJ3yavVgoncyPUS2Q/ss.jpg 1.33x" width="30">

        </a>

    </div>



    </div>

                        <div class="media-story">
            <a class="photo-desc" href="/biz_photos/peter-luger-brooklyn-2?select=DnReRUkXRtsmKycQEYl0pg">
                Luger&#39;s Special German Fried Potatoes…
            </a>
        <span class="author">
                by         <a class="user-display-name js-analytics-click" href="/user_details?userid=0H2ggT37lK8hsbH8aGR1Ug" data-hovercard-id="jhcC53NaX4oPCiMVW_V-pg" data-analytics-label="about_me" id="dropdown_user-name">Dean C.</a>
        </span>
    </div>

                </div>
        </div>

    </div>

        





    <div class="js-photo photo photo-3 photo-grid" data-ga-label="right_photo" data-media-id="KRpKd1ZdcYSPRDhlb12Vjw" data-media-index="2">
        <div class="showcase-photo-box">
                    <a href="/biz_photos/peter-luger-brooklyn-2?select=KRpKd1ZdcYSPRDhlb12Vjw">
            
                <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Big slice o bacon" class="photo-box-img" height="250" src="https://s3-media4.fl.yelpcdn.com/bphoto/KRpKd1ZdcYSPRDhlb12Vjw/ls.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/KRpKd1ZdcYSPRDhlb12Vjw/258s.jpg 1.03x,https://s3-media4.fl.yelpcdn.com/bphoto/KRpKd1ZdcYSPRDhlb12Vjw/348s.jpg 1.39x,https://s3-media4.fl.yelpcdn.com/bphoto/KRpKd1ZdcYSPRDhlb12Vjw/300s.jpg 1.20x" width="250">


        </a>

                            <a href="/biz_photos/peter-luger-brooklyn-2?select=IoUOmcfnYddp9NL4rUVJmw">
            
                <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Special &#34;Holy Cow&#34; Hot Fudge Sundae $13.95" class="photo-box-img" height="180" src="https://s3-media4.fl.yelpcdn.com/bphoto/IoUOmcfnYddp9NL4rUVJmw/180s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/IoUOmcfnYddp9NL4rUVJmw/ls.jpg 1.39x,https://s3-media4.fl.yelpcdn.com/bphoto/IoUOmcfnYddp9NL4rUVJmw/300s.jpg 1.67x,https://s3-media4.fl.yelpcdn.com/bphoto/IoUOmcfnYddp9NL4rUVJmw/258s.jpg 1.43x,https://s3-media4.fl.yelpcdn.com/bphoto/IoUOmcfnYddp9NL4rUVJmw/348s.jpg 1.93x" width="180">


        </a>

                            <a href="/biz_photos/peter-luger-brooklyn-2?select=vcU4N1jvvQl-8cEUBCeSBA">
            
                <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak For Two $107.90 &#34;USDA Prime Beef, family selected, dry aged in our own aging box&#34;" class="photo-box-img" height="180" src="https://s3-media3.fl.yelpcdn.com/bphoto/vcU4N1jvvQl-8cEUBCeSBA/180s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/vcU4N1jvvQl-8cEUBCeSBA/ls.jpg 1.39x,https://s3-media3.fl.yelpcdn.com/bphoto/vcU4N1jvvQl-8cEUBCeSBA/300s.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/bphoto/vcU4N1jvvQl-8cEUBCeSBA/258s.jpg 1.43x,https://s3-media3.fl.yelpcdn.com/bphoto/vcU4N1jvvQl-8cEUBCeSBA/348s.jpg 1.93x" width="180">


        </a>

                            <a href="/biz_photos/peter-luger-brooklyn-2?select=qekn1pkUON5Q8efRdoUH8Q">
            
                <img alt="Photo of Peter Luger - Brooklyn, NY, United States. A Piece of My Medium Rare Steak (yes, it&#39;s raw)" class="photo-box-img" height="180" src="https://s3-media2.fl.yelpcdn.com/bphoto/qekn1pkUON5Q8efRdoUH8Q/180s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/qekn1pkUON5Q8efRdoUH8Q/ls.jpg 1.39x,https://s3-media2.fl.yelpcdn.com/bphoto/qekn1pkUON5Q8efRdoUH8Q/300s.jpg 1.67x,https://s3-media2.fl.yelpcdn.com/bphoto/qekn1pkUON5Q8efRdoUH8Q/258s.jpg 1.43x,https://s3-media2.fl.yelpcdn.com/bphoto/qekn1pkUON5Q8efRdoUH8Q/348s.jpg 1.93x" width="180">


        </a>

        </div>


                <a class="see-more show-all-overlay" href="/biz_photos/peter-luger-brooklyn-2">
        <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-grid icon--size-24 icon--inverse icon--fallback-inverted show-all-overlay_icon">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_grid" />
    </svg>
</span>
        See all 7692
    </a>

    </div>


            </div>
        </div>

    </div>

        </div>
    </div>

        </div>
    </div>

                    <div id="super-container" class="content-container">


    <div class="container">
    <div class="clearfix layout-block layout-a layout-border ysection js-position-with-scroll-container position-with-scroll-container layout--biz-details">
    <div class="column column-alpha main-section">

                





                    
    <div class="popular-dishes js-popular-dishes popular-dishes--for-status-quo">
            
    <div class="popular-dishes-section-header arrange arrange--baseline arrange--12 u-space-b1">
            <div class="section-title arrange_unit arrange_unit--fill">
            <h2 class="u-space-b0">
                Popular Dishes
            </h2>
    </div>

            <div class="full-menu-url js-full-menu-url arrange_unit u-nowrap">
        <a href="https://www.yelp.com/biz_redir?cachebuster=1563162273&amp;s=eaba562b9cf91311153e81d45ca3452ab74d246ec165d3bd1943618a543a5d70&amp;src_bizid=4yPqqJDJOQX69gC66YUDkA&amp;url=http%3A%2F%2Fpeterluger.com%2Fbrooklyn-menu%2F&amp;website_link_type=menu" target="_blank">View Full Menu</a>
    </div>


                <div class="carousel-navigator carousel-navigator_disabled js-left-navigator arrange_unit u-text-centered ">
        <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-chevron-left icon--size-14">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_chevron_left" />
    </svg>
</span>
    </div>
    <div class="carousel-navigator js-right-navigator arrange_unit u-text-centered">
        <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-chevron-right icon--size-14">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_chevron_right" />
    </svg>
</span>
    </div>

    </div>

                <div class="popular-dishes-list popular-dishes-list--carousel js-carousel-slider arrange">
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="0">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/creamed-spinach-for-2" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media4.fl.yelpcdn.com/bphoto/c07W-yn_NdvnE19sE_aAog/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Creamed Spinach
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                173 Photos
            </small>

            <small class="review-count bullet-before">
                1350 Reviews
            </small>
        </div>
    </div>

            </div>
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="1">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/apple-strudel" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media3.fl.yelpcdn.com/bphoto/rtAqMUEqIE3rWucrQuormw/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Apple Strudel
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                59 Photos
            </small>

            <small class="review-count bullet-before">
                169 Reviews
            </small>
        </div>
    </div>

            </div>
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="2">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/pecan-pie" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media4.fl.yelpcdn.com/bphoto/D-Qbx49XV_urCqiGIX0HQA/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Pecan Pie
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                26 Photos
            </small>

            <small class="review-count bullet-before">
                130 Reviews
            </small>
        </div>
    </div>

            </div>
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="3">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/iceberg-wedge-salad" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media1.fl.yelpcdn.com/bphoto/0E8NhyZGBbBE5m1wJjzjLw/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Iceberg Wedge Salad
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                23 Photos
            </small>

            <small class="review-count bullet-before">
                71 Reviews
            </small>
        </div>
    </div>

            </div>
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="4">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/jumbo-shrimp-cocktail-extra-large-portion-6" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media3.fl.yelpcdn.com/bphoto/7cJsWx5AcOvXltPbsx3CIw/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Jumbo Shrimp Cocktail, Extra Large Portion
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                15 Photos
            </small>

            <small class="review-count bullet-before">
                59 Reviews
            </small>
        </div>
    </div>

            </div>
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="5">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/rib-steak" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media2.fl.yelpcdn.com/bphoto/Kg4VVHIqYnDiyn3P66fkdw/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Rib Steak
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                45 Photos
            </small>

            <small class="review-count bullet-before">
                60 Reviews
            </small>
        </div>
    </div>

            </div>
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="6">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/single-steak" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media4.fl.yelpcdn.com/bphoto/JFTrKejowhqWrAy3bL4C5Q/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Single Steak
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                16 Photos
            </small>

            <small class="review-count bullet-before">
                54 Reviews
            </small>
        </div>
    </div>

            </div>
            <div class="js-popular-dish-item arrange_unit u-space-r2" data-dish-index="7">
                    <div class="popular-dish-content">
        <a href="/menu/peter-luger-brooklyn-2/item/roast-prime-ribs-of-beef-usda-prime-4" class="popular-dish-menu-item-link">
            <span></span>
        </a>

        <div class="popular-dish-image-price-card u-space-b2">
            <img class="popular-dish-image" src="https://s3-media1.fl.yelpcdn.com/bphoto/Usd-p4b0XTHSXhl76kl-uA/258s.jpg">
        </div>

        <div class="h4 alternate u-text-truncate popular-dish-title">
            Roast Prime Ribs of Beef Usda Prime
        </div>

        <div class="popular-dish-additional-info u-text-subtle">
            <small class="photo-count">
                14 Photos
            </small>

            <small class="review-count bullet-before">
                64 Reviews
            </small>
        </div>
    </div>

            </div>
    </div>

    </div>


                    

                    <div class="media-details js-media-details js-media-details-template hidden">
        <div class="media-details_container media-details_container--embed media-details_container--with-sidebar">
            <div class="media-container js-media-container"></div>
            <div class="media-nav js-media-nav"></div>
        </div>
    </div>


                        <img class="offscreen" src="https://idsync.rlcdn.com/398766.gif?partner_uid=DE1CA55AE6AF7DF3" height="0" width="0">





                <div >

                    
        <div class="ysection questions">
                    
    <div class="section-header u-space-b0">
        <h2>Ask the Community</h2>
    </div>


                    <ul class="questions_feed ylist ylist-bordered u-space-b3">
                            <li>
                                

            <div class="arrange arrange--6">
        <div class="arrange_unit arrange_unit--fill">
            <h3 class="alternate u-line-break">How do I reserve? I&#39;ve called a bunch of times but their line is always busy!</h3>
        </div>
    </div>

    <div class="u-space-t1">
                        
    <div class="arrange arrange--12">
        <div  class="arrange_unit">
                        <div class="photo-box pb-30s" data-hovercard-id="qC4YYEVpoWav1820z1WVsQ">
                <a href="/user_details?userid=CIAFysUee8isjjxGLp4uSw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Kelsey A." class="photo-box-img" height="30" src="https://s3-media4.fl.yelpcdn.com/photo/Hn4Thc_4TtBYLvdtpeNISA/30s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/photo/Hn4Thc_4TtBYLvdtpeNISA/60s.jpg 2.00x,https://s3-media4.fl.yelpcdn.com/photo/Hn4Thc_4TtBYLvdtpeNISA/90s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/photo/Hn4Thc_4TtBYLvdtpeNISA/ss.jpg 1.33x" width="30">

        </a>

    </div>



        </div>
        <div class="arrange_unit arrange_unit--fill">
            
    <div class="js-answer-expander">
        <p class="js-content-expanded u-hidden u-line-break u-break-word u-space-b1">I had to call multiple times, multiple days. I called about 4PM and was on hold for about fifteen minutes before the picked up. My reservation is about a month away and they still had only a select amount of time slots. Cal about a month-month and a half in advance and block off about twenty minutes.             <a href="javascript:;" class="js-read-less">Read less</a>
        </p>
        <p class="js-content-truncated u-line-break u-break-word u-space-b1">I had to call multiple times, multiple days. I called about 4PM and was on hold for about fifteen minutes before the picked up. My reservation is about a month away and they still had only a select amount of time slots. Cal about a month-month and a…                 <a href="javascript:;" class="js-read-more">Read more</a>
        </p>
    </div>


            
    <span class="time-stamp u-time-stamp">
        7 months ago
    </span>

                

        <small class="feedback-summary u-text-subtle bullet-before">
            18 people found this helpful
        </small>



        </div>
    </div>




    </div>


    <div class="u-space-t2">
            <div class="u-space-t2">
                <a class="question-details" href="/questions/peter-luger-how-do-i-reserve-ive-called-a-bunch-of-times-but-their-line-is-always-busy/FzyF96AazelmeUbIsNAVew">
                    View 24 more answers
                </a>
            </div>
    </div>

                            </li>
                            <li>
                                

            <div class="arrange arrange--6">
        <div class="arrange_unit arrange_unit--fill">
            <h3 class="alternate u-line-break">How much does it cost for 1 couple to eat here?</h3>
        </div>
    </div>

    <div class="u-space-t1">
                        
    <div class="arrange arrange--12">
        <div  class="arrange_unit">
                        <div class="photo-box pb-30s" data-hovercard-id="qg7N_fGlWw4h33LX-Pks-Q">
                <a href="/user_details?userid=mi-2YfIj88R2VlDy0DAvxg" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Maki S." class="photo-box-img" height="30" src="https://s3-media4.fl.yelpcdn.com/photo/lzive9Cxy6NSttvO2g83UA/30s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/photo/lzive9Cxy6NSttvO2g83UA/60s.jpg 2.00x,https://s3-media4.fl.yelpcdn.com/photo/lzive9Cxy6NSttvO2g83UA/90s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/photo/lzive9Cxy6NSttvO2g83UA/ss.jpg 1.33x" width="30">

        </a>

    </div>



        </div>
        <div class="arrange_unit arrange_unit--fill">
            
    <div class="js-answer-expander">
        <p class="js-content-expanded u-hidden u-line-break u-break-word u-space-b1">My bill was $152 before tips with 2 pieces of bacon, steak for two, creamed spinach for 2, and 2 sodas.    If you order wine, it&#39;s easily over $200.             <a href="javascript:;" class="js-read-less">Read less</a>
        </p>
        <p class="js-content-truncated u-line-break u-break-word u-space-b1">My bill was $152 before tips with 2 pieces of bacon, steak for two, creamed spinach for 2, and 2 sodas.    If you order wine, it&#39;s easily over $200.         </p>
    </div>


            
    <span class="time-stamp u-time-stamp">
        1 year ago
    </span>

                

        <small class="feedback-summary u-text-subtle bullet-before">
            24 people found this helpful
        </small>



        </div>
    </div>




    </div>


    <div class="u-space-t2">
            <div class="u-space-t2">
                <a class="question-details" href="/questions/peter-luger-how-much-does-it-cost-for-1-couple-to-eat-here/6z2t09QbZ1lQhjhFzTx3DA">
                    View 19 more answers
                </a>
            </div>
    </div>

                            </li>
                    </ul>

                <a href="/questions/peter-luger-brooklyn-2" class="u-cursor-pointer">
                    See all 43 questions for Peter Luger
                </a>
        </div>

                    

    <div class="feed">
                <div class="feed_header">
            <div class="section-header section-header--no-spacing">
                        <h2>Recommended Reviews <b>for Peter Luger</b></h2>

                        <div class="feed_trust-banner">
            <div class="arrange arrange--12 arrange--middle">
                <div class="arrange_unit">
                    <span aria-hidden="true" style="fill: #c41200; width: 24px; height: 24px;" class="icon icon--24-yelp icon--size-24">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_yelp" />
    </svg>
</span>
                </div>
                <div class="arrange_unit arrange_unit--fill">
                    <span class="legal-copy">
                            <b>Your trust is our top concern,</b> so businesses can't pay to alter or remove their reviews. <a href="/advertiser_faq">Learn more.</a>
                    </span>
                </div>
                <div class="arrange_unit">
                    <span class="dismiss-link u-text-mid js-dismiss-trust-banner" role="button" aria-label="Close">
                        &times;
                    </span>
                </div>
            </div>
        </div>

                    <div class="feed_filters u-space-t1 u-space-b1">
                <div class="section-header_block u-space-0">
                    <div class="arrange arrange--middle">
                            <div class="arrange_unit arrange_unit--fill feed_search">
                                <div class="section-header_search u-space-r5">
                                        <form class="yform yform--continuous arrange" name="q" action="https://www.yelp.com/biz/peter-luger-brooklyn-2" method="GET">
            <label for="q" class="offscreen">Search within the reviews</label>
    <div class="arrange_unit arrange_unit--fill">
        <input type="text" placeholder="Search within the reviews" name="q" value="" autocomplete="on" >
    </div>
    <div class="arrange_unit">
        <button type="submit" value="submit" class="ybtn ybtn--primary ybtn--small"><span><span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-search-small icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_search_small" />
    </svg>
</span></span></button>
    </div>

    </form>

                                </div>
                            </div>
                            <div class="arrange_unit u-nowrap">
                                <div class="feed_sort js-review-feed-sort">
                                    

    <div class="dropdown js-dropdown dropdown--tab dropdown--arrow dropdown--hover dropdown--restricted">
        

    <div class="dropdown_toggle js-dropdown-toggle" tabindex="-1">
        <a
            class="dropdown_toggle-action"
                href="javascript:;"
                data-dropdown-prefix="Sort by"
            aria-haspopup="true"
            role="button"
        >
                <span class="dropdown_prefix">
                    Sort by
                </span>
            <span class="dropdown_toggle-text js-dropdown-toggle-text" data-dropdown-initial-text="Yelp Sort">Yelp Sort</span>
            <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-triangle-down icon--size-14 icon--currentColor u-triangle-direction-down dropdown_arrow">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_triangle_down" />
    </svg>
</span>
        </a>
    </div>

            <div class="dropdown_menu-container">
        <div class="dropdown_menu js-dropdown-menu">
            <div class="dropdown_menu-inner">
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="true" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown is-selected" data-review-feed-label="Yelp Sort" data-sort="relevance" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?sort_by=relevance_desc&amp;start=0" role="tab"
    >
                <span class="tab-link_label">Yelp Sort</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-order_by="desc" data-review-feed-label="Newest First" data-sort="date" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?sort_by=date_desc&amp;start=0" role="tab"
    >
                <span class="tab-link_label">Newest First</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-order_by="asc" data-review-feed-label="Oldest First" data-sort="date" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?sort_by=date_asc&amp;start=0" role="tab"
    >
                <span class="tab-link_label">Oldest First</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-order_by="desc" data-review-feed-label="Highest Rated" data-sort="rating" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?sort_by=rating_desc&amp;start=0" role="tab"
    >
                <span class="tab-link_label">Highest Rated</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-order_by="asc" data-review-feed-label="Lowest Rated" data-sort="rating" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?sort_by=rating_asc&amp;start=0" role="tab"
    >
                <span class="tab-link_label">Lowest Rated</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-review-feed-label="Elites" data-sort="elites" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?sort_by=elites_desc&amp;start=0" role="tab"
    >
                <span class="tab-link_label">Elites</span>
    </a>

                            </li>
                    </ul>
            </div>
        </div>
    </div>

    </div>

                                </div>
                            </div>
                            <div class="arrange_unit u-nowrap feed_language js-review-feed-language dropdown--right">
                                    



        <div class="tab-nav-container">
            <ul class="tab-nav js-tab-nav" role="tablist">
                    
            <li class="tab-nav_item tab-nav_item--last" role="presentation">
                        

    <div class="dropdown js-dropdown dropdown--tab dropdown--arrow dropdown--hover dropdown--restricted">
        

    <div class="dropdown_toggle js-dropdown-toggle" tabindex="-1">
        <a
            class="dropdown_toggle-action"
                href="javascript:;"
                data-dropdown-prefix="Language"
            aria-haspopup="true"
            role="button"
        >
                <span class="dropdown_prefix">
                    Language
                </span>
            <span class="dropdown_toggle-text js-dropdown-toggle-text" data-dropdown-initial-text="English (5385)">English (5385)</span>
            <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-triangle-down icon--size-14 icon--currentColor u-triangle-direction-down dropdown_arrow">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_triangle_down" />
    </svg>
</span>
        </a>
    </div>

            <div class="dropdown_menu-container">
        <div class="dropdown_menu js-dropdown-menu">
            <div class="dropdown_menu-inner">
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item" role="presentation">
                                





    <span aria-selected="true" class="tab-link tab-link--dropdown is-selected" data-lang="en" role="tab"
    >
                <span class="tab-link_label">English</span>
            <span class="tab-link_count">(5385)</span>
    </span>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="de" href="https://www.yelp.de/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">German</span>
            <span class="tab-link_count">(22)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="zh" href="https://zh.yelp.com.hk/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">Chinese</span>
            <span class="tab-link_count">(10)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="ja" href="https://www.yelp.co.jp/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">Japanese</span>
            <span class="tab-link_count">(10)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="es" href="https://www.yelp.es/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">Spanish</span>
            <span class="tab-link_count">(3)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="fr" href="https://www.yelp.fr/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">French</span>
            <span class="tab-link_count">(2)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="sv" href="https://www.yelp.se/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">Swedish</span>
            <span class="tab-link_count">(2)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="nl" href="https://www.yelp.nl/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">Dutch</span>
            <span class="tab-link_count">(1)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="pt" href="https://www.yelp.com.br/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">Portuguese</span>
            <span class="tab-link_count">(1)</span>
    </a>

                            </li>
                            <li class="dropdown_item" role="presentation">
                                





    <a aria-selected="false" class="tab-link js-dropdown-link tab-link--dropdown js-tab-link--dropdown" data-lang="it" href="https://www.yelp.it/biz/peter-luger-brooklyn-2" role="tab"
    >
                <span class="tab-link_label">Italian</span>
            <span class="tab-link_count">(1)</span>
    </a>

                            </li>
                    </ul>
            </div>
        </div>
    </div>

    </div>


            </li>



            </ul>
        </div>


                            </div>
                    </div>
                </div>
    </div>

                
            </div>
    </div>

        <div class="review-list">
                <ul class="ylist ylist-bordered reviews">
                <li>
                    <div class="js-war-widget war-widget--compose review review--with-sidebar">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <img src="https://s3-media4.fl.yelpcdn.com/assets/2/www/img/1f2e356daa5c/writeareview/empty_profile.png" srcset="https://s3-media2.fl.yelpcdn.com/assets/2/www/img/f52f768da99a/writeareview/empty_profile@2x.png 2x" width="148"
                    height="68">
        </div>
    </div>


        <div class='review-wrapper'>
            <div class="island clearfix u-text-centered">
                <div class="u-border-bottom u-space-b2 war-widget_stars--top">
                    <div class="u-space-b2">
                            

    <fieldset data-war-url="/writeareview/biz/4yPqqJDJOQX69gC66YUDkA?return_url=%2Fbiz%2F4yPqqJDJOQX69gC66YUDkA" class="star-selector js-star-selector" data-original-rating="0">
        <legend class="star-selector_legend offscreen">Rating</legend>
        <ul class="star-selector_stars i-selector-stars i-selector-stars--extra-large-0 js-star-selector_stars">
                <li class="star-selector_star js-star-selector_star show-tooltip" data-label="Eek! Methinks not.">
                    <input class="star-selector_input js-star-selector_input" id="rating-1" name="rating" type="radio" value="1">
                    <label class="star-selector_label" for="rating-1">1 (Eek! Methinks not.)</label>
                </li>
                <li class="star-selector_star js-star-selector_star show-tooltip" data-label="Meh. I've experienced better.">
                    <input class="star-selector_input js-star-selector_input" id="rating-2" name="rating" type="radio" value="2">
                    <label class="star-selector_label" for="rating-2">2 (Meh. I've experienced better.)</label>
                </li>
                <li class="star-selector_star js-star-selector_star show-tooltip" data-label="A-OK.">
                    <input class="star-selector_input js-star-selector_input" id="rating-3" name="rating" type="radio" value="3">
                    <label class="star-selector_label" for="rating-3">3 (A-OK.)</label>
                </li>
                <li class="star-selector_star js-star-selector_star show-tooltip" data-label="Yay! I'm a fan.">
                    <input class="star-selector_input js-star-selector_input" id="rating-4" name="rating" type="radio" value="4">
                    <label class="star-selector_label" for="rating-4">4 (Yay! I'm a fan.)</label>
                </li>
                <li class="star-selector_star js-star-selector_star show-tooltip star-selector_star--last" data-label="Woohoo! As good as it gets!">
                    <input class="star-selector_input js-star-selector_input" id="rating-5" name="rating" type="radio" value="5">
                    <label class="star-selector_label" for="rating-5">5 (Woohoo! As good as it gets!)</label>
                </li>
        </ul>
    </fieldset>


                    </div>
                </div>
                <a class="js-war-text-link" href="/writeareview/biz/4yPqqJDJOQX69gC66YUDkA?return_url=%2Fbiz%2F4yPqqJDJOQX69gC66YUDkA">
                    Start your review of <strong>Peter Luger</strong>.
                </a>
            </div>
        </div>
    </div>


    </li>


                    <li>
                <div class="review review--with-sidebar" data-review-id="ttR7bpdlggK9dNJeldRSwg" data-signup-object="user_id:PpPJs8yG-sovCiuT6QYr3w">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="XYjtdA9XQl27dQGXBu5s3w">
                <a href="/user_details?userid=PpPJs8yG-sovCiuT6QYr3w" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Jordyn S." class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/photo/WKfNvKkO21QDgU_cPQ9r5w/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/photo/WKfNvKkO21QDgU_cPQ9r5w/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/photo/WKfNvKkO21QDgU_cPQ9r5w/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/photo/WKfNvKkO21QDgU_cPQ9r5w/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/photo/WKfNvKkO21QDgU_cPQ9r5w/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/photo/WKfNvKkO21QDgU_cPQ9r5w/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=PpPJs8yG-sovCiuT6QYr3w" data-hovercard-id="XYjtdA9XQl27dQGXBu5s3w" data-analytics-label="about_me" id="dropdown_user-name">Jordyn S.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>San Diego, CA</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>28</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>47</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>50</b> photos
            </li>
        
    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/ttR7bpdlggK9dNJeldRSwg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/ttR7bpdlggK9dNJeldRSwg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Jordyn S.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="PpPJs8yG-sovCiuT6QYr3w">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Jordyn S.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="PpPJs8yG-sovCiuT6QYr3w">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Jordyn S.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        7/8/2019
    </span>

    </div>

        
                <p lang="en">It had been about 10 years since the last time we went. Nothing has changed which I love it!! The. Best. Steak. Ever. <br><br>My now teens used to come when they were toddlers and they didn&#39;t use to eat, or like steaks much, but now... They killed it. <br><br>I&#39;ve been on diet and don&#39;t eat dinner much for the past 2 years, but I did it. I ate some steaks and even dessert, a cheesecake with a whole a lot of cream on it. I shared it with my younger daughter, but we are it all!<br><br>Love the atmosphere, food and service. This is a MUST place when we visit NYC.</p> 
            
    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="ttR7bpdlggK9dNJeldRSwg" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
            Was this review &hellip;?
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="l1-pSziyvy3aUiOOAN4F5w" data-signup-object="user_id:WMm0cPBWeYz-4aKTTCwfZA">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="FNnzpBMYsYv8zqMRIY7b0A">
                <a href="/user_details?userid=WMm0cPBWeYz-4aKTTCwfZA" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Guadalupe L." class="photo-box-img" height="60" src="https://s3-media2.fl.yelpcdn.com/photo/emCd_wIUt8FD4qe6vtE2Mg/60s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/photo/emCd_wIUt8FD4qe6vtE2Mg/90s.jpg 1.50x,https://s3-media2.fl.yelpcdn.com/photo/emCd_wIUt8FD4qe6vtE2Mg/168s.jpg 2.80x,https://s3-media2.fl.yelpcdn.com/photo/emCd_wIUt8FD4qe6vtE2Mg/ms.jpg 1.67x,https://s3-media2.fl.yelpcdn.com/photo/emCd_wIUt8FD4qe6vtE2Mg/180s.jpg 3.00x,https://s3-media2.fl.yelpcdn.com/photo/emCd_wIUt8FD4qe6vtE2Mg/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=WMm0cPBWeYz-4aKTTCwfZA" data-hovercard-id="FNnzpBMYsYv8zqMRIY7b0A" data-analytics-label="about_me" id="dropdown_user-name">Guadalupe L.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>New York, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>1</b> friend
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>43</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>1</b> photo
            </li>
        
    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/l1-pSziyvy3aUiOOAN4F5w" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/l1-pSziyvy3aUiOOAN4F5w" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Guadalupe L.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="WMm0cPBWeYz-4aKTTCwfZA">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Guadalupe L.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="WMm0cPBWeYz-4aKTTCwfZA">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Guadalupe L.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-3 rating-large" title="3.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="3.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        7/5/2019
    </span>

    </div>

        
                <p lang="en">The cuts of meat are delicious, the service is good, the prices are above average, its location is somewhat removed from Manhattan, I suggest a remodeling to give it a more fresh and innovative look.<br><br>Los cortes de carne son  deliciosos, el servicio es bueno, los precios por arriba del promedio, su ubicacion es algo retirada de Manhattan, suguiero una remodelacion para darle un look mas fresco e inovativo</p> 
            
    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="l1-pSziyvy3aUiOOAN4F5w" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Matthew B. and 2 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="vtdTqDJw8sufc6kr0UffYQ" data-signup-object="user_id:ElEUh1-Ux6npfDs56sGPSw">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="VFY28x6UCmRZ2os3Ug1ymw">
                <a href="/user_details?userid=ElEUh1-Ux6npfDs56sGPSw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Wendy S." class="photo-box-img" height="60" src="https://s3-media2.fl.yelpcdn.com/photo/_44QLq0toYqajvPi4T2VUw/60s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/photo/_44QLq0toYqajvPi4T2VUw/90s.jpg 1.50x,https://s3-media2.fl.yelpcdn.com/photo/_44QLq0toYqajvPi4T2VUw/168s.jpg 2.80x,https://s3-media2.fl.yelpcdn.com/photo/_44QLq0toYqajvPi4T2VUw/ms.jpg 1.67x,https://s3-media2.fl.yelpcdn.com/photo/_44QLq0toYqajvPi4T2VUw/180s.jpg 3.00x,https://s3-media2.fl.yelpcdn.com/photo/_44QLq0toYqajvPi4T2VUw/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=ElEUh1-Ux6npfDs56sGPSw" data-hovercard-id="VFY28x6UCmRZ2os3Ug1ymw" data-analytics-label="about_me" id="dropdown_user-name">Wendy S.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Tucson, AZ</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>14</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>122</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>292</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/vtdTqDJw8sufc6kr0UffYQ" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/vtdTqDJw8sufc6kr0UffYQ" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Wendy S.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="ElEUh1-Ux6npfDs56sGPSw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Wendy S.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="ElEUh1-Ux6npfDs56sGPSw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Wendy S.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        7/4/2019
    </span>

    </div>

        
                <p lang="en">This is the place where beef becomes steak and steak becomes manna. After all the hoopla surrounding Peter Luger, it was without question that I had to pay a visit and digest some of their world-renowned steaks. I had made a reservation through their website a month in advance for dinner on a Thursday at 7:45 p.m. to celebrate my beloved sister&#39;s birthday. The online reservation system was very user-friendly and convenient. <br><br>We arrived 15 minutes early and were told to wait a few minutes for our table. We didn&#39;t have to wait long and were seated next to one of the large windows overlooking Broadway Street. Our waiter was jovial, knowledgeable and efficient. He had a dry sense of humor which we found amusing. <br><br>We started off with the Caesar salad which managed to be amazing with crunchy, fresh lettuce and grated Pecorino Romano. We ordered the Steak for Two (which is the Porterhouse steak) and the Rib Steak which has more marbling and fat. We also chose two sides, the creamed spinach and the Luger&#39;s Special German fried potatoes. Both sides are good for 2 people. <br><br>When the food arrived, we took one look at our overladen table and predicted that leftovers would definitely result due to our ambitious ordering strategy. But against all odds and to everyone&#39;s amazement, no steak, potato or spinach was left standing. We looked with pride and a sense of accomplishment at the bare bones left on our otherwise empty plates. My sister and I both concluded that we had just experienced meat nirvana. My sister preferred the Porterhouse and I liked the Rib-eye more. Our waiter also recommended eating the steaks sans any gravy or sauce in order to fully enjoy the seasoning and core flavors of the meats.<br><br>Payment choices include cash, personal check with ID and debit card. We also got a pile of milk chocolate Peter Luger coins at the end of our incredible meal. We were too stuffed to order any desserts.<br><br>They have a tiny nook of a gift shop in the restaurant which sells Peter Luger merchandise including the Luger sauce, but one can only buy meats online. They ship all over the United States and they accept credit cards for online purchases.<br><br>I encourage people to enter the world of Peter Luger where great food and German beer-hallesque ambience abound.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="1" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=vtdTqDJw8sufc6kr0UffYQ" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=A4amdOgcc9LtSxh3Rjh5nA>
        
                <img data-async-src=https://s3-media1.fl.yelpcdn.com/bphoto/A4amdOgcc9LtSxh3Rjh5nA/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Left - Steak for two; Right - Rib Steak" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Left - Steak for two; Right - Rib Steak" class="photo-box-img" height="348" src="https://s3-media1.fl.yelpcdn.com/bphoto/A4amdOgcc9LtSxh3Rjh5nA/348s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/A4amdOgcc9LtSxh3Rjh5nA/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Left - Steak for two; Right - Rib Steak</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=vtdTqDJw8sufc6kr0UffYQ&amp;select=A4amdOgcc9LtSxh3Rjh5nA">
            <span class="offscreen">Left - Steak for two; Right - Rib Steak</span>
    </a>

    </div>

                </li>

        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="vtdTqDJw8sufc6kr0UffYQ" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Larry F. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">1</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="Iu_X4aMAsTwxsEgx9pLYBw" data-signup-object="user_id:XbWP6vAli3waJ0iR_lcjOg">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="hB232lBxxW1Ga2gbX9Ba-Q">
                <a href="/user_details?userid=XbWP6vAli3waJ0iR_lcjOg" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Nathan A." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/hv3wlO_oCqUy6eOUtPKslA/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/hv3wlO_oCqUy6eOUtPKslA/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/hv3wlO_oCqUy6eOUtPKslA/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/hv3wlO_oCqUy6eOUtPKslA/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/hv3wlO_oCqUy6eOUtPKslA/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/hv3wlO_oCqUy6eOUtPKslA/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=XbWP6vAli3waJ0iR_lcjOg" data-hovercard-id="hB232lBxxW1Ga2gbX9Ba-Q" data-analytics-label="about_me" id="dropdown_user-name">Nathan A.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Toronto, Canada</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>119</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>110</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>184</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/Iu_X4aMAsTwxsEgx9pLYBw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/Iu_X4aMAsTwxsEgx9pLYBw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Nathan A.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="XbWP6vAli3waJ0iR_lcjOg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Nathan A.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="XbWP6vAli3waJ0iR_lcjOg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Nathan A.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-4 rating-large" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        6/28/2019
    </span>

    </div>

        
                <p lang="en">Visiting from Toronto, knew we had to stop by here to try out the famous Peter Luger. Came as a party of 6, with a reservation. Even with the reservation expect to chill a bit at the front/bar while they wait and prep your table.<br><br>Between the 6 of us, we ordered the thick-cut bacon, Caesar Salad for starters, Steak for Four, Rib Steak for mains, and the potatoes and creamed Spinach for sides. Also on the table is a gravy boat of their house sauce, which to me tasted similar to a cocktail sauce, definitely has some horseradish.<br><br>Going straight to the steaks, ordering-wise there really wasn&#39;t much options so don&#39;t worry about not being able to decide what to get. The huge steaks come out on a hot plate, cut up already. The waiter will serve you a couple slices, and drizzle over some oil/butter from the plate. Good crust, strong beef taste. The tenderloin part of the porterhouse is very tender. A good amount of fat, and just a hint of aged flavour.<br><br>Service is not necessarily the best, but the waiters can be chatty if you engage with them. Price was about ~$70-80 per person with a drink.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="4" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=Iu_X4aMAsTwxsEgx9pLYBw" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=vEoQJwntiwLSBcLyaigRFg>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/vEoQJwntiwLSBcLyaigRFg/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Rib Steak" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Rib Steak" class="photo-box-img" height="348" src="https://s3-media2.fl.yelpcdn.com/bphoto/vEoQJwntiwLSBcLyaigRFg/348s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/vEoQJwntiwLSBcLyaigRFg/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Rib Steak</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=Iu_X4aMAsTwxsEgx9pLYBw&amp;select=vEoQJwntiwLSBcLyaigRFg">
            <span class="offscreen">Rib Steak</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=d6570AUiss3oHfztDK-toQ>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/d6570AUiss3oHfztDK-toQ/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Two" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Two" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/d6570AUiss3oHfztDK-toQ/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/d6570AUiss3oHfztDK-toQ/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/d6570AUiss3oHfztDK-toQ/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/d6570AUiss3oHfztDK-toQ/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/d6570AUiss3oHfztDK-toQ/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/d6570AUiss3oHfztDK-toQ/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Steak for Two</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=Iu_X4aMAsTwxsEgx9pLYBw&amp;select=d6570AUiss3oHfztDK-toQ">
            <span class="offscreen">Steak for Two</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=5loLYjHYs3s5OCXzHuCOAQ>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/5loLYjHYs3s5OCXzHuCOAQ/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Thick cut Bacon (2 orders)" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Thick cut Bacon (2 orders)" class="photo-box-img" height="168" src="https://s3-media2.fl.yelpcdn.com/bphoto/5loLYjHYs3s5OCXzHuCOAQ/168s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/5loLYjHYs3s5OCXzHuCOAQ/258s.jpg 1.54x,https://s3-media2.fl.yelpcdn.com/bphoto/5loLYjHYs3s5OCXzHuCOAQ/180s.jpg 1.07x,https://s3-media2.fl.yelpcdn.com/bphoto/5loLYjHYs3s5OCXzHuCOAQ/300s.jpg 1.79x,https://s3-media2.fl.yelpcdn.com/bphoto/5loLYjHYs3s5OCXzHuCOAQ/348s.jpg 2.07x,https://s3-media2.fl.yelpcdn.com/bphoto/5loLYjHYs3s5OCXzHuCOAQ/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Thick cut Bacon (2 orders)</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=Iu_X4aMAsTwxsEgx9pLYBw&amp;select=5loLYjHYs3s5OCXzHuCOAQ">
            <span class="offscreen">Thick cut Bacon (2 orders)</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=XbWP6vAli3waJ0iR_lcjOg" class="js-content-expander">
                        See all photos from Nathan A. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="Iu_X4aMAsTwxsEgx9pLYBw" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Millie S. and 3 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">3</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">1</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="e5S2X5Og3sbMMyGFgci6qw" data-signup-object="user_id:KE8v6mZkoC6C-4E0jnN5Aw">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="ed4CkTyTxVLZKfmsEcgfrg">
                <a href="/user_details?userid=KE8v6mZkoC6C-4E0jnN5Aw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Celia N." class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/photo/r_YZ2lZFI7Usu1RvNjkTyw/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/photo/r_YZ2lZFI7Usu1RvNjkTyw/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/photo/r_YZ2lZFI7Usu1RvNjkTyw/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/photo/r_YZ2lZFI7Usu1RvNjkTyw/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/photo/r_YZ2lZFI7Usu1RvNjkTyw/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/photo/r_YZ2lZFI7Usu1RvNjkTyw/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=KE8v6mZkoC6C-4E0jnN5Aw" data-hovercard-id="ed4CkTyTxVLZKfmsEcgfrg" data-analytics-label="about_me" id="dropdown_user-name">Celia N.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Nashville, TN</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>129</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>79</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>407</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/e5S2X5Og3sbMMyGFgci6qw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/e5S2X5Og3sbMMyGFgci6qw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Celia N.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="KE8v6mZkoC6C-4E0jnN5Aw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Celia N.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="KE8v6mZkoC6C-4E0jnN5Aw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Celia N.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        6/24/2019
    </span>

    </div>

        
                <p lang="en">Im still having trouble finding the words to express how happy and how glad I am after my second Peter Luger experience of the year.<br><br>I get reservation for 4 to celebrate my fathers birthday, and enjoying our meeting after 2 years. I asked if we could have a nice table by the window in one of the salons downstairs giving the reason why we would have that special meal, and the nailed it. <br><br>We enjoyed an amazing meal with the best service. We started with some of their delicious breads and butter (it&#39;s hard to stop eating! They are completely addictive!) with the tender and juicy bacon and their signature sauce. And following that such an amazing bite, we had the steak for 3 and a green salad.<br><br>What can I say besides the meat was absorbed perfect, I enjoyed so much better the second time. It&#39;s so delicious that even if you&#39;re not a meat fan you will enjoy it. You can taste the level of quality in the meat, that comes without any salt. You can add it or having it with their sauce, I enjoyed it just with a tiny bit of salt to embrace the flavors.<br><br>We didn&#39;t finish hungry at all, but we needed to have the cheesecake at least, and I ordered also the pecan pie a la mode because either my parents or my sister ever had one before. <br><br>So the service arrived, singing louder happy birthday, bringing the pie with the candle, and giving him a huge surprise (because we are not used to do the same in Spain).<br><br>We had a memorable meal, not just because of the delicious food, the place is just magical.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="5" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=e5S2X5Og3sbMMyGFgci6qw" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=NK4_uuEb1MY6Bd46H28vtQ>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/NK4_uuEb1MY6Bd46H28vtQ/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Three" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Three" class="photo-box-img" height="348" src="https://s3-media2.fl.yelpcdn.com/bphoto/NK4_uuEb1MY6Bd46H28vtQ/348s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/NK4_uuEb1MY6Bd46H28vtQ/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Steak for Three</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=e5S2X5Og3sbMMyGFgci6qw&amp;select=NK4_uuEb1MY6Bd46H28vtQ">
            <span class="offscreen">Steak for Three</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=McFiFPIh1AXIkU5H8dEDEQ>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/McFiFPIh1AXIkU5H8dEDEQ/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Sizzling Bacon, Extra Thick by the Slice" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Sizzling Bacon, Extra Thick by the Slice" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/McFiFPIh1AXIkU5H8dEDEQ/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/McFiFPIh1AXIkU5H8dEDEQ/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/McFiFPIh1AXIkU5H8dEDEQ/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/McFiFPIh1AXIkU5H8dEDEQ/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/McFiFPIh1AXIkU5H8dEDEQ/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/McFiFPIh1AXIkU5H8dEDEQ/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Luger&#39;s Sizzling Bacon, Extra Thick by the Slice</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=e5S2X5Og3sbMMyGFgci6qw&amp;select=McFiFPIh1AXIkU5H8dEDEQ">
            <span class="offscreen">Luger&#39;s Sizzling Bacon, Extra Thick by the Slice</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=4m_3aS0bUzoJ-qPOxEhYfQ>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/4m_3aS0bUzoJ-qPOxEhYfQ/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Three" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Three" class="photo-box-img" height="168" src="https://s3-media2.fl.yelpcdn.com/bphoto/4m_3aS0bUzoJ-qPOxEhYfQ/168s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/4m_3aS0bUzoJ-qPOxEhYfQ/258s.jpg 1.54x,https://s3-media2.fl.yelpcdn.com/bphoto/4m_3aS0bUzoJ-qPOxEhYfQ/180s.jpg 1.07x,https://s3-media2.fl.yelpcdn.com/bphoto/4m_3aS0bUzoJ-qPOxEhYfQ/300s.jpg 1.79x,https://s3-media2.fl.yelpcdn.com/bphoto/4m_3aS0bUzoJ-qPOxEhYfQ/348s.jpg 2.07x,https://s3-media2.fl.yelpcdn.com/bphoto/4m_3aS0bUzoJ-qPOxEhYfQ/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Steak for Three</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=e5S2X5Og3sbMMyGFgci6qw&amp;select=4m_3aS0bUzoJ-qPOxEhYfQ">
            <span class="offscreen">Steak for Three</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=KE8v6mZkoC6C-4E0jnN5Aw" class="js-content-expander">
                        See all photos from Celia N. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="e5S2X5Og3sbMMyGFgci6qw" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Alexis B. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">2</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="inedWqbvw5TqeM6HS1m71Q" data-signup-object="user_id:VU3YmnLbVRonct9uvS4okQ">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="HmJLxd2YZ7w3WjKFbVqmmA">
                <a href="/user_details?userid=VU3YmnLbVRonct9uvS4okQ" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Sumin C." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/-jS0sca__OW-atcd2dEwXQ/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/-jS0sca__OW-atcd2dEwXQ/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/-jS0sca__OW-atcd2dEwXQ/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/-jS0sca__OW-atcd2dEwXQ/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/-jS0sca__OW-atcd2dEwXQ/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/-jS0sca__OW-atcd2dEwXQ/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=VU3YmnLbVRonct9uvS4okQ" data-hovercard-id="HmJLxd2YZ7w3WjKFbVqmmA" data-analytics-label="about_me" id="dropdown_user-name">Sumin C.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>New York, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>622</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>145</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>904</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/inedWqbvw5TqeM6HS1m71Q" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/inedWqbvw5TqeM6HS1m71Q" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Sumin C.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="VU3YmnLbVRonct9uvS4okQ">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Sumin C.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="VU3YmnLbVRonct9uvS4okQ">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Sumin C.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        6/22/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">Came here with my family and I must say this restaurant did meet my expectations. With all the hype around their steak, I&#39;ve been wanting to try Peter Luger and I was pleased with the food, service, and price. <br><br>My parents and I got their iceberg Wedge salad, Steak for two, Creamed spinach, and for dessert, we got their Key lime pie with their homemade whipped cream. <br><br>The salad was decent and I especially like the bacon bits. I know they are also known for their bacon and we were going to get that as an appetizer, but I&#39;m glad I got to try some of it from the salad. <br><br>The steak was awesome -- the meat was so juicy, extremely tender, and it literally melted in my mouth. They gave us their special steak sauce, but I mostly just ate the steak by itself with some salt because the meat itself was delicious. It went well with the creamed spinach and I loved it, but my parents thought it was a bit bland. <br><br>I&#39;m also glad I got their key lime pie for dessert even though I was so full. The pie was refreshing and tasty, and the whipped cream went well with it even though I usually don&#39;t like whipped cream. <br><br>Overall, it was definitely one of the best steaks I&#39;ve ever had and I think it does live up to its reputation.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="2" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=inedWqbvw5TqeM6HS1m71Q" data-starting-index="0">
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=JAZ8o323_yhzbCUOlLW6mg>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/JAZ8o323_yhzbCUOlLW6mg/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Creamed Spinach (for 2)" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Creamed Spinach (for 2)" class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/JAZ8o323_yhzbCUOlLW6mg/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/JAZ8o323_yhzbCUOlLW6mg/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/JAZ8o323_yhzbCUOlLW6mg/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/JAZ8o323_yhzbCUOlLW6mg/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/JAZ8o323_yhzbCUOlLW6mg/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/JAZ8o323_yhzbCUOlLW6mg/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Creamed Spinach (for 2)</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=inedWqbvw5TqeM6HS1m71Q&amp;select=JAZ8o323_yhzbCUOlLW6mg">
            <span class="offscreen">Creamed Spinach (for 2)</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=oe40-dpii1-gAXL9UtnMjA>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/oe40-dpii1-gAXL9UtnMjA/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Two" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Two" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/oe40-dpii1-gAXL9UtnMjA/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/oe40-dpii1-gAXL9UtnMjA/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/oe40-dpii1-gAXL9UtnMjA/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/oe40-dpii1-gAXL9UtnMjA/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/oe40-dpii1-gAXL9UtnMjA/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/oe40-dpii1-gAXL9UtnMjA/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Steak for Two</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=inedWqbvw5TqeM6HS1m71Q&amp;select=oe40-dpii1-gAXL9UtnMjA">
            <span class="offscreen">Steak for Two</span>
    </a>

    </div>

                </li>

        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="inedWqbvw5TqeM6HS1m71Q" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Kim T.</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">1</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="jbfIOj9PtfOdbE08xqwCng" data-signup-object="user_id:N9m-ToOPrr_PIpMZJOUr5g">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="Xi94R3p7LMmUoK4uMqSjog">
                <a href="/user_details?userid=N9m-ToOPrr_PIpMZJOUr5g" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Michelle S." class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/photo/BlyfK8-Nh1QY5StYDBP-bg/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/photo/BlyfK8-Nh1QY5StYDBP-bg/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/photo/BlyfK8-Nh1QY5StYDBP-bg/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/photo/BlyfK8-Nh1QY5StYDBP-bg/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/photo/BlyfK8-Nh1QY5StYDBP-bg/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/photo/BlyfK8-Nh1QY5StYDBP-bg/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=N9m-ToOPrr_PIpMZJOUr5g" data-hovercard-id="Xi94R3p7LMmUoK4uMqSjog" data-analytics-label="about_me" id="dropdown_user-name">Michelle S.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Brooklyn, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>664</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>326</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>666</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/jbfIOj9PtfOdbE08xqwCng" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/jbfIOj9PtfOdbE08xqwCng" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Michelle S.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="N9m-ToOPrr_PIpMZJOUr5g">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Michelle S.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="N9m-ToOPrr_PIpMZJOUr5g">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Michelle S.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        7/14/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">48 hours later after eating at Peter Luger&#39;s I finally woke up from the food coma that ensued after feasting. I was warned that the food is so good it&#39;ll knock you out and that was true! I made reservations a month in advance and secured a time of 4:45pm on a Saturday. It wasn&#39;t busy and we were able to be seated when we arrived at 4:15pm. Our server was Danny and he was a very attentive server. We ordered a pellegrino (sparkling water helps digestion), bacon (2), porterhouse for 2, spinach and the potatoes. I have to say, the bacon was the best I&#39;ve ever had. I&#39;ve had bacon from all parts of the spectrum but this bacon brought me to heaven (and clogged my arteries). When the steak came, it was still sizzling in a pool of buttery goodness. My server basted the steak and then served it to me which was definitely a nice touch. The appetizers, spinach and potatoes were definitely up there as well. I&#39;ve never had creamed spinach like the one at Peter Luger&#39;s, it definitely pairs well with the steak. For the steak itself... I suggest rare because it is still cooking as it&#39;s brought to the table. I got the medium rare which was still tasty. Don&#39;t forget to drench the signature steak sauce onto it! Our bill ended up to be -$172 before tip. I&#39;ll definitely return after I take a few months to unclog my arteries. I can&#39;t wait to do it all over again.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="4" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=jbfIOj9PtfOdbE08xqwCng" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=d2kPYQ2ovVRFdveNvDwB8Q>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/d2kPYQ2ovVRFdveNvDwB8Q/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. The menu" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. The menu" class="photo-box-img" height="348" src="https://s3-media4.fl.yelpcdn.com/bphoto/d2kPYQ2ovVRFdveNvDwB8Q/348s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/d2kPYQ2ovVRFdveNvDwB8Q/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">The menu</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=jbfIOj9PtfOdbE08xqwCng&amp;select=d2kPYQ2ovVRFdveNvDwB8Q">
            <span class="offscreen">The menu</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=I9RmL1S2HtMkjjY41i-2lw>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/I9RmL1S2HtMkjjY41i-2lw/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Porterhouse Steak for Two" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Porterhouse Steak for Two" class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/I9RmL1S2HtMkjjY41i-2lw/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/I9RmL1S2HtMkjjY41i-2lw/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/I9RmL1S2HtMkjjY41i-2lw/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/I9RmL1S2HtMkjjY41i-2lw/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/I9RmL1S2HtMkjjY41i-2lw/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/I9RmL1S2HtMkjjY41i-2lw/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Porterhouse Steak for Two</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=jbfIOj9PtfOdbE08xqwCng&amp;select=I9RmL1S2HtMkjjY41i-2lw">
            <span class="offscreen">Porterhouse Steak for Two</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=K0ggKQI-eoAQgDJaZznH8g>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/K0ggKQI-eoAQgDJaZznH8g/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Sizzling Bacon, Extra Thick by the Slice" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Sizzling Bacon, Extra Thick by the Slice" class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/K0ggKQI-eoAQgDJaZznH8g/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/K0ggKQI-eoAQgDJaZznH8g/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/K0ggKQI-eoAQgDJaZznH8g/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/K0ggKQI-eoAQgDJaZznH8g/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/K0ggKQI-eoAQgDJaZznH8g/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/K0ggKQI-eoAQgDJaZznH8g/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Luger&#39;s Sizzling Bacon, Extra Thick by the Slice</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=jbfIOj9PtfOdbE08xqwCng&amp;select=K0ggKQI-eoAQgDJaZznH8g">
            <span class="offscreen">Luger&#39;s Sizzling Bacon, Extra Thick by the Slice</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=N9m-ToOPrr_PIpMZJOUr5g" class="js-content-expander">
                        See all photos from Michelle S. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="jbfIOj9PtfOdbE08xqwCng" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Shabo P. and 4 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">5</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">2</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="VoMzsCqn_RgldFT6UzmQqQ" data-signup-object="user_id:Pjy5Cq8lsZjpN9hm5leBJA">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="1W4x4Az_lSCHmzDBvIHSZg">
                <a href="/user_details?userid=Pjy5Cq8lsZjpN9hm5leBJA" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Mark S." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/xQ8sxjZ7bKa03ZiMdYV1Bg/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/xQ8sxjZ7bKa03ZiMdYV1Bg/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/xQ8sxjZ7bKa03ZiMdYV1Bg/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/xQ8sxjZ7bKa03ZiMdYV1Bg/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/xQ8sxjZ7bKa03ZiMdYV1Bg/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/xQ8sxjZ7bKa03ZiMdYV1Bg/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=Pjy5Cq8lsZjpN9hm5leBJA" data-hovercard-id="1W4x4Az_lSCHmzDBvIHSZg" data-analytics-label="about_me" id="dropdown_user-name">Mark S.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Cary, NC</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>174</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>295</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>491</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/VoMzsCqn_RgldFT6UzmQqQ" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/VoMzsCqn_RgldFT6UzmQqQ" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Mark S.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="Pjy5Cq8lsZjpN9hm5leBJA">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Mark S.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="Pjy5Cq8lsZjpN9hm5leBJA">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Mark S.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        6/4/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">When you visit NYC, you have to check out their most famous steakhouse. We made reservation 1 month ahead and I was super excited to try it out.<br><br>Dinner was at 7:45pm. When we arrived, the place was packed. After getting my name called, we get seated and a very nice gentleman gave us the menu. We knew exactly what we wanted: Porterhouse for 3 with creamed spinach and fried potatoes. <br><br>The bread came first but pro tip: take the bread home because you&#39;re here for the steak. The steak was absolutely mind blowing. We ordered it med rare and it was cooked perfectly. It was very juicy and full of flavor. The creamed spinach was also very tasty.<br><br>The service was absolutely top notch. I will recommend this place to anyone looking for a great steak house experience.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="2" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=VoMzsCqn_RgldFT6UzmQqQ" data-starting-index="0">
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=JBv1kNmonOt4CmJmRKS6sw>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/JBv1kNmonOt4CmJmRKS6sw/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Thick cut bacon as an appetizer" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Thick cut bacon as an appetizer" class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/JBv1kNmonOt4CmJmRKS6sw/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/JBv1kNmonOt4CmJmRKS6sw/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/JBv1kNmonOt4CmJmRKS6sw/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/JBv1kNmonOt4CmJmRKS6sw/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/JBv1kNmonOt4CmJmRKS6sw/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/JBv1kNmonOt4CmJmRKS6sw/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Thick cut bacon as an appetizer</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=VoMzsCqn_RgldFT6UzmQqQ&amp;select=JBv1kNmonOt4CmJmRKS6sw">
            <span class="offscreen">Thick cut bacon as an appetizer</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=rcY81jcp4DY9N5dZo28PFg>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/rcY81jcp4DY9N5dZo28PFg/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Porterhouse for 3. Probably the best steak I&#39;ve had." class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Porterhouse for 3. Probably the best steak I&#39;ve had." class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/rcY81jcp4DY9N5dZo28PFg/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/rcY81jcp4DY9N5dZo28PFg/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/rcY81jcp4DY9N5dZo28PFg/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/rcY81jcp4DY9N5dZo28PFg/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/rcY81jcp4DY9N5dZo28PFg/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/rcY81jcp4DY9N5dZo28PFg/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Porterhouse for 3. Probably the best steak I&#39;ve had.</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=VoMzsCqn_RgldFT6UzmQqQ&amp;select=rcY81jcp4DY9N5dZo28PFg">
            <span class="offscreen">Porterhouse for 3. Probably the best steak I&#39;ve had.</span>
    </a>

    </div>

                </li>

        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="VoMzsCqn_RgldFT6UzmQqQ" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Nini J. and 3 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">2</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="ObisS9pEr99Zbo9lGkydjA" data-signup-object="user_id:mDQhgesXXOKaDQW2o00sMg">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="4fhV-7gD1xndOcAfYoK_Rw">
                <a href="/user_details?userid=mDQhgesXXOKaDQW2o00sMg" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Ron W." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/Xx9fEW7zV-__nCoKmYRFVA/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/Xx9fEW7zV-__nCoKmYRFVA/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/Xx9fEW7zV-__nCoKmYRFVA/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/Xx9fEW7zV-__nCoKmYRFVA/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/Xx9fEW7zV-__nCoKmYRFVA/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/Xx9fEW7zV-__nCoKmYRFVA/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=mDQhgesXXOKaDQW2o00sMg" data-hovercard-id="4fhV-7gD1xndOcAfYoK_Rw" data-analytics-label="about_me" id="dropdown_user-name">Ron W.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Santa Ana, CA</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>235</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>362</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>2398</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/ObisS9pEr99Zbo9lGkydjA" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/ObisS9pEr99Zbo9lGkydjA" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Ron W.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="mDQhgesXXOKaDQW2o00sMg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Ron W.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="mDQhgesXXOKaDQW2o00sMg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Ron W.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/17/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">They just signed on with Resy which allows them to ignore the 4,000 calls per day they get (which they are highly skilled at). After wrestling with Resy and losing the match I decided Resy is only good for raising your blood pressure. <br><br>I started calling direct...for months. Yes, I got a ring but couldn&#39;t get in the ring. <br><br>Their system is truly democratic. They can/will abuse everyone. They don&#39;t care how many Twitter followers you have or who you influence, or that you have staff to handle nasty things like this. It&#39;s the same steep climb to the summit. A table.<br>  <br>Close to yelling P.L. F. U. and throwing the phone in the pool in frustration I tried one last time and got Michael the Manager. What a Prince. &#34;You&#39;re all set&#34;, he assured me after a pleasant chat. (By the time I got there months later they had me down for a table twice the same night and also for three consecutive nights). That was a small glitch and no matter, I was in the system, in the club, and finally...in the door. <br><br>I walked in after a 3 mile, one hour Uber from Manhattan and it was bedlam. 300 people were packed 4 deep at an old fashioned bar. (Probably the same one they had when they opened in 1837). Even though I had a reservation that just got me on the list of waiters waiting to be waited on. <br><br>An old fashioned bar tender made me an old fashioned Dirty Martini (stirred, not shaken. Take that, Mr. Bond) and it was easily the finest Martini I&#39;ve ever enjoyed. The serious pour had me staring at the floor...from close up. <br><br>P.L. is not pretentious. Don&#39;t bother draggin&#39; your Judith Lieber bag, or strappin&#39; on the Patek Phillipe. No one there cares. Just bring a fat bank roll because they only accept cash and that they do care about. AMEX Reward points sacrificed, it&#39;s still worth it for the Bovine bounty you&#39;re about to receive.<br><br>Be advised there are rules here, and they are strictly enforced. (A) You will suffer to get in. No walk-in&#39;s allowed. (B) You will pay cash (and tip generously). (C) You will have to wait even if you&#39;re early for your reservation. (D) You will be grilled (like a steak) to make sure your entire party is present before you&#39;re seated, (even if you&#39;re dining alone). <br><br>So, why, you ask, am I here? Because, (A) It&#39;s Peter Luger. (B) It will exceed all the hype. (C) You&#39;re gonna have the best hunk-a-meat you&#39;ve ever had or will have. (Tinder notwithstanding.) <br><br>The Bar: Basics. No chemistry set, periodic table concoctions. Whatever you&#39;re handed will have no peer.<br><br>The Vibe: Bucket-listers, wannabe A-listers, real A-listers, Wise Guys, and those who have chosen wisely.  <br><br>The Staff: Jerek was a Big Brother, a Dutch Uncle, Father Confessor, BFF, Master of Meat...a great guy and a top tier true pro. <br><br>The Food: You can memorize the menu in one glance. <br><br>Have the thick slice of house cured Bacon even if you&#39;re a devout Orthodox Kosher, Halal, Vegan. Once you eat it there can be no turning back. Close your eyes and jump, I say. <br><br>The Caesar Salad (despite the boxed croutons, for which I forgive you, P.L.) was exceptional. Creamy and Garlic-y it was so good, in fact, the ghost of Caesar Cardini shows up periodically to get his fix.<br><br>The Porterhouse for 1, 2, 3, 4, or the whole 182nd Airborne is why you came and it will be dead perfect. Period. <br><br>The steak arrives pre-cut on a platter with a splatter of Jus as large as an Olympic pool. (No, the steak was not swimming in it, it was just luxuriating in it.)<br><br>The meat comes direct from a 1,200 degree grill/oven and is so hot, your face will tan. Remember you&#39;re enrolled in steak Master Class. This is the way a Master serves it. <br><br>Your server will double spoon the first few slices onto your plate, and  anoint it with Jus, (so good you&#39;ll be tempted to ask for a Go Cup of it.) You have just been served food made by God&#39;s own Chef.  <br><br>Add to that the necessary/required Creamed Spinach (deliciously more Spinach than Cream...Lawry&#39;s eat your heart out). Luger&#39;s version of this is divine. <br><br>The Onion Rings got a flour dredge and then a flash fry. Oh my, were they fabulous? Yes, they were.  <br><br>No room for dessert? No normal human would have available space for it. And, yes you ate the entire Porterhouse, and then gnawed on the bone like a Hyenna in public. <br><br>No matter that you can&#39;t manage dessert under any circumstances, you&#39;re gonna have the Apple Strudel under an Everest of Schlag anyway.You&#39;re gonna eat it, and you&#39;re gonna like it. I loved it. <br><br>Walking out late required wading through 300 people 4 deep at the bar anxiously waiting for their name to be called and there&#39;s a reason for that.  <br><br>Yes, there&#39;s a reason this place is off the chart and unique. But, you don&#39;t need a reason to go. Just go. Go now!</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="10" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=ObisS9pEr99Zbo9lGkydjA" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=jonzGJ2zcwIg2v7lqob0eQ>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/jonzGJ2zcwIg2v7lqob0eQ/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Vegan&#39;s sinful desire." class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Vegan&#39;s sinful desire." class="photo-box-img" height="348" src="https://s3-media2.fl.yelpcdn.com/bphoto/jonzGJ2zcwIg2v7lqob0eQ/348s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/jonzGJ2zcwIg2v7lqob0eQ/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Vegan&#39;s sinful desire.</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=ObisS9pEr99Zbo9lGkydjA&amp;select=jonzGJ2zcwIg2v7lqob0eQ">
            <span class="offscreen">Vegan&#39;s sinful desire.</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=svofe4j2PtWPA-sRfceTaQ>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/svofe4j2PtWPA-sRfceTaQ/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Whisky Sour to whisk you away," class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Whisky Sour to whisk you away," class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/svofe4j2PtWPA-sRfceTaQ/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/svofe4j2PtWPA-sRfceTaQ/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/svofe4j2PtWPA-sRfceTaQ/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/svofe4j2PtWPA-sRfceTaQ/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/svofe4j2PtWPA-sRfceTaQ/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/svofe4j2PtWPA-sRfceTaQ/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Whisky Sour to whisk you away,</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=ObisS9pEr99Zbo9lGkydjA&amp;select=svofe4j2PtWPA-sRfceTaQ">
            <span class="offscreen">Whisky Sour to whisk you away,</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=Af-vTK16M8imo1DywGumCA>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/Af-vTK16M8imo1DywGumCA/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. The Bar stirred me." class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. The Bar stirred me." class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/Af-vTK16M8imo1DywGumCA/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/Af-vTK16M8imo1DywGumCA/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/Af-vTK16M8imo1DywGumCA/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/Af-vTK16M8imo1DywGumCA/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/Af-vTK16M8imo1DywGumCA/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/Af-vTK16M8imo1DywGumCA/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">The Bar stirred me.</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=ObisS9pEr99Zbo9lGkydjA&amp;select=Af-vTK16M8imo1DywGumCA">
            <span class="offscreen">The Bar stirred me.</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=mDQhgesXXOKaDQW2o00sMg" class="js-content-expander">
                        See all photos from Ron W. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="ObisS9pEr99Zbo9lGkydjA" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Nini J. and 25 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">13</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">17</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">9</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="fzcfxGA4vtWYhXzRNmVwnw" data-signup-object="user_id:8_L5NTHp6L5skX-Hz3eH3Q">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="Qr3KRyxCY9KvmsxxxPdAFw">
                <a href="/user_details?userid=8_L5NTHp6L5skX-Hz3eH3Q" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Rashad A." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/_cMxsSWJhw8-Uh6zQmKPfw/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/_cMxsSWJhw8-Uh6zQmKPfw/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/_cMxsSWJhw8-Uh6zQmKPfw/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/_cMxsSWJhw8-Uh6zQmKPfw/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/_cMxsSWJhw8-Uh6zQmKPfw/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/_cMxsSWJhw8-Uh6zQmKPfw/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=8_L5NTHp6L5skX-Hz3eH3Q" data-hovercard-id="Qr3KRyxCY9KvmsxxxPdAFw" data-analytics-label="about_me" id="dropdown_user-name">Rashad A.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Los Angeles, CA</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>501</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>108</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>243</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/fzcfxGA4vtWYhXzRNmVwnw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/fzcfxGA4vtWYhXzRNmVwnw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Rashad A.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="8_L5NTHp6L5skX-Hz3eH3Q">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Rashad A.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="8_L5NTHp6L5skX-Hz3eH3Q">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Rashad A.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-4 rating-large" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/30/2019
    </span>

    </div>

        
                <p lang="en">Burger review: very good, tasty, patty preparation is biased towards rare (we asked for medium rare) which I personally don&#39;t mind, but figured this is good to know. You have the option to add cheese to the burger- totally up to you. I added cheese and enjoyed it but will definitely have the burger without cheese on the next run so I can better appreciate the beef. Quality bun, not too soft and has form. Slightly on the firmer side. <br><br>Tried an order of the thick cut bacon- nice fat/meat ratio. Really melts in your mouth, but flavor is slightly lacking. <br><br>Came on a Thursday during lunch with a friend, we had to wait about 20 minutes. Coming on a Friday or weekend expect a longer wait. Cash only.<br><br>Overall a solid burger and will return for another.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="2" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=fzcfxGA4vtWYhXzRNmVwnw" data-starting-index="0">
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=2XENFpTbtb524F2kN98toA>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/2XENFpTbtb524F2kN98toA/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/2XENFpTbtb524F2kN98toA/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/2XENFpTbtb524F2kN98toA/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/2XENFpTbtb524F2kN98toA/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/2XENFpTbtb524F2kN98toA/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/2XENFpTbtb524F2kN98toA/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/2XENFpTbtb524F2kN98toA/ls.jpg 1.49x" width="168">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=fzcfxGA4vtWYhXzRNmVwnw&amp;select=2XENFpTbtb524F2kN98toA">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=ZYzgYj42aXR_pvV5lgcOXA>
        
                <img data-async-src=https://s3-media1.fl.yelpcdn.com/bphoto/ZYzgYj42aXR_pvV5lgcOXA/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="168" src="https://s3-media1.fl.yelpcdn.com/bphoto/ZYzgYj42aXR_pvV5lgcOXA/168s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/ZYzgYj42aXR_pvV5lgcOXA/258s.jpg 1.54x,https://s3-media1.fl.yelpcdn.com/bphoto/ZYzgYj42aXR_pvV5lgcOXA/180s.jpg 1.07x,https://s3-media1.fl.yelpcdn.com/bphoto/ZYzgYj42aXR_pvV5lgcOXA/300s.jpg 1.79x,https://s3-media1.fl.yelpcdn.com/bphoto/ZYzgYj42aXR_pvV5lgcOXA/348s.jpg 2.07x,https://s3-media1.fl.yelpcdn.com/bphoto/ZYzgYj42aXR_pvV5lgcOXA/ls.jpg 1.49x" width="168">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=fzcfxGA4vtWYhXzRNmVwnw&amp;select=ZYzgYj42aXR_pvV5lgcOXA">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>

        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="fzcfxGA4vtWYhXzRNmVwnw" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Nini J. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="6i4Ul4evnsudCKjy8nvcRw" data-signup-object="user_id:yC-lvSrksRBMW8aRXOYTAA">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="yvjcS4GtOdSe7qv2JyDx8g">
                <a href="/user_details?userid=yC-lvSrksRBMW8aRXOYTAA" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Vince L." class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/photo/-_qg73otQL_G4mIQmAUuUg/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/photo/-_qg73otQL_G4mIQmAUuUg/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/photo/-_qg73otQL_G4mIQmAUuUg/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/photo/-_qg73otQL_G4mIQmAUuUg/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/photo/-_qg73otQL_G4mIQmAUuUg/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/photo/-_qg73otQL_G4mIQmAUuUg/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=yC-lvSrksRBMW8aRXOYTAA" data-hovercard-id="yvjcS4GtOdSe7qv2JyDx8g" data-analytics-label="about_me" id="dropdown_user-name">Vince L.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Toronto, Canada</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>3</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>81</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>399</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/6i4Ul4evnsudCKjy8nvcRw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/6i4Ul4evnsudCKjy8nvcRw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Vince L.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="yC-lvSrksRBMW8aRXOYTAA">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Vince L.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="yC-lvSrksRBMW8aRXOYTAA">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Vince L.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/25/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">I came to NYC for this. It&#39;s been a fee years since we&#39;ve last had it and in my mind it&#39;s the best steak ever. <br><br>This is definitely at the top of my list for best Steakhouses ever. The perfect cut, aged the right time and the finishing touch. Its full of flavour and juices. Pair it with some houSr wine and you&#39;ve got the combo ready. And dont forget to finish the bone, don&#39;t be shy. Wish I could have this every month. And dont forget to pair it with some creamed spinach and a milkshake if you still have room. <br><br>Service is pretty good in general. Everytime we&#39;ve been the server&#39;s have been quite the jokers. This place is cash or debit only so keep that in mind. And make sure to have a reservation no matter the time . Otherwise you would be waiting in line with the tourists for a good 2 hours. It&#39;ll be a painful wait</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="8" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=6i4Ul4evnsudCKjy8nvcRw" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=zTLdMKiiQ8Jsk5xD6CqBQw>
        
                <img data-async-src=https://s3-media1.fl.yelpcdn.com/bphoto/zTLdMKiiQ8Jsk5xD6CqBQw/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="348" src="https://s3-media1.fl.yelpcdn.com/bphoto/zTLdMKiiQ8Jsk5xD6CqBQw/348s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/zTLdMKiiQ8Jsk5xD6CqBQw/1000s.jpg 2.87x" width="348">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=6i4Ul4evnsudCKjy8nvcRw&amp;select=zTLdMKiiQ8Jsk5xD6CqBQw">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=GGa683jfAscWt7G-yA_Xrg>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/GGa683jfAscWt7G-yA_Xrg/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="168" src="https://s3-media2.fl.yelpcdn.com/bphoto/GGa683jfAscWt7G-yA_Xrg/168s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/GGa683jfAscWt7G-yA_Xrg/258s.jpg 1.54x,https://s3-media2.fl.yelpcdn.com/bphoto/GGa683jfAscWt7G-yA_Xrg/180s.jpg 1.07x,https://s3-media2.fl.yelpcdn.com/bphoto/GGa683jfAscWt7G-yA_Xrg/300s.jpg 1.79x,https://s3-media2.fl.yelpcdn.com/bphoto/GGa683jfAscWt7G-yA_Xrg/348s.jpg 2.07x,https://s3-media2.fl.yelpcdn.com/bphoto/GGa683jfAscWt7G-yA_Xrg/ls.jpg 1.49x" width="168">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=6i4Ul4evnsudCKjy8nvcRw&amp;select=GGa683jfAscWt7G-yA_Xrg">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=GJDjphLPLE9OsPYkXrgcWw>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/GJDjphLPLE9OsPYkXrgcWw/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/GJDjphLPLE9OsPYkXrgcWw/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/GJDjphLPLE9OsPYkXrgcWw/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/GJDjphLPLE9OsPYkXrgcWw/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/GJDjphLPLE9OsPYkXrgcWw/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/GJDjphLPLE9OsPYkXrgcWw/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/GJDjphLPLE9OsPYkXrgcWw/ls.jpg 1.49x" width="168">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=6i4Ul4evnsudCKjy8nvcRw&amp;select=GJDjphLPLE9OsPYkXrgcWw">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=yC-lvSrksRBMW8aRXOYTAA" class="js-content-expander">
                        See all photos from Vince L. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="6i4Ul4evnsudCKjy8nvcRw" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Jun-Jun V. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">1</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="8vBRdx7HDxVybveVahinUg" data-signup-object="user_id:jFweUfeEqM4y-njcQN-Npg">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="VneU3GUXPZqUFePdgQWKPQ">
                <a href="/user_details?userid=jFweUfeEqM4y-njcQN-Npg" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Lesley Y." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/WK7S9Gow1yeTDGhsmdJdvA/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/WK7S9Gow1yeTDGhsmdJdvA/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/WK7S9Gow1yeTDGhsmdJdvA/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/WK7S9Gow1yeTDGhsmdJdvA/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/WK7S9Gow1yeTDGhsmdJdvA/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/WK7S9Gow1yeTDGhsmdJdvA/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=jFweUfeEqM4y-njcQN-Npg" data-hovercard-id="VneU3GUXPZqUFePdgQWKPQ" data-analytics-label="about_me" id="dropdown_user-name">Lesley Y.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>San Francisco Bay Area, CA</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>410</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>174</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>442</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/8vBRdx7HDxVybveVahinUg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/8vBRdx7HDxVybveVahinUg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Lesley Y.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="jFweUfeEqM4y-njcQN-Npg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Lesley Y.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="jFweUfeEqM4y-njcQN-Npg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Lesley Y.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/16/2019
    </span>

    </div>

        
                <p lang="en">We ventured one hour to get here. It&#39;s a classic must-go steakhouse of Brooklyn / New York. When we got there, there was still a wait despite reservations. <br><br>Make sure you get a reservation! <br><br>We were then seated and got their prime beef steak for two. Then also got some creamed spinach.<br><br>At this point of the trip we had already eaten a lot, so the fact that I could still taste how delicious the steak is (juicy, tender, flavourful) shows how great it is.<br><br>Bill came out to be $120.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="1" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=8vBRdx7HDxVybveVahinUg" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=rSA5qtd0N40Nm-QyCKSM8g>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/rSA5qtd0N40Nm-QyCKSM8g/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak" class="photo-box-img" height="348" src="https://s3-media4.fl.yelpcdn.com/bphoto/rSA5qtd0N40Nm-QyCKSM8g/348s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/rSA5qtd0N40Nm-QyCKSM8g/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Steak</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=8vBRdx7HDxVybveVahinUg&amp;select=rSA5qtd0N40Nm-QyCKSM8g">
            <span class="offscreen">Steak</span>
    </a>

    </div>

                </li>

        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="8vBRdx7HDxVybveVahinUg" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Adam C. and 3 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">2</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="CgE6lZ1FcbJpXvt8camoSw" data-signup-object="user_id:rMgKJHpA0kyU8fUk-Ef0NQ">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="y758x13KXYUXLNiCXPW--A">
                <a href="/user_details?userid=rMgKJHpA0kyU8fUk-Ef0NQ" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Dona F." class="photo-box-img" height="60" src="https://s3-media2.fl.yelpcdn.com/photo/GpukMVM-JBa_dgHzalaSSg/60s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/photo/GpukMVM-JBa_dgHzalaSSg/90s.jpg 1.50x,https://s3-media2.fl.yelpcdn.com/photo/GpukMVM-JBa_dgHzalaSSg/168s.jpg 2.80x,https://s3-media2.fl.yelpcdn.com/photo/GpukMVM-JBa_dgHzalaSSg/ms.jpg 1.67x,https://s3-media2.fl.yelpcdn.com/photo/GpukMVM-JBa_dgHzalaSSg/180s.jpg 3.00x,https://s3-media2.fl.yelpcdn.com/photo/GpukMVM-JBa_dgHzalaSSg/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=rMgKJHpA0kyU8fUk-Ef0NQ" data-hovercard-id="y758x13KXYUXLNiCXPW--A" data-analytics-label="about_me" id="dropdown_user-name">Dona F.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Staten Island, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>31</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>247</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>770</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/CgE6lZ1FcbJpXvt8camoSw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/CgE6lZ1FcbJpXvt8camoSw" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Dona F.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="rMgKJHpA0kyU8fUk-Ef0NQ">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Dona F.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="rMgKJHpA0kyU8fUk-Ef0NQ">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Dona F.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/8/2019
    </span>

    </div>

        
                <p lang="en">I was last here 20 years ago but I&#39;m worth it so it was time to revisit!!<br>No more sawdust. <br>Need reservations weeks in advance <br>Amazing servers who have been there for ages and know their meat!!<br>Just awesome steak cooked exactly how you ask. <br>Side of creamer spinach - delicious <br>German potatoes - cooked with butter then baked to perfect crunch. <br>Apple strudel was good but the home made Schlag is even better. I can eat that plain</p> 
            
    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="CgE6lZ1FcbJpXvt8camoSw" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Claudine V. and 4 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">4</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">3</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">3</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="020u0HPv9jWqgeoOTjFZ8w" data-signup-object="user_id:4vc1PIObqy8kxsX7jyV2dw">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="3w4dYmOlrHq47ZRg_XR-Lg">
                <a href="/user_details?userid=4vc1PIObqy8kxsX7jyV2dw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Melissa G." class="photo-box-img" height="60" src="https://s3-media2.fl.yelpcdn.com/photo/LgWr964Fd09-amUJWh5JQA/60s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/photo/LgWr964Fd09-amUJWh5JQA/90s.jpg 1.50x,https://s3-media2.fl.yelpcdn.com/photo/LgWr964Fd09-amUJWh5JQA/168s.jpg 2.80x,https://s3-media2.fl.yelpcdn.com/photo/LgWr964Fd09-amUJWh5JQA/ms.jpg 1.67x,https://s3-media2.fl.yelpcdn.com/photo/LgWr964Fd09-amUJWh5JQA/180s.jpg 3.00x,https://s3-media2.fl.yelpcdn.com/photo/LgWr964Fd09-amUJWh5JQA/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=4vc1PIObqy8kxsX7jyV2dw" data-hovercard-id="3w4dYmOlrHq47ZRg_XR-Lg" data-analytics-label="about_me" id="dropdown_user-name">Melissa G.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>New York, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>26</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>213</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>125</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/020u0HPv9jWqgeoOTjFZ8w" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/020u0HPv9jWqgeoOTjFZ8w" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Melissa G.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="4vc1PIObqy8kxsX7jyV2dw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Melissa G.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="4vc1PIObqy8kxsX7jyV2dw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Melissa G.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-4 rating-large" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/12/2019
    </span>

    </div>

        
                <p lang="en">Dinner at Peter Luger exceeded my expectations. It&#39;s quite casual for a place with a Michelin star; dressing up is unnecessary, yet you won&#39;t be out of place if you want to make an occasion of it. I had heard rumors of cold service but our experience was quite the opposite; our waiter was patient and helpful -- he suggested pairing Malbec wine with our steak and it was a wonderful choice.<br><br>Although I usually prefer warm bread, the rolls that they brought to our table were delicious. I especially liked the roll filled with onion, and found it hard to pace myself knowing all the food to come! We started off with the thick-cut bacon and a caesar salad. I&#39;m not even a huge bacon person but I will say that it&#39;s something you should try here. Tastes great with the steak sauce! The caesar salad was fantastic and dressed perfectly -- probably one of my favorite parts of the meal actually. The steak though was truly amazing. So incredibly tender, definitely among the best steaks I&#39;ve had. Portion size was quite impressive! We also got the baked potato and onion rings as sides, both of which I would pass on next time. The potato was pretty standard (nothing you couldn&#39;t do at home), and the onion rings were thin and stringy. <br><br>For dessert, we got the pecan pie with the homemade schlag of course! The pie was good although not super memorable, but we did enjoy the mound of whipped cream! To make you feel a little better about the sum you spend on this dinner, they send you away with some chocolate coins -- a nice touch. An experience very worth having.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="5" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=020u0HPv9jWqgeoOTjFZ8w" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=oQzIjfLs2c_u-YbzgYsy7Q>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/oQzIjfLs2c_u-YbzgYsy7Q/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. pecan pie + schlag" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. pecan pie + schlag" class="photo-box-img" height="348" src="https://s3-media4.fl.yelpcdn.com/bphoto/oQzIjfLs2c_u-YbzgYsy7Q/348s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/oQzIjfLs2c_u-YbzgYsy7Q/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">pecan pie + schlag</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=020u0HPv9jWqgeoOTjFZ8w&amp;select=oQzIjfLs2c_u-YbzgYsy7Q">
            <span class="offscreen">pecan pie + schlag</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=HKdl0K6UT8XFdRBGNJogDg>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/HKdl0K6UT8XFdRBGNJogDg/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. steak!" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. steak!" class="photo-box-img" height="168" src="https://s3-media2.fl.yelpcdn.com/bphoto/HKdl0K6UT8XFdRBGNJogDg/168s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/HKdl0K6UT8XFdRBGNJogDg/258s.jpg 1.54x,https://s3-media2.fl.yelpcdn.com/bphoto/HKdl0K6UT8XFdRBGNJogDg/180s.jpg 1.07x,https://s3-media2.fl.yelpcdn.com/bphoto/HKdl0K6UT8XFdRBGNJogDg/300s.jpg 1.79x,https://s3-media2.fl.yelpcdn.com/bphoto/HKdl0K6UT8XFdRBGNJogDg/348s.jpg 2.07x,https://s3-media2.fl.yelpcdn.com/bphoto/HKdl0K6UT8XFdRBGNJogDg/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">steak!</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=020u0HPv9jWqgeoOTjFZ8w&amp;select=HKdl0K6UT8XFdRBGNJogDg">
            <span class="offscreen">steak!</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=MvXs5X_-3H8TJaq8nbTwNQ>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/MvXs5X_-3H8TJaq8nbTwNQ/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. filled with onion -- scrumptious" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. filled with onion -- scrumptious" class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/MvXs5X_-3H8TJaq8nbTwNQ/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/MvXs5X_-3H8TJaq8nbTwNQ/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/MvXs5X_-3H8TJaq8nbTwNQ/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/MvXs5X_-3H8TJaq8nbTwNQ/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/MvXs5X_-3H8TJaq8nbTwNQ/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/MvXs5X_-3H8TJaq8nbTwNQ/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">filled with onion -- scrumptious</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=020u0HPv9jWqgeoOTjFZ8w&amp;select=MvXs5X_-3H8TJaq8nbTwNQ">
            <span class="offscreen">filled with onion -- scrumptious</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=4vc1PIObqy8kxsX7jyV2dw" class="js-content-expander">
                        See all photos from Melissa G. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="020u0HPv9jWqgeoOTjFZ8w" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">hans simon f. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="5wURUiNYfdVKin_j2C70cg" data-signup-object="user_id:_jqTa-3-fqmBxnkANJHtFw">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="RfvNs7gOVK3icFzFX70ghQ">
                <a href="/user_details?userid=_jqTa-3-fqmBxnkANJHtFw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Brian F." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/N1XsgzOZPfFsxtwXEC0rFg/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/N1XsgzOZPfFsxtwXEC0rFg/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/N1XsgzOZPfFsxtwXEC0rFg/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/N1XsgzOZPfFsxtwXEC0rFg/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/N1XsgzOZPfFsxtwXEC0rFg/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/N1XsgzOZPfFsxtwXEC0rFg/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=_jqTa-3-fqmBxnkANJHtFw" data-hovercard-id="RfvNs7gOVK3icFzFX70ghQ" data-analytics-label="about_me" id="dropdown_user-name">Brian F.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Bellevue, WA</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>502</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>617</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>2119</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/5wURUiNYfdVKin_j2C70cg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/5wURUiNYfdVKin_j2C70cg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Brian F.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="_jqTa-3-fqmBxnkANJHtFw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Brian F.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="_jqTa-3-fqmBxnkANJHtFw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Brian F.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        4/29/2019
    </span>

    </div>

        
                <p lang="en">Cash only<br><br>Have you ever eaten a piece of steak and thought to yourself, &#34;this is the best steak I&#39;ve ever had in my life&#34;? Well that&#39;s what happened to me at Peter Luger. This place is generally seen as one of the best steak houses in New York and I am glad I was about to treat myself here for my birthday.<br><br>I&#39;d highly recommend making reservations. You can make them online or over the phone. I actually arrived half an hour before my reserved time and got seated within a couple minutes.<br><br>I ordered the sizzling bacon ($6.95) which was smoky and fatty and delicious. It was so easy to cut through. I also got the single steak ($54.95) which was heavenly! Seriously, don&#39;t get me started on the steak... it was amazing and cooked perfectly medium rare. The crispiness and the slight char on the outside and then the juiciness and tenderness on the inside. I&#39;m salivating just thinking about it. It comes with dipping sauce which is sweet and tastes like it&#39;s ketchup based. Both of these were just melt in your mouth and I recommend getting these. I&#39;ve also heard great things about the burger but sadly they only have it for lunch so I wasn&#39;t able to try it.<br><br>Note: I&#39;ve heard of the troubles of making a reservation here. You can actually make reservations on their website now! I still recommend you make one in advance; I made mine three weeks in advanced or so and there were no problems. If you&#39;d prefer to speak to someone you can call them. If you call, I recommend calling them right when they open as I only needed to wait a couple minutes until someone answered.<br><br>Is this the best steak I&#39;ve ever had? Yup.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="5" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=5wURUiNYfdVKin_j2C70cg" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=fm_YjUsUe-5oE6m1thE49w>
        
                <img data-async-src=https://s3-media1.fl.yelpcdn.com/bphoto/fm_YjUsUe-5oE6m1thE49w/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Menu (April 2019)" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Menu (April 2019)" class="photo-box-img" height="348" src="https://s3-media1.fl.yelpcdn.com/bphoto/fm_YjUsUe-5oE6m1thE49w/348s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/fm_YjUsUe-5oE6m1thE49w/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Menu (April 2019)</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=5wURUiNYfdVKin_j2C70cg&amp;select=fm_YjUsUe-5oE6m1thE49w">
            <span class="offscreen">Menu (April 2019)</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=Qbbb-uCbZlWIRK9hV4jGYw>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/Qbbb-uCbZlWIRK9hV4jGYw/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Complimentary bread" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Complimentary bread" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/Qbbb-uCbZlWIRK9hV4jGYw/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/Qbbb-uCbZlWIRK9hV4jGYw/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/Qbbb-uCbZlWIRK9hV4jGYw/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/Qbbb-uCbZlWIRK9hV4jGYw/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/Qbbb-uCbZlWIRK9hV4jGYw/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/Qbbb-uCbZlWIRK9hV4jGYw/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Complimentary bread</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=5wURUiNYfdVKin_j2C70cg&amp;select=Qbbb-uCbZlWIRK9hV4jGYw">
            <span class="offscreen">Complimentary bread</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=IEvmrJN9G8JCFXf7czWBew>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/IEvmrJN9G8JCFXf7czWBew/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Sizzling bacon ($6.95)" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Sizzling bacon ($6.95)" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/IEvmrJN9G8JCFXf7czWBew/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/IEvmrJN9G8JCFXf7czWBew/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/IEvmrJN9G8JCFXf7czWBew/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/IEvmrJN9G8JCFXf7czWBew/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/IEvmrJN9G8JCFXf7czWBew/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/IEvmrJN9G8JCFXf7czWBew/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Sizzling bacon ($6.95)</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=5wURUiNYfdVKin_j2C70cg&amp;select=IEvmrJN9G8JCFXf7czWBew">
            <span class="offscreen">Sizzling bacon ($6.95)</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=_jqTa-3-fqmBxnkANJHtFw" class="js-content-expander">
                        See all photos from Brian F. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="5wURUiNYfdVKin_j2C70cg" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Vincent L. and 9 others</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">5</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">3</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">9</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="4sQGtQc24_kQzscu2_D81g" data-signup-object="user_id:btdMwcbzwNPbsCcI-cLl4g">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="HipkAWLhI1Zzgs1gFGWg4Q">
                <a href="/user_details?userid=btdMwcbzwNPbsCcI-cLl4g" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Phoebe W." class="photo-box-img" height="60" src="https://s3-media2.fl.yelpcdn.com/photo/y64EW8JeCuwoDCgKEX2hsQ/60s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/photo/y64EW8JeCuwoDCgKEX2hsQ/90s.jpg 1.50x,https://s3-media2.fl.yelpcdn.com/photo/y64EW8JeCuwoDCgKEX2hsQ/168s.jpg 2.80x,https://s3-media2.fl.yelpcdn.com/photo/y64EW8JeCuwoDCgKEX2hsQ/ms.jpg 1.67x,https://s3-media2.fl.yelpcdn.com/photo/y64EW8JeCuwoDCgKEX2hsQ/180s.jpg 3.00x,https://s3-media2.fl.yelpcdn.com/photo/y64EW8JeCuwoDCgKEX2hsQ/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=btdMwcbzwNPbsCcI-cLl4g" data-hovercard-id="HipkAWLhI1Zzgs1gFGWg4Q" data-analytics-label="about_me" id="dropdown_user-name">Phoebe W.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>New York, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>609</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>423</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>1023</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/4sQGtQc24_kQzscu2_D81g" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/4sQGtQc24_kQzscu2_D81g" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Phoebe W.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="btdMwcbzwNPbsCcI-cLl4g">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Phoebe W.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="btdMwcbzwNPbsCcI-cLl4g">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Phoebe W.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/29/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">Came here with my family for a post-graduation meal! It was 100% worth the trek to Brooklyn and the 8:45 PM reservation time (make your reservations well in advance, because this was the earliest I could get 2 months in advance for a Wednesday evening!).<br><br>The four of us ordered the steak for two, the creamed spinach, the bacon, and the onion rings. I can&#39;t even describe just how good everything was -- that bacon was unlike anything else I&#39;d ever put into my mouth. The steak too -- nothing is as good as the first bite of steak, and paired with the Peter Luger&#39;s sauce, was a dish I will never forget. Creamed spinach -- I thought it was good, my family thought it was amazing. Even my sister and dad who hate spinach ate it by the spoonful. The onion rings were fresh and hot and tasty as well, a little bit on the greasy side but that&#39;s what makes them taste so good the next day.<br><br>Our server was so funny and attentive, he made us feel welcomed from the start. My dad left his backpack at the restaurant on accident and they went above and beyond in helping him get it back the following day. Definitely a great splurge-worthy meal and experience!</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="4" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=4sQGtQc24_kQzscu2_D81g" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=cZyPFCaH28vFSCYyLk6Kwg>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/cZyPFCaH28vFSCYyLk6Kwg/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="348" src="https://s3-media3.fl.yelpcdn.com/bphoto/cZyPFCaH28vFSCYyLk6Kwg/348s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/cZyPFCaH28vFSCYyLk6Kwg/1000s.jpg 2.87x" width="348">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=4sQGtQc24_kQzscu2_D81g&amp;select=cZyPFCaH28vFSCYyLk6Kwg">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=gLMm_9MnFID3SkCSEs1dww>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/gLMm_9MnFID3SkCSEs1dww/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Two, Onion Rings (for 2)" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Steak for Two, Onion Rings (for 2)" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/gLMm_9MnFID3SkCSEs1dww/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/gLMm_9MnFID3SkCSEs1dww/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/gLMm_9MnFID3SkCSEs1dww/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/gLMm_9MnFID3SkCSEs1dww/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/gLMm_9MnFID3SkCSEs1dww/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/gLMm_9MnFID3SkCSEs1dww/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Steak for Two, Onion Rings (for 2)</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=4sQGtQc24_kQzscu2_D81g&amp;select=gLMm_9MnFID3SkCSEs1dww">
            <span class="offscreen">Steak for Two, Onion Rings (for 2)</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=b_qynTPan5w5bWXMiXWERg>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/b_qynTPan5w5bWXMiXWERg/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Menu" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Menu" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/b_qynTPan5w5bWXMiXWERg/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/b_qynTPan5w5bWXMiXWERg/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/b_qynTPan5w5bWXMiXWERg/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/b_qynTPan5w5bWXMiXWERg/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/b_qynTPan5w5bWXMiXWERg/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/b_qynTPan5w5bWXMiXWERg/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Menu</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=4sQGtQc24_kQzscu2_D81g&amp;select=b_qynTPan5w5bWXMiXWERg">
            <span class="offscreen">Menu</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=btdMwcbzwNPbsCcI-cLl4g" class="js-content-expander">
                        See all photos from Phoebe W. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="4sQGtQc24_kQzscu2_D81g" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
            Was this review &hellip;?
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="CmDZZg3jrd2HVxDn7ZVvvg" data-signup-object="user_id:1xZDeCio4x3917dv1sX4sg">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="uDbveDJMD3vg2InuMXJhuA">
                <a href="/user_details?userid=1xZDeCio4x3917dv1sX4sg" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Jae H." class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/photo/xJJC1lq6bT2SpNlCb9jxVQ/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/photo/xJJC1lq6bT2SpNlCb9jxVQ/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/photo/xJJC1lq6bT2SpNlCb9jxVQ/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/photo/xJJC1lq6bT2SpNlCb9jxVQ/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/photo/xJJC1lq6bT2SpNlCb9jxVQ/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/photo/xJJC1lq6bT2SpNlCb9jxVQ/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=1xZDeCio4x3917dv1sX4sg" data-hovercard-id="uDbveDJMD3vg2InuMXJhuA" data-analytics-label="about_me" id="dropdown_user-name">Jae H.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>New York, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>37</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>316</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>1440</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/CmDZZg3jrd2HVxDn7ZVvvg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/CmDZZg3jrd2HVxDn7ZVvvg" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Jae H.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="1xZDeCio4x3917dv1sX4sg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Jae H.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="1xZDeCio4x3917dv1sX4sg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Jae H.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-4 rating-large" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/9/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">I made reservation for my friends whom visited in NY from LA. One of my friend birthday coming up and I decided to took my friends at this place. I&#39;ve been tried this place but I could tell you that their prime steak was the best I ever had in this place before. To honestly I like this place but My experience wasn&#39;t good in here. <br><br>Stopped by for dinner with friends. We ordered 3 of prime steak and side of creamy spinach, mixed green salad. 3 of cosmopolitan cocktails, one of bartender special cocktail and key lime pie with coffee, vanilla ice cream for dessert.<br><br>Prime steak(4.5/5) was comes out hot also tasted great than I expected. Very juicy and tender, soft. Peter Luger steak sauce is famous. I loved their sauce and my all of friends were looks like they love steak with sauce. <br><br>Creamy spinach(2/5)wasn&#39;t creamy than I expected. I&#39;ve had eat better elsewhere. Not creamy much but savory. If you like less creamy style, you would like it.<br><br>Mixed green salad(3/5) was ok but very fresh vegetables. Nothing special, you can eat this salad at your home too. <br><br>Key lime pie(3/5)was good but I&#39;ve eat better elsewhere. I am a huge fan of dessert also I love key lime pie. I eaten that in my whole life but this place key lime pie wasn&#39;t the best. Home made whipped cream was the best. Not too sweet but kinda of savory with very soft. Even if you don&#39;t like sweets, you can try key lime pie and their home made whipped cream. But I would recommend you that if you want to eat key lime pie, drink coffee. <br><br>Vanilla ice cream (2/5) just ok. <br><br>Cosmopolitan(2/5) were really strong cocktails. Actually it isn&#39;t strong cocktails but the bartender made our cocktails very strong. I had to request them to extra cranberry juice. <br><br>Service(3/5)-our waiter was accommodating and very friendly. They load up our plate with steak and side. But there&#39;s busy and it was difficult to get for what we need. Very little of space between table to table. One waiter was keeping touch my chair or moved my chair by accidentally but nothing apologized. I understand how they&#39;re busy and not enough space for served but i personally thought he could say at least sorry or excuse me. Am I too picky? <br><br>Ambience(2.5/5)- old restaurant and their interior is wood style. It&#39;s simple interior but I couldn&#39;t see any of special point of interior. <br><br>This place is cash only. If you have plan for go this place then have some cash. No dress cord here.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="10" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=CmDZZg3jrd2HVxDn7ZVvvg" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=z2uliNbMp7Dx7pW2mamU1g>
        
                <img data-async-src=https://s3-media1.fl.yelpcdn.com/bphoto/z2uliNbMp7Dx7pW2mamU1g/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Dessert menu" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Dessert menu" class="photo-box-img" height="348" src="https://s3-media1.fl.yelpcdn.com/bphoto/z2uliNbMp7Dx7pW2mamU1g/348s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/z2uliNbMp7Dx7pW2mamU1g/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Dessert menu</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=CmDZZg3jrd2HVxDn7ZVvvg&amp;select=z2uliNbMp7Dx7pW2mamU1g">
            <span class="offscreen">Dessert menu</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=ejA9Bplxq_YTxJpuPDbtiw>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/ejA9Bplxq_YTxJpuPDbtiw/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Prime steak for 3" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Prime steak for 3" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/ejA9Bplxq_YTxJpuPDbtiw/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/ejA9Bplxq_YTxJpuPDbtiw/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/ejA9Bplxq_YTxJpuPDbtiw/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/ejA9Bplxq_YTxJpuPDbtiw/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/ejA9Bplxq_YTxJpuPDbtiw/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/ejA9Bplxq_YTxJpuPDbtiw/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Prime steak for 3</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=CmDZZg3jrd2HVxDn7ZVvvg&amp;select=ejA9Bplxq_YTxJpuPDbtiw">
            <span class="offscreen">Prime steak for 3</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=XxZRmNC08ZiauVRJSPV02A>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/XxZRmNC08ZiauVRJSPV02A/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/XxZRmNC08ZiauVRJSPV02A/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/XxZRmNC08ZiauVRJSPV02A/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/XxZRmNC08ZiauVRJSPV02A/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/XxZRmNC08ZiauVRJSPV02A/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/XxZRmNC08ZiauVRJSPV02A/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/XxZRmNC08ZiauVRJSPV02A/ls.jpg 1.49x" width="168">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=CmDZZg3jrd2HVxDn7ZVvvg&amp;select=XxZRmNC08ZiauVRJSPV02A">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=1xZDeCio4x3917dv1sX4sg" class="js-content-expander">
                        See all photos from Jae H. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="CmDZZg3jrd2HVxDn7ZVvvg" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Sama C. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">2</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count">1</span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="8Dlo0HKBTG0KYmktLTNx4w" data-signup-object="user_id:lc6KhKhuCk6_ldm9gAviKg">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="uCiu2Xc-oipHsEymwB3Tyw">
                <a href="/user_details?userid=lc6KhKhuCk6_ldm9gAviKg" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Danny L." class="photo-box-img" height="60" src="https://s3-media4.fl.yelpcdn.com/photo/NoHwy-Az-rIe8q-WaQbG3Q/60s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/photo/NoHwy-Az-rIe8q-WaQbG3Q/90s.jpg 1.50x,https://s3-media4.fl.yelpcdn.com/photo/NoHwy-Az-rIe8q-WaQbG3Q/168s.jpg 2.80x,https://s3-media4.fl.yelpcdn.com/photo/NoHwy-Az-rIe8q-WaQbG3Q/ms.jpg 1.67x,https://s3-media4.fl.yelpcdn.com/photo/NoHwy-Az-rIe8q-WaQbG3Q/180s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/photo/NoHwy-Az-rIe8q-WaQbG3Q/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=lc6KhKhuCk6_ldm9gAviKg" data-hovercard-id="uCiu2Xc-oipHsEymwB3Tyw" data-analytics-label="about_me" id="dropdown_user-name">Danny L.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>New York, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>105</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>75</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>286</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/8Dlo0HKBTG0KYmktLTNx4w" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/8Dlo0HKBTG0KYmktLTNx4w" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Danny L.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="lc6KhKhuCk6_ldm9gAviKg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Danny L.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="lc6KhKhuCk6_ldm9gAviKg">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Danny L.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-5 rating-large" title="5.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="5.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/14/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">Where to start....I have been live in nyc for almost a decade and this is the FIRST TIME I came to Peter Luger, I have to say I absolutely love it !!! I know this steakhouse is very famous, but I didn&#39;t expect the fact that I have to make a reservation one week ahead. However it all worth it at the end! <br><br>It was PACKED when we get here and people already waiting in the line. The interior design is old fashioned, classic steakhouse style which I love. <br><br>We ordered their famous bacon strips, porterhouse steak for two, wedge salad, cream Spanish , French fries, and pecan pie with their homemade whip cream. Everything tasted sooo delicious. Especially the thick cut bacon strip with their house steak Sauce.. it was like in heaven... hahaha.  The steak was cooked beautifully , tender and juicy, the sound of the sizzling pan make it extra mouth watering... My friend is not a dessert lover, but she was to able to finish the whole pecan pie herself...  <br><br>Food price is very reasonable too, we (two people)spent around $300 + tip. Which is very normal for a good steak house in nyc.<br><br>Highly recommended !!!!</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="5" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=8Dlo0HKBTG0KYmktLTNx4w" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=J9e9A1uAstnOcMmdXFs9mw>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/J9e9A1uAstnOcMmdXFs9mw/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="348" src="https://s3-media3.fl.yelpcdn.com/bphoto/J9e9A1uAstnOcMmdXFs9mw/348s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/J9e9A1uAstnOcMmdXFs9mw/1000s.jpg 2.87x" width="348">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=8Dlo0HKBTG0KYmktLTNx4w&amp;select=J9e9A1uAstnOcMmdXFs9mw">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=wUPOUTDW3dgsjvQwKdtYVg>
        
                <img data-async-src=https://s3-media4.fl.yelpcdn.com/bphoto/wUPOUTDW3dgsjvQwKdtYVg/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="168" src="https://s3-media4.fl.yelpcdn.com/bphoto/wUPOUTDW3dgsjvQwKdtYVg/168s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/wUPOUTDW3dgsjvQwKdtYVg/258s.jpg 1.54x,https://s3-media4.fl.yelpcdn.com/bphoto/wUPOUTDW3dgsjvQwKdtYVg/180s.jpg 1.07x,https://s3-media4.fl.yelpcdn.com/bphoto/wUPOUTDW3dgsjvQwKdtYVg/300s.jpg 1.79x,https://s3-media4.fl.yelpcdn.com/bphoto/wUPOUTDW3dgsjvQwKdtYVg/348s.jpg 2.07x,https://s3-media4.fl.yelpcdn.com/bphoto/wUPOUTDW3dgsjvQwKdtYVg/ls.jpg 1.49x" width="168">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=8Dlo0HKBTG0KYmktLTNx4w&amp;select=wUPOUTDW3dgsjvQwKdtYVg">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=LH8kSUsZT-2AFlHy1tZZAg>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/LH8kSUsZT-2AFlHy1tZZAg/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States" class="photo-box-img" height="168" src="https://s3-media3.fl.yelpcdn.com/bphoto/LH8kSUsZT-2AFlHy1tZZAg/168s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/LH8kSUsZT-2AFlHy1tZZAg/258s.jpg 1.54x,https://s3-media3.fl.yelpcdn.com/bphoto/LH8kSUsZT-2AFlHy1tZZAg/180s.jpg 1.07x,https://s3-media3.fl.yelpcdn.com/bphoto/LH8kSUsZT-2AFlHy1tZZAg/300s.jpg 1.79x,https://s3-media3.fl.yelpcdn.com/bphoto/LH8kSUsZT-2AFlHy1tZZAg/348s.jpg 2.07x,https://s3-media3.fl.yelpcdn.com/bphoto/LH8kSUsZT-2AFlHy1tZZAg/ls.jpg 1.49x" width="168">

            </noscript>



        

            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=8Dlo0HKBTG0KYmktLTNx4w&amp;select=LH8kSUsZT-2AFlHy1tZZAg">
            <span class="offscreen"></span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=lc6KhKhuCk6_ldm9gAviKg" class="js-content-expander">
                        See all photos from Danny L. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="8Dlo0HKBTG0KYmktLTNx4w" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Joyce H. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="qTTmNquLIKqQbRAT5B4EEA" data-signup-object="user_id:HVQKFDi-vZbawtmLMq3ZOw">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="3OL1nZ8E8RtKvOPquMAyJw">
                <a href="/user_details?userid=HVQKFDi-vZbawtmLMq3ZOw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Gab G." class="photo-box-img" height="60" src="https://s3-media4.fl.yelpcdn.com/photo/uellgqIek0sU1pDDbWZ77Q/60s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/photo/uellgqIek0sU1pDDbWZ77Q/90s.jpg 1.50x,https://s3-media4.fl.yelpcdn.com/photo/uellgqIek0sU1pDDbWZ77Q/168s.jpg 2.80x,https://s3-media4.fl.yelpcdn.com/photo/uellgqIek0sU1pDDbWZ77Q/ms.jpg 1.67x,https://s3-media4.fl.yelpcdn.com/photo/uellgqIek0sU1pDDbWZ77Q/180s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/photo/uellgqIek0sU1pDDbWZ77Q/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=HVQKFDi-vZbawtmLMq3ZOw" data-hovercard-id="3OL1nZ8E8RtKvOPquMAyJw" data-analytics-label="about_me" id="dropdown_user-name">Gab G.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>New York, NY</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>211</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>1720</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>5278</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/qTTmNquLIKqQbRAT5B4EEA" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/qTTmNquLIKqQbRAT5B4EEA" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Gab G.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="HVQKFDi-vZbawtmLMq3ZOw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Gab G.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="HVQKFDi-vZbawtmLMq3ZOw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Gab G.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-3 rating-large" title="3.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="3.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/12/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">Made a reservation over the phone for my boyfriend&#39;s birthday about a month in advance. Service was great, but unfortunately we found the food to be hit or miss. Rib eye, Caesar salad, creamed spinach and apple strudel were all great, but the German potatoes were dry and the key lime pie had a weird chemical taste to it. On top of that, the dishes we enjoyed were just that - enjoyable. At a place like this, I want to be blown away when I take a bite. That wasn&#39;t really the case here, and so I think I am better off returning to places like 4 Charles and Del Frisco&#39;s to get my meat and potatoes fix. This is a classic place that I&#39;m glad I tried once, but I&#39;m not sure I&#39;d recommend it if you&#39;re looking for a singular NYC steak experience. <br><br>3.5 stars.</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="5" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=qTTmNquLIKqQbRAT5B4EEA" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=zl7TorOSyxmJ96-G9Ae60w>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/zl7TorOSyxmJ96-G9Ae60w/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Rib eye" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Rib eye" class="photo-box-img" height="348" src="https://s3-media2.fl.yelpcdn.com/bphoto/zl7TorOSyxmJ96-G9Ae60w/348s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/zl7TorOSyxmJ96-G9Ae60w/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Rib eye</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=qTTmNquLIKqQbRAT5B4EEA&amp;select=zl7TorOSyxmJ96-G9Ae60w">
            <span class="offscreen">Rib eye</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=23JF2LMMI0OI8e7u_-Qrng>
        
                <img data-async-src=https://s3-media2.fl.yelpcdn.com/bphoto/23JF2LMMI0OI8e7u_-Qrng/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Special German Fried Potatoes" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Luger&#39;s Special German Fried Potatoes" class="photo-box-img" height="168" src="https://s3-media2.fl.yelpcdn.com/bphoto/23JF2LMMI0OI8e7u_-Qrng/168s.jpg" srcset="https://s3-media2.fl.yelpcdn.com/bphoto/23JF2LMMI0OI8e7u_-Qrng/258s.jpg 1.54x,https://s3-media2.fl.yelpcdn.com/bphoto/23JF2LMMI0OI8e7u_-Qrng/180s.jpg 1.07x,https://s3-media2.fl.yelpcdn.com/bphoto/23JF2LMMI0OI8e7u_-Qrng/300s.jpg 1.79x,https://s3-media2.fl.yelpcdn.com/bphoto/23JF2LMMI0OI8e7u_-Qrng/348s.jpg 2.07x,https://s3-media2.fl.yelpcdn.com/bphoto/23JF2LMMI0OI8e7u_-Qrng/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Luger&#39;s Special German Fried Potatoes</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=qTTmNquLIKqQbRAT5B4EEA&amp;select=23JF2LMMI0OI8e7u_-Qrng">
            <span class="offscreen">Luger&#39;s Special German Fried Potatoes</span>
    </a>

    </div>

                </li>
                <li style="width: 168px; height: 168px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=iqdhTingeD5Qoy8rmqD2TA>
        
                <img data-async-src=https://s3-media1.fl.yelpcdn.com/bphoto/iqdhTingeD5Qoy8rmqD2TA/168s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Apple Strudel and Key Lime Pie" class="photo-box-img no-js-hidden" height="168" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="168">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Apple Strudel and Key Lime Pie" class="photo-box-img" height="168" src="https://s3-media1.fl.yelpcdn.com/bphoto/iqdhTingeD5Qoy8rmqD2TA/168s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/iqdhTingeD5Qoy8rmqD2TA/258s.jpg 1.54x,https://s3-media1.fl.yelpcdn.com/bphoto/iqdhTingeD5Qoy8rmqD2TA/180s.jpg 1.07x,https://s3-media1.fl.yelpcdn.com/bphoto/iqdhTingeD5Qoy8rmqD2TA/300s.jpg 1.79x,https://s3-media1.fl.yelpcdn.com/bphoto/iqdhTingeD5Qoy8rmqD2TA/348s.jpg 2.07x,https://s3-media1.fl.yelpcdn.com/bphoto/iqdhTingeD5Qoy8rmqD2TA/ls.jpg 1.49x" width="168">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Apple Strudel and Key Lime Pie</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=qTTmNquLIKqQbRAT5B4EEA&amp;select=iqdhTingeD5Qoy8rmqD2TA">
            <span class="offscreen">Apple Strudel and Key Lime Pie</span>
    </a>

    </div>

                </li>

                <li class="more-review-photos">
                    <a href="/biz_photos/peter-luger-brooklyn-2?userid=HVQKFDi-vZbawtmLMq3ZOw" class="js-content-expander">
                        See all photos from Gab G. for Peter Luger
                    </a>
                </li>
        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="qTTmNquLIKqQbRAT5B4EEA" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Matthew B. and 1 other</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>
        <li>
                <div class="review review--with-sidebar" data-review-id="DEr0Cu2Bs4ho932v7EWG5Q" data-signup-object="user_id:5yGgNgkkNvsv6dYiA_whFw">
            <div class="review-sidebar">
        <div class="review-sidebar-content">
                <div class="ypassport media-block">
        <div class="media-avatar responsive-photo-box">
                        <div class="photo-box pb-60s" data-hovercard-id="jJsc46A9uPxxss-dgUlBSg">
                <a href="/user_details?userid=5yGgNgkkNvsv6dYiA_whFw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Yonnie C." class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/photo/g5sRIcgRA0ID_hC2GwB11w/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/photo/g5sRIcgRA0ID_hC2GwB11w/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/photo/g5sRIcgRA0ID_hC2GwB11w/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/photo/g5sRIcgRA0ID_hC2GwB11w/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/photo/g5sRIcgRA0ID_hC2GwB11w/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/photo/g5sRIcgRA0ID_hC2GwB11w/120s.jpg 2.00x" width="60">

        </a>

    </div>



        </div>
        <div class="media-story">
                <ul class="user-passport-info">
        <li class="user-name">
                    <a class="user-display-name js-analytics-click" href="/user_details?userid=5yGgNgkkNvsv6dYiA_whFw" data-hovercard-id="jJsc46A9uPxxss-dgUlBSg" data-analytics-label="about_me" id="dropdown_user-name">Yonnie C.</a>
        </li>
        <li class="user-location responsive-hidden-small">
            <b>Los Angeles, CA</b>
        </li>
    </ul>

            

    <ul class="user-passport-stats">
        <li class="friend-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-friends icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_friends" />
    </svg>
</span>
            <b>96</b> friends
        </li>
        <li class="review-count responsive-small-display-inline-block">
            <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-review icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_review" />
    </svg>
</span>
            <b>2682</b> reviews
        </li>
            <li class="photo-count responsive-small-display-inline-block">
                <span aria-hidden="true" style="fill: #f15c00; width: 18px; height: 18px;" class="icon icon--18-camera icon--size-18">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_camera" />
    </svg>
</span>
                <b>2444</b> photos
            </li>
                <li class="is-elite responsive-small-display-inline-block">
            <a href="/elite">Elite ’19</a>
        </li>

    </ul>

        </div>
    </div>

                    <ul class="action-link-list action-link-list--small">
        
        <li>
            


    <a class="arrange arrange--middle send-to-friend" data-pop-uri="/send_to_friend/review/DEr0Cu2Bs4ho932v7EWG5Q" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-share icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_share" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Share review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle embed-review" data-pop-uri="/review_embed_modal/DEr0Cu2Bs4ho932v7EWG5Q" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-embed icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_embed" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Embed review
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-compliment" href="javascript:;" rel="">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-compliment icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_compliment" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Compliment
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle send-pm" href="javascript:;" rel="Yonnie C.">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-speech icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_speech" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Send message
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-add" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="5yGgNgkkNvsv6dYiA_whFw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-following icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_following" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Follow Yonnie C.
        </div>
    </a>

        </li>

        <li>
            


    <a class="arrange arrange--middle manage-following manage-following-remove hidden" data-csrf-token="614d59ff4ebd1b8fe3edd7248f4f310b5cd3253a58c1c85604b9bc1203d58e60" href="javascript:;" rel="5yGgNgkkNvsv6dYiA_whFw">
        <div class="action-link_icon arrange_unit">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-unfollow icon--size-18 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_unfollow" />
    </svg>
</span>
        </div>
        <div class="action-link_label arrange_unit arrange_unit--fill">
            Stop following Yonnie C.
        </div>
    </a>

        </li>


    </ul>

        </div>
    </div>

        <div class="review-wrapper">
                <div class="review-content">
            <div class="biz-rating biz-rating-large biz-rating-large--wrap clearfix">
        <div class="biz-rating__stars" >
                    



    <div class="i-stars i-stars--regular-4 rating-large" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>



        </div>
            <span class="rating-qualifier">
        5/8/2019
    </span>

    </div>

                        <ul class="review-tags">
                <li class="review-tags_item">

        <span >
            <span aria-hidden="true" style="fill: #0077bc; width: 18px; height: 18px;" class="icon icon--18-check-in icon--size-18 u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_check_in" />
    </svg>
</span>1 check-in
        </span>
    </li>

    </ul>


                <p lang="en">Ordered Steak for 2. Asked for medium but got medium rare. Because the plate was so hot, we waited for the steak on the plate to baste itself more before consuming. The porterhouse steak is on the leaner side, not as marbled as we would like, but is still good. <br><br>Recommend to make a reservation, lines for walk in gets long. No credit cards, cash or debit.<br><br>(This review written by husband)</p> 
                        
        <ul class="photo-box-grid clearfix js-content-expandable lightbox-media-parent" data-ad-logging-csrf="fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265" data-ad-logging-uri="/ad_acknowledgment" data-ga-path="media_lightbox/servlet:biz_details/type:biz" data-logging-csrf="c5f3294d518afd93de858c7ee0c61370ce179937ef89daa9e224848196ae2922" data-logging-uri="/biz_photos/4yPqqJDJOQX69gC66YUDkA/log_views" data-media-count="1" data-media-url="/biz_photos/get_media_slice/4yPqqJDJOQX69gC66YUDkA?reviewid=DEr0Cu2Bs4ho932v7EWG5Q" data-starting-index="0">
                <li style="width: 348px; height: 348px;">
                        <div class="photo-box photo-box--interactive is-loading" data-photo-id=Fls_69hsX-HJjQydn1e1sg>
        
                <img data-async-src=https://s3-media3.fl.yelpcdn.com/bphoto/Fls_69hsX-HJjQydn1e1sg/348s.jpg alt="Photo of Peter Luger - Brooklyn, NY, United States. Sizzling in butter, $109" class="photo-box-img no-js-hidden" height="348" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/978c1bee49d7/assets/img/1x1.png" width="348">

            <noscript>
                    <img alt="Photo of Peter Luger - Brooklyn, NY, United States. Sizzling in butter, $109" class="photo-box-img" height="348" src="https://s3-media3.fl.yelpcdn.com/bphoto/Fls_69hsX-HJjQydn1e1sg/348s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/Fls_69hsX-HJjQydn1e1sg/1000s.jpg 2.87x" width="348">

            </noscript>



                <div class="photo-box-overlay js-overlay">
                <div class="photo-box-overlay_caption">Sizzling in butter, $109</div>
        </div>


            <a class="biz-shim js-lightbox-media-link js-analytics-click" data-analytics-label="biz-photo"  href="/biz_photos/peter-luger-brooklyn-2?reviewid=DEr0Cu2Bs4ho932v7EWG5Q&amp;select=Fls_69hsX-HJjQydn1e1sg">
            <span class="offscreen">Sizzling in butter, $109</span>
    </a>

    </div>

                </li>

        </ul>

    </div>
        <div class="review-footer clearfix">
                    <div class="rateReview voting-feedback" data-review-id="DEr0Cu2Bs4ho932v7EWG5Q" data-origin="biz_details_review_feed">
                    <p class="voting-intro voting-prompt">
                <a class="js-review-feedback-voter-list" href="javascript:;">Shabo P.</a> voted for this review
    </p>
    <ul class="voting-buttons" data-csrf-token="b105964f9be8a1b0337190cb073b12daa48a9788f5265d89acd6386257d9b345">
            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary useful js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="useful">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-useful-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_useful_outline" />
    </svg>
</span>    <span class="vote-type">Useful</span>
    <span class="count">1</span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary funny js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="funny">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-funny-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_funny_outline" />
    </svg>
</span>    <span class="vote-type">Funny</span>
    <span class="count"></span>
    </a>
    </li>

            
    <li class="vote-item inline-block">

    <a class="ybtn ybtn--small ybtn--secondary cool js-analytics-click" data-analytics-label="ufc" href="javascript:;" rel="cool">
            <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-cool-outline icon--size-18 icon--currentColor button-content u-space-r-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_cool_outline" />
    </svg>
</span>    <span class="vote-type">Cool</span>
    <span class="count"></span>
    </a>
    </li>

    </ul>
    <div class="js-deanon-modal-container">
            <div class="js-deanon-modal u-hidden">
    <div class="modal_body">
            <div class="u-text-centered u-space-t6 u-space-b6">
                <h3 class="u-space-b3">Others will see how you vote!</h3>

                <div class="u-space-b3">
                    <img src="https://s3-media3.fl.yelpcdn.com/assets/srv0/yelp_deanonymize_ufc/f871a0ff7872/lib/img/200x120_ufc_globe.png">
                </div>

                <div class="u-space-b2" style="width:76%;margin:auto;">
                    Heads up: From now on, other Yelpers will be able to see how you voted. Want to chime in?
                </div>




    <form action="/review_feedback/deanonymize/1.0.0" class="js-deanon-form" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="591c8eca28158193261c6733acf958e23522b2173ec4b1696e1758aef9bb2590">

                    <button href="javascript:;" type="submit" value="submit" class="ybtn ybtn--primary ybtn--small js-let-me-vote-button u-space-b1"><span>Yes, Let Me Vote!</span></button>
    </form>
                <div class="js-maybe-later-button u-pseudo-link">
                    <small>Maybe Later</small>
                </div>
            </div>
    </div>
    </div>

    </div>

        </div>

    </div>

    
    

        </div>
    </div>

        </li>

    </ul>

        </div>
        <div class="review-pager">
            
        <div class="pagination-block">
            <div class="arrange arrange--stack arrange--baseline arrange--6">
                    <div class="page-of-pages arrange_unit arrange_unit--fill">
        Page 1 of 270
    </div>

                        <div class="pagination-links arrange_unit">
        <div class="arrange arrange--baseline">

                    <div class="arrange_unit page-option current">
                        <span class="pagination-links_anchor">1</span>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=20">
                            2
                        </a>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=40">
                            3
                        </a>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=60">
                            4
                        </a>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=80">
                            5
                        </a>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=100">
                            6
                        </a>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=120">
                            7
                        </a>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=140">
                            8
                        </a>
                    </div>
                    <div class="arrange_unit page-option">
                        <a class="available-number pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=160">
                            9
                        </a>
                    </div>

                <div class="arrange_unit">
                    <a class="u-decoration-none next pagination-links_anchor" href="https://www.yelp.com/biz/peter-luger-brooklyn-2?start=20">
                        <span class="pagination-label responsive-hidden-small pagination-links_anchor">Next</span>
                        <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-chevron-right icon--size-24 icon--currentColor">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_chevron_right" />
    </svg>
</span>
                    </a>
                </div>
        </div>
    </div>

            </div>
        </div>

        </div>
    </div>



                                <div class="not-recommended ysection">
                <a class="subtle-text inline-block js-expander-link" href="/not_recommended_reviews/peter-luger-brooklyn-2" rel="nofollow">
            893 other reviews that are not currently recommended
    </a>
    <div class="not-recommended-reviews-drawer filtered-reviews-content js-expander-content hidden"></div>

        </div>

                        

                    <div class="ysection js-best-of-yelp-container best-of-yelp-container eight-col"></div>
                    


                </div>
    </div>

    <div class="column column-beta sidebar">
                <div class="open-rail clearfix white-rail">









                    

                    

                        


                    


                        <h3 class="offscreen">Business info summary</h3>
        <div class="island summary">
            <ul class="iconed-list">
                                    <li class="biz-hours iconed-list-item">
        <div class="iconed-list-avatar">
                <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-clock icon--size-24 icon--error">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_clock" />
    </svg>
</span>
        </div>
        <div class="iconed-list-story">
            <dl class="short-def-list">
                <dt class="attribute-key">Today</dt>
                <dd>
                    <strong class="u-space-r-half"><span class="nowrap">1:00 pm</span> - <span class="nowrap">9:30 pm</span></strong>
                        <span class="nowrap extra closed">Closed now</span>
                </dd>
            </dl>
        </div>
    </li>



                                <li class="menu-link-block iconed-list-item">
        <div class="iconed-list-avatar">
            <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-food icon--size-24 menu-icon">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_food" />
    </svg>
</span>
        </div>
        <div class="iconed-list-story js-add-url-tagging">
            <b>    <a href="https://www.yelp.com/biz_redir?cachebuster=1563162273&amp;s=eaba562b9cf91311153e81d45ca3452ab74d246ec165d3bd1943618a543a5d70&amp;src_bizid=4yPqqJDJOQX69gC66YUDkA&amp;url=http%3A%2F%2Fpeterluger.com%2Fbrooklyn-menu%2F&amp;website_link_type=menu" target='_blank' rel='no-follow' class="external-menu js-external-menu">
        <span class="offscreen">Opens an external link</span>
        Full menu<span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-external-link icon--size-18 icon--currentColor u-space-l-half">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_external_link" />
    </svg>
</span>
    </a>
</b>
        </div>
    </li>


                        <li class="iconed-list-item">
            <div class="iconed-list-avatar">
                    
        <span class="business-attribute price-range" data-remainder="">$$$$</span>

            </div>
            <div class="iconed-list-story">
                <dl class="short-def-list">
                    <dt class="attribute-key">Price range</dt>
                    <dd class="nowrap price-description">
                        Above $61
                    </dd>
                </dl>
            </div>
        </li>

                        <li class="iconed-list-item health-score">
            <div class="iconed-list-avatar">
                <div class="score-block">
                    <span aria-hidden="true" style="width: 24px; height: 24px;" class="icon icon--24-medical icon--size-24">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_medical" />
    </svg>
</span>
                </div>
            </div>
            <div class="iconed-list-story">
                <div class="health-score-info">
                    <dl class="short-def-list">
                        <dt class="attribute-key">
                            <b><a href="/inspections/peter-luger-brooklyn-2">Health Score</a></b>
                        </dt>
                            <dd class="nowrap health-score-description">
                                A
                            </dd>
                            <dt class="nowrap u-text-subtle score-attribution">
                                Powered by HDScores
                            </dt>
                    </dl>
                </div>
            </div>
        </li>

                
            </ul>
        </div>


                    <div class="js-sidebar-promotion-container">
                    </div>

                </div>
                <div class="bordered-rail">
                        <div class="ywidget biz-hours">
        <h3 >
            Hours
        </h3>
                    <table class="table table-simple hours-table">
            <tbody>
                            <tr>
        <th scope="row">Mon</th>
        <td>
            <span class="nowrap">12:00 pm</span> - <span class="nowrap">9:30 pm</span>
        </td>
        <td class="extra">
        </td>
    </tr>

                            <tr>
        <th scope="row">Tue</th>
        <td>
            <span class="nowrap">12:00 pm</span> - <span class="nowrap">9:30 pm</span>
        </td>
        <td class="extra">
        </td>
    </tr>

                            <tr>
        <th scope="row">Wed</th>
        <td>
            <span class="nowrap">12:00 pm</span> - <span class="nowrap">9:30 pm</span>
        </td>
        <td class="extra">
        </td>
    </tr>

                            <tr>
        <th scope="row">Thu</th>
        <td>
            <span class="nowrap">12:00 pm</span> - <span class="nowrap">9:30 pm</span>
        </td>
        <td class="extra">
        </td>
    </tr>

                            <tr>
        <th scope="row">Fri</th>
        <td>
            <span class="nowrap">12:00 pm</span> - <span class="nowrap">10:30 pm</span>
        </td>
        <td class="extra">
        </td>
    </tr>

                            <tr>
        <th scope="row">Sat</th>
        <td>
            <span class="nowrap">12:00 pm</span> - <span class="nowrap">10:30 pm</span>
        </td>
        <td class="extra">
        </td>
    </tr>

                            <tr>
        <th scope="row">Sun</th>
        <td>
            <span class="nowrap">1:00 pm</span> - <span class="nowrap">9:30 pm</span>
        </td>
        <td class="extra">
                <span class="nowrap closed">Closed now</span>
        </td>
    </tr>

            </tbody>
        </table>


                    <a href="/biz_attribute?biz_id=4yPqqJDJOQX69gC66YUDkA">
        <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-pencil icon--size-14 icon--currentColor link-more edit-info u-space-r1">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_pencil" />
    </svg>
</span>Edit business info
    </a>

    </div>

                    

                            <div class="ywidget">
            <h3>More business info</h3>

            <ul class="ylist">
                        <li>
            <div class="short-def-list">
                    <dl>
                        <dt class="attribute-key">
                            Kids Activities Nearby
                        </dt>
                        <dd>
                            Yes
                        </dd>
                    </dl>
                    <dl>
                        <dt class="attribute-key">
                            High Chairs
                        </dt>
                        <dd>
                            Yes
                        </dd>
                    </dl>
                    <dl>
                        <dt class="attribute-key">
                            Keto Options
                        </dt>
                        <dd>
                            Yes
                        </dd>
                    </dl>


                        <dl>
                            <dt class="attribute-key">
                                Takes Reservations
                            </dt>
                            <dd>
                                Yes
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Delivery
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Take-out
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Accepts Credit Cards
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Accepts Apple Pay
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Accepts Google Pay
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Good For
                            </dt>
                            <dd>
                                Dinner
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Parking
                            </dt>
                            <dd>
                                Garage, Street, Private Lot
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Bike Parking
                            </dt>
                            <dd>
                                Yes
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Wheelchair Accessible
                            </dt>
                            <dd>
                                Yes
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Good for Kids
                            </dt>
                            <dd>
                                Yes
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Good for Groups
                            </dt>
                            <dd>
                                Yes
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Attire
                            </dt>
                            <dd>
                                Casual
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Ambience
                            </dt>
                            <dd>
                                Classy
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Noise Level
                            </dt>
                            <dd>
                                Average
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Alcohol
                            </dt>
                            <dd>
                                Full Bar
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Outdoor Seating
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Wi-Fi
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Has TV
                            </dt>
                            <dd>
                                Yes
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Dogs Allowed
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Waiter Service
                            </dt>
                            <dd>
                                Yes
                            </dd>
                        </dl>

                        <dl>
                            <dt class="attribute-key">
                                Caters
                            </dt>
                            <dd>
                                No
                            </dd>
                        </dl>

            </div>
        </li>

            </ul>
            
        </div>



                    
        <div class="media-block first-to-review ywidget">
            <div class="media-avatar">
                <span aria-hidden="true" style="fill: #2a9c9f; width: 24px; height: 24px;" class="icon icon--24-first icon--size-24">
    <svg role="img" class="icon_svg">
        <use xlink:href="#24x24_first" />
    </svg>
</span>
                            <div class="photo-box pb-30s" data-hovercard-id="G_rIDjEagT12ucyEg4blog">
                <a href="/user_details?userid=Gw4UiyZjER7GuyGbwSRIQw" class="js-analytics-click" data-analytics-label="user-photo">
                <img alt="Edward I." class="photo-box-img" height="30" src="https://s3-media1.fl.yelpcdn.com/photo/GJM4XLPnoRNLV_5FwmRUsA/30s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/photo/GJM4XLPnoRNLV_5FwmRUsA/60s.jpg 2.00x,https://s3-media1.fl.yelpcdn.com/photo/GJM4XLPnoRNLV_5FwmRUsA/90s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/photo/GJM4XLPnoRNLV_5FwmRUsA/ss.jpg 1.33x" width="30">

        </a>

    </div>



            </div>
            <div class="media-story">
                <div class="media-title">        <a class="user-display-name js-analytics-click" href="/user_details?userid=Gw4UiyZjER7GuyGbwSRIQw" data-hovercard-id="G_rIDjEagT12ucyEg4blog" data-analytics-label="about_me" id="dropdown_user-name">Edward I.</a></div>
                <a href="/biz/peter-luger-brooklyn-2?hrid=Nz6rROvHoBZodrtGZbtLgg" class="disableable-link label" data-review-id="Nz6rROvHoBZodrtGZbtLgg">
                    First to review
                </a>
            </div>
        </div>

                    


                    
        <div class="ywidget related-businesses js-related-businesses related-businesses--with-max-height">
            <h3>People also viewed</h3>
            <ul class="ylist">
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/keens-steakhouse-new-york?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Keens Steakhouse" class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/bphoto/Wt-lr0rX2f-lncs05SWDlA/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/Wt-lr0rX2f-lncs05SWDlA/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/bphoto/Wt-lr0rX2f-lncs05SWDlA/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/bphoto/Wt-lr0rX2f-lncs05SWDlA/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/bphoto/Wt-lr0rX2f-lncs05SWDlA/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/bphoto/Wt-lr0rX2f-lncs05SWDlA/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="eHWtYkovY5UmYe9kyHuV3w" href="/biz/keens-steakhouse-new-york?page_src=related_bizes"><span >Keens Steakhouse</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4 rating" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>


                <span class="review-count rating-qualifier">
            2406 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$$</span>
                        </span>

                        <span class="category-str-list">
                            Steakhouses, Desserts, Pubs
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Midtown West
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/del-friscos-double-eagle-steakhouse-new-york?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Del Frisco&#39;s Double Eagle Steakhouse" class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/bphoto/FyQ0V-QtndWStk5B8qwjXw/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/FyQ0V-QtndWStk5B8qwjXw/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/bphoto/FyQ0V-QtndWStk5B8qwjXw/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/bphoto/FyQ0V-QtndWStk5B8qwjXw/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/bphoto/FyQ0V-QtndWStk5B8qwjXw/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/bphoto/FyQ0V-QtndWStk5B8qwjXw/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="Oi19yd85XNCNYoT87K98AA" href="/biz/del-friscos-double-eagle-steakhouse-new-york?page_src=related_bizes"><span >Del Frisco’s Double Eagle Steakhouse</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4-half rating" title="4.5 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.5 star rating">
    </div>


                <span class="review-count rating-qualifier">
            2816 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$$</span>
                        </span>

                        <span class="category-str-list">
                            Steakhouses, American (New), Seafood
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Theater District
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/cote-new-york?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Cote" class="photo-box-img" height="60" src="https://s3-media4.fl.yelpcdn.com/bphoto/5iRixCvfWHjaHEFSpn7ZvQ/60s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/5iRixCvfWHjaHEFSpn7ZvQ/90s.jpg 1.50x,https://s3-media4.fl.yelpcdn.com/bphoto/5iRixCvfWHjaHEFSpn7ZvQ/168s.jpg 2.80x,https://s3-media4.fl.yelpcdn.com/bphoto/5iRixCvfWHjaHEFSpn7ZvQ/ms.jpg 1.67x,https://s3-media4.fl.yelpcdn.com/bphoto/5iRixCvfWHjaHEFSpn7ZvQ/180s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/bphoto/5iRixCvfWHjaHEFSpn7ZvQ/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="Q7WhrbFa7HO5PB-CcmXosg" href="/biz/cote-new-york?page_src=related_bizes"><span >Cote</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4 rating" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>


                <span class="review-count rating-qualifier">
            689 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$</span>
                        </span>

                        <span class="category-str-list">
                            Korean, Steakhouses, Cocktail Bars
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Flatiron
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/boucherie-union-square-new-york?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Boucherie Union Square" class="photo-box-img" height="60" src="https://s3-media4.fl.yelpcdn.com/bphoto/U90cr6FAs7CRF83TkbvrZg/60s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/U90cr6FAs7CRF83TkbvrZg/90s.jpg 1.50x,https://s3-media4.fl.yelpcdn.com/bphoto/U90cr6FAs7CRF83TkbvrZg/168s.jpg 2.80x,https://s3-media4.fl.yelpcdn.com/bphoto/U90cr6FAs7CRF83TkbvrZg/ms.jpg 1.67x,https://s3-media4.fl.yelpcdn.com/bphoto/U90cr6FAs7CRF83TkbvrZg/180s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/bphoto/U90cr6FAs7CRF83TkbvrZg/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="Up8vugqGau762YGEZIuRzw" href="/biz/boucherie-union-square-new-york?page_src=related_bizes"><span >Boucherie Union Square</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4-half rating" title="4.5 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.5 star rating">
    </div>


                <span class="review-count rating-qualifier">
            526 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$</span>
                        </span>

                        <span class="category-str-list">
                            Steakhouses, French, Cocktail Bars
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Gramercy
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/club-a-steakhouse-new-york?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Club A Steakhouse" class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/bphoto/8UNPe4NiTidCHN2VpxGbfg/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/8UNPe4NiTidCHN2VpxGbfg/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/bphoto/8UNPe4NiTidCHN2VpxGbfg/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/bphoto/8UNPe4NiTidCHN2VpxGbfg/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/bphoto/8UNPe4NiTidCHN2VpxGbfg/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/bphoto/8UNPe4NiTidCHN2VpxGbfg/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="g3xou3jd_TrJILJGFprNzw" href="/biz/club-a-steakhouse-new-york?page_src=related_bizes"><span >Club A Steakhouse</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4-half rating" title="4.5 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.5 star rating">
    </div>


                <span class="review-count rating-qualifier">
            2612 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$</span>
                        </span>

                        <span class="category-str-list">
                            Steakhouses
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Midtown East
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/boucherie-west-village-new-york-3?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Boucherie West Village" class="photo-box-img" height="60" src="https://s3-media4.fl.yelpcdn.com/bphoto/-5TXMVrjARsqKQHDg-hKbQ/60s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/-5TXMVrjARsqKQHDg-hKbQ/90s.jpg 1.50x,https://s3-media4.fl.yelpcdn.com/bphoto/-5TXMVrjARsqKQHDg-hKbQ/168s.jpg 2.80x,https://s3-media4.fl.yelpcdn.com/bphoto/-5TXMVrjARsqKQHDg-hKbQ/ms.jpg 1.67x,https://s3-media4.fl.yelpcdn.com/bphoto/-5TXMVrjARsqKQHDg-hKbQ/180s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/bphoto/-5TXMVrjARsqKQHDg-hKbQ/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="BT5mMck4BbzRTPQsignsuA" href="/biz/boucherie-west-village-new-york-3?page_src=related_bizes"><span >Boucherie West Village</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4-half rating" title="4.5 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.5 star rating">
    </div>


                <span class="review-count rating-qualifier">
            1105 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$</span>
                        </span>

                        <span class="category-str-list">
                            French, Cocktail Bars, Steakhouses
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            West Village
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/quality-meats-new-york?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Quality Meats" class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/bphoto/S0LrrLcBBkLThXw0qZLQLw/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/S0LrrLcBBkLThXw0qZLQLw/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/bphoto/S0LrrLcBBkLThXw0qZLQLw/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/bphoto/S0LrrLcBBkLThXw0qZLQLw/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/bphoto/S0LrrLcBBkLThXw0qZLQLw/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/bphoto/S0LrrLcBBkLThXw0qZLQLw/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="tTRAxtEWuaT-qOzyTIHlXw" href="/biz/quality-meats-new-york?page_src=related_bizes"><span >Quality Meats</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4 rating" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>


                <span class="review-count rating-qualifier">
            1912 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$$</span>
                        </span>

                        <span class="category-str-list">
                            Steakhouses
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Midtown West
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/fogo-de-chao-brazilian-steakhouse-new-york-4?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Fogo de Chao Brazilian Steakhouse" class="photo-box-img" height="60" src="https://s3-media4.fl.yelpcdn.com/bphoto/IcEOfdd4OS-lTrLLPigq9A/60s.jpg" srcset="https://s3-media4.fl.yelpcdn.com/bphoto/IcEOfdd4OS-lTrLLPigq9A/90s.jpg 1.50x,https://s3-media4.fl.yelpcdn.com/bphoto/IcEOfdd4OS-lTrLLPigq9A/168s.jpg 2.80x,https://s3-media4.fl.yelpcdn.com/bphoto/IcEOfdd4OS-lTrLLPigq9A/ms.jpg 1.67x,https://s3-media4.fl.yelpcdn.com/bphoto/IcEOfdd4OS-lTrLLPigq9A/180s.jpg 3.00x,https://s3-media4.fl.yelpcdn.com/bphoto/IcEOfdd4OS-lTrLLPigq9A/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="3ATN8HniUs93y-UU_T6JmQ" href="/biz/fogo-de-chao-brazilian-steakhouse-new-york-4?page_src=related_bizes"><span >Fogo de Chao Brazilian Steakhouse</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4 rating" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>


                <span class="review-count rating-qualifier">
            1084 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$</span>
                        </span>

                        <span class="category-str-list">
                            Brazilian, Steakhouses
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Midtown West
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/gallaghers-steakhouse-new-york?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Gallaghers Steakhouse" class="photo-box-img" height="60" src="https://s3-media1.fl.yelpcdn.com/bphoto/5Vq76A_SH0pHQdR3F6tvpw/60s.jpg" srcset="https://s3-media1.fl.yelpcdn.com/bphoto/5Vq76A_SH0pHQdR3F6tvpw/90s.jpg 1.50x,https://s3-media1.fl.yelpcdn.com/bphoto/5Vq76A_SH0pHQdR3F6tvpw/168s.jpg 2.80x,https://s3-media1.fl.yelpcdn.com/bphoto/5Vq76A_SH0pHQdR3F6tvpw/ms.jpg 1.67x,https://s3-media1.fl.yelpcdn.com/bphoto/5Vq76A_SH0pHQdR3F6tvpw/180s.jpg 3.00x,https://s3-media1.fl.yelpcdn.com/bphoto/5Vq76A_SH0pHQdR3F6tvpw/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="an1RDhCSkCX0AsH9uKpE_w" href="/biz/gallaghers-steakhouse-new-york?page_src=related_bizes"><span >Gallaghers Steakhouse</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-4 rating" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>


                <span class="review-count rating-qualifier">
            1030 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$</span>
                        </span>

                        <span class="category-str-list">
                            Steakhouses, Seafood, American (Traditional)
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Theater District
        </span>


            </div>
        </div>
    </div>

                    </li>
                    <li>
                            <div class="media-block media-block--12 biz-listing-medium js-related-business">
        <div class="media-avatar">
                            <div class="photo-box pb-60s">
                <a href="/biz/nusr-et-steakhouse-new-york-4?page_src=related_bizes" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Nusr-Et Steakhouse" class="photo-box-img" height="60" src="https://s3-media3.fl.yelpcdn.com/bphoto/PpAy466RGrnsf_8ckppw2w/60s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/PpAy466RGrnsf_8ckppw2w/90s.jpg 1.50x,https://s3-media3.fl.yelpcdn.com/bphoto/PpAy466RGrnsf_8ckppw2w/168s.jpg 2.80x,https://s3-media3.fl.yelpcdn.com/bphoto/PpAy466RGrnsf_8ckppw2w/ms.jpg 1.67x,https://s3-media3.fl.yelpcdn.com/bphoto/PpAy466RGrnsf_8ckppw2w/180s.jpg 3.00x,https://s3-media3.fl.yelpcdn.com/bphoto/PpAy466RGrnsf_8ckppw2w/120s.jpg 2.00x" width="60">

        </a>

    </div>



 </div>
        <div class="media-story">
            <div class="media-title clearfix">
                    


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="_NzX2Wq7fhvpenNYlShNeg" href="/biz/nusr-et-steakhouse-new-york-4?page_src=related_bizes"><span >Nusr-Et Steakhouse</span></a>


            </div>

            
                        <div class="biz-rating biz-rating-medium clearfix">
                



    <div class="i-stars i-stars--small-3-half rating" title="3.5 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="3.5 star rating">
    </div>


                <span class="review-count rating-qualifier">
            420 reviews
    </span>

        </div>



            <div class="related-business-passport-text">
                <div class="price-category">
                        <span class="bullet-after">
                            
        <span class="business-attribute price-range">$$$$</span>
                        </span>

                        <span class="category-str-list">
                            Steakhouses, Turkish, Halal
                        </span>
                </div>

                            <span class="neighborhood-str-list">
            Midtown West
        </span>


            </div>
        </div>
    </div>

                    </li>
            </ul>
        </div>


                                <div class="ywidget related-collections">
            <h3>
                Collections including Peter Luger
            </h3>
                <ul class="ylist">
            <li>
                <div class="media-block media-block--12">
                    <div class="media-avatar">
    <a class="collection-card-container u-inline-block" href="/collection/TdLOVnbx3Y3Qh1P9eL4xSw">
        <div class="collection-card_photo-box u-inline-block" style="background-image: url(https://s3-media4.fl.yelpcdn.com/bphoto/OB9dC1RKjEf14-pREQpVVg/300s.jpg)">
            <div class="collection-card_blur-overlay" style="background-image: url(https://s3-media4.fl.yelpcdn.com/bphoto/OB9dC1RKjEf14-pREQpVVg/300s.jpg)"></div>
            <div class="collection-card_text-overlay u-flex u-flex-dead-center">
                            <span class="collection-item-count">
                                <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-save icon--size-14 icon--white icon--fallback-inverted">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_save" />
    </svg>
</span>
                            </span>
            </div>
        </div>
    </a>
                    </div>
                    <div class="media-story">
                        <div class="media-title clearfix u-text-truncate">
                            <a class="" href="/collection/TdLOVnbx3Y3Qh1P9eL4xSw">Michelin Star Restaurant List For NYC</a>
                        </div>
                        <div class="u-text-small">By Erica A.</div>
                        <div class="u-absolute u-sticky-bottom">76 Places</div>
                    </div>
                </div>
            </li>
            <li>
                <div class="media-block media-block--12">
                    <div class="media-avatar">
    <a class="collection-card-container u-inline-block" href="/collection/85GGXVpkCMwKsElQewGdug">
        <div class="collection-card_photo-box u-inline-block" style="background-image: url(https://s3-media2.fl.yelpcdn.com/bphoto/noQBG75BIW_V3Fxcrct2Cw/300s.jpg)">
            <div class="collection-card_blur-overlay" style="background-image: url(https://s3-media2.fl.yelpcdn.com/bphoto/noQBG75BIW_V3Fxcrct2Cw/300s.jpg)"></div>
            <div class="collection-card_text-overlay u-flex u-flex-dead-center">
                            <span class="collection-item-count">
                                <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-save icon--size-14 icon--white icon--fallback-inverted">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_save" />
    </svg>
</span>
                            </span>
            </div>
        </div>
    </a>
                    </div>
                    <div class="media-story">
                        <div class="media-title clearfix u-text-truncate">
                            <a class="" href="/collection/85GGXVpkCMwKsElQewGdug">best restaurants to take the parents.</a>
                        </div>
                        <div class="u-text-small">By Peter D.</div>
                        <div class="u-absolute u-sticky-bottom">38 Places</div>
                    </div>
                </div>
            </li>
            <li>
                <div class="media-block media-block--12">
                    <div class="media-avatar">
    <a class="collection-card-container u-inline-block" href="/collection/lvVjeu7166VzSh_MeUAG_w">
        <div class="collection-card_photo-box u-inline-block" style="background-image: url(https://s3-media3.fl.yelpcdn.com/bphoto/NyWw-T2wPUDqqyVcz5qHsQ/300s.jpg)">
            <div class="collection-card_blur-overlay" style="background-image: url(https://s3-media3.fl.yelpcdn.com/bphoto/NyWw-T2wPUDqqyVcz5qHsQ/300s.jpg)"></div>
            <div class="collection-card_text-overlay u-flex u-flex-dead-center">
                            <span class="collection-item-count">
                                <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-save icon--size-14 icon--white icon--fallback-inverted">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_save" />
    </svg>
</span>
                            </span>
            </div>
        </div>
    </a>
                    </div>
                    <div class="media-story">
                        <div class="media-title clearfix u-text-truncate">
                            <a class="" href="/collection/lvVjeu7166VzSh_MeUAG_w">NYC Must Try</a>
                        </div>
                        <div class="u-text-small">By Charles N.</div>
                        <div class="u-absolute u-sticky-bottom">274 Places</div>
                    </div>
                </div>
            </li>
    </ul>
    <a class="link-more link-more-collections" href="/collections">More Collections</a>

        </div>


                            <div class="ywidget nearby-search-internal-link">
            <h3>Other Steakhouses Nearby</h3>
        <ul class="ylist">
                <li>
                    <a class="arrange arrange--6" href="/search?cflt=steak&amp;find_near=peter-luger-brooklyn-2">
                        <span class="arrange_unit arrange_unit--fill">
                            Find more Steakhouses near Peter Luger
                        </span>
                    </a>
                </li>
        </ul>
    </div>


                            <div class="ywidget nearby-links">
        <h3>Browse Nearby</h3>
        <ul class="ylist">
            <li>
                <a class="arrange arrange--6" href="/search?find_desc=Restaurants&amp;find_loc=178+Broadway%2C+Brooklyn%2C+NY+11211">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-food icon--size-18 icon--currentColor arrange_unit">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_food" />
    </svg>
</span>
                    <span class="arrange_unit arrange_unit--fill">Restaurants</span>
                </a>
            </li>
            <li>
                <a class="arrange arrange--6" href="/search?find_desc=Nightlife&amp;find_loc=178+Broadway%2C+Brooklyn%2C+NY+11211">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-nightlife icon--size-18 icon--currentColor arrange_unit">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_nightlife" />
    </svg>
</span>
                    <span class="arrange_unit arrange_unit--fill">Nightlife</span>
                </a>
            </li>
            <li>
                <a class="arrange arrange--6" href="/search?find_desc=Shopping&amp;find_loc=178+Broadway%2C+Brooklyn%2C+NY+11211">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-shopping icon--size-18 icon--currentColor arrange_unit">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_shopping" />
    </svg>
</span>
                    <span class="arrange_unit arrange_unit--fill">Shopping</span>
                </a>
            </li>
            <li>
                <a class="arrange arrange--6" href="/search?find_loc=178+Broadway%2C+Brooklyn%2C+NY+11211">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-overflow icon--size-18 icon--currentColor arrange_unit">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_overflow" />
    </svg>
</span>
                    <span class="arrange_unit arrange_unit--fill">Show all</span>
                </a>
            </li>
        </ul>
    </div>




                            <div class="ywidget">
        <h3>Dining in Brooklyn</h3>
        <ul class="ylist">
                <li>
                    <a class="arrange arrange--6" href="/search?attrs=RestaurantsReservations&amp;cflt=restaurants&amp;find_loc=Brooklyn%2C+NY">
                        <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-search-small icon--size-18 icon--currentColor arrange_unit">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_search_small" />
    </svg>
</span>
                        <span class="arrange_unit arrange_unit--fill">Search for Reservations</span>
                    </a>
                </li>
                <li>
                    <a class="arrange arrange--6" href="/search?attrs=OnlineReservations&amp;cflt=restaurants&amp;find_loc=Brooklyn%2C+NY">
                        <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-reservation icon--size-18 icon--currentColor arrange_unit">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_reservation" />
    </svg>
</span>
                        <span class="arrange_unit arrange_unit--fill">Book a Table in Brooklyn</span>
                    </a>
                </li>
        </ul>
    </div>


                                <div class="ywidget city-page-link">
            <h3>Best of Brooklyn</h3>
            <ul class="ylist">
                <li>
                    <a class="arrange arrange--6" href="/city/brooklyn">
                        <span class="arrange_unit arrange_unit--fill">
                            Things to do in Brooklyn
                        </span>
                    </a>
                </li>
            </ul>
        </div>


                    
        <div class="ywidget related-searches">
            <h3>People found Peter Luger by searching for…</h3>
            <ul class="ylist">
                    <li>
                            <a href="/search?find_desc=Gordon+Ramsay+Restaurant&amp;find_loc=Brooklyn%2C+NY">Gordon Ramsay Restaurant Brooklyn</a>
                    </li>
                    <li>
                            <a href="/search?find_desc=Restaurants+-+Steakhouses&amp;find_loc=Brooklyn%2C+NY">Restaurants - Steakhouses Brooklyn</a>
                    </li>
                    <li>
                            <a href="/c/brooklyn/steak">Steak House Brooklyn</a>
                    </li>
            </ul>
        </div>


                            <div class="ywidget sidebar-internal-links">
        <h3>Near Me</h3>
        <ul class="ylist">
                <li>
                    <a class="arrange arrange--6" href="https://www.yelp.com/nearme/claim-jumpers">
                        <span class="arrange_unit arrange_unit--fill">
                            Claim Jumpers Near Me
                        </span>
                    </a>
                </li>
                <li>
                    <a class="arrange arrange--6" href="https://www.yelp.com/nearme/steak">
                        <span class="arrange_unit arrange_unit--fill">
                            Steak Near Me
                        </span>
                    </a>
                </li>
        </ul>
    </div>



                </div>
    </div>
    </div>
        
        
    </div>


    
    
    

        
        
        
        
        
        

    
    
    
    
            
    
    
    
    
        
        



        
    


    

    

    

        




        
    

        




    


            <img height="1" width="1" alt="" style="display:none" src="https://www.facebook.com/tr?id=1432291440371688&amp;ev=NoScript&amp;cd%5Bcategory%5D%5B%5D=steak" />


    

        <div class="hidden send-to-phone-modal-content send-to-phone">
    <div class="modal_head">
            <h2>Text to Phone</h2>
    </div>
    <div class="modal_body">
            <div class="sms-modal-error hidden">


    <div class="alert alert-error alert--section alert--dismissible">
        <div class="alert-content-wrapper arrange">
            <div class="arrange_unit arrange_unit--fill">
                <div class="alert-message">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-warning icon--size-18 u-space-l1 u-space-r2">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_warning" />
    </svg>
</span>
                    Oops, looks like something’s wrong. Try again!
                </div>
            </div>
                <div class="arrange_unit">
                    <a class="dismiss-link js-alert-dismiss" href="javascript:;" title="Close">&times;</a>
                </div>
        </div>
    </div>
            </div>
            <div class="sms-modal-captcha-message hidden">


    <div class="alert alert-error alert--section alert--dismissible">
        <div class="alert-content-wrapper arrange">
            <div class="arrange_unit arrange_unit--fill">
                <div class="alert-message">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-warning icon--size-18 u-space-l1 u-space-r2">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_warning" />
    </svg>
</span>
                    Are you a human? Please complete the bot challenge below.
                </div>
            </div>
                <div class="arrange_unit">
                    <a class="dismiss-link js-alert-dismiss" href="javascript:;" title="Close">&times;</a>
                </div>
        </div>
    </div>
            </div>
            <div class="sms-modal-success send-success hidden">


    <div class="alert alert-success alert--section alert--dismissible">
        <div class="alert-content-wrapper arrange">
            <div class="arrange_unit arrange_unit--fill">
                <div class="alert-message">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-checkmark icon--size-18 u-space-l1 u-space-r2">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_checkmark" />
    </svg>
</span>
                    <strong>Sent!</strong> Check your phone to view the link now!
                </div>
            </div>
                <div class="arrange_unit">
                    <a class="dismiss-link js-alert-dismiss" href="javascript:;" title="Close">&times;</a>
                </div>
        </div>
    </div>
            </div>
            <div class="sms-modal-success initial-success hidden">


    <div class="alert alert-success alert--section alert--dismissible">
        <div class="alert-content-wrapper arrange">
            <div class="arrange_unit arrange_unit--fill">
                <div class="alert-message">
                    <span aria-hidden="true" style="width: 18px; height: 18px;" class="icon icon--18-checkmark icon--size-18 u-space-l1 u-space-r2">
    <svg role="img" class="icon_svg">
        <use xlink:href="#18x18_checkmark" />
    </svg>
</span>
                    
                </div>
            </div>
                <div class="arrange_unit">
                    <a class="dismiss-link js-alert-dismiss" href="javascript:;" title="Close">&times;</a>
                </div>
        </div>
    </div>
            </div>
            <div class="arrange arrange--middle arrange-30 u-text-centered">
                <div class="sms-modal-body cross-browser-center">
                        <div class="arrange arrange--9 arrange-equal u-space-b5 u-space-t2">
                            <div class="arrange_unit u-flex-inline u-text-left">
                                    <div class="media-block media-block--12 biz-listing-medium nowrap">
        <div class="media-avatar">
                            <div class="photo-box pb-90s">
                <a href="/biz/peter-luger-brooklyn-2" class="js-analytics-click" data-analytics-label="biz-photo">
                <img alt="Peter Luger" class="photo-box-img" height="90" src="https://s3-media3.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/90s.jpg" srcset="https://s3-media3.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/180s.jpg 2.00x,https://s3-media3.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/ms.jpg 1.11x,https://s3-media3.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/120s.jpg 1.33x,https://s3-media3.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/168s.jpg 1.87x,https://s3-media3.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/ls.jpg 2.78x,https://s3-media3.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/258s.jpg 2.87x" width="90">

        </a>

    </div>




        </div>
        <div>
            <div class="media-title clearfix">
                        


        <a class="biz-name js-analytics-click" data-analytics-label="biz-name" data-hovercard-id="FM6GMdzV5FR83nQvEZJu3A" href="/biz/peter-luger-brooklyn-2"><span >Peter Luger</span></a>


            </div>
                        <div class="biz-rating biz-rating-large clearfix" >
                



    <div class="i-stars i-stars--regular-4 rating-large" title="4.0 star rating">
        <img class="offscreen" height="303" src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png" width="84" alt="4.0 star rating">
    </div>


                    <span class="review-count rating-qualifier">
            5437 reviews
    </span>

        </div>


                        <div class="price-category">
                <span class="bullet-after">
            
        <span class="business-attribute price-range">$$$$</span>
        </span>
            <span class="category-str-list">
                    <a href="/search?cflt=steak&amp;find_loc=Brooklyn%2C+NY">Steakhouses</a>
    </span>


    </div>


                    <small class="biz-city">Brooklyn, NY</small>

        </div>
    </div>

                            </div>
                        </div>
                        <div class="sms-modal-secondary-text u-space-b5 u-space-r5 u-space-l5">Text a link to your phone so you can quickly get directions, see photos, and read reviews on the go!</div>
                    <div class="arrange arrange--9 arrange-equal">
                        <div class="arrange_unit u-flex-inline">



    <form action="/sms_app_link/peter-luger-brooklyn-2" class="sms-modal-form yform u-space-b0 u-space-r2" data-modal-name="send-to-phone" method="POST">
        <input     type="hidden"    name="csrftok"    class="csrftok"    value="f091e5a484d734f178b59138f4759aad66870b161bc619a9062a0d7b53d13374">
                                    <div class="arrange phone-input">
        <div class="arrange_unit">


    <div class="yselect">
                <select class="u-relative" name="country_code">
                            <option value="+1" selected>+1</option>
                            <option value="+31">+31</option>
                            <option value="+32">+32</option>
                            <option value="+33">+33</option>
                            <option value="+34">+34</option>
                            <option value="+39">+39</option>
                            <option value="+41">+41</option>
                            <option value="+43">+43</option>
                            <option value="+44">+44</option>
                            <option value="+45">+45</option>
                            <option value="+46">+46</option>
                            <option value="+47">+47</option>
                            <option value="+48">+48</option>
                            <option value="+49">+49</option>
                            <option value="+52">+52</option>
                            <option value="+54">+54</option>
                            <option value="+55">+55</option>
                            <option value="+56">+56</option>
                            <option value="+60">+60</option>
                            <option value="+61">+61</option>
                            <option value="+63">+63</option>
                            <option value="+64">+64</option>
                            <option value="+65">+65</option>
                            <option value="+81">+81</option>
                            <option value="+90">+90</option>
                            <option value="+351">+351</option>
                            <option value="+353">+353</option>
                            <option value="+358">+358</option>
                            <option value="+420">+420</option>
                            <option value="+852">+852</option>
                            <option value="+886">+886</option>
                </select>
        <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-triangle-down icon--size-14 icon--currentColor u-triangle-direction-down yselect_arrow">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_triangle_down" />
    </svg>
</span>
    </div>
        </div>

        <div class="arrange_unit">
            <input name="national_number" class="phone-input-field u-relative" placeholder="Phone Number" value="" type="text">
        </div>
    </div>


                                        <input type="hidden" name="fb:app_id" value="97534753161">
                                        <input type="hidden" name="og:description" value="Steakhouses in Brooklyn, NY">
                                        <input type="hidden" name="og:image" value="https://s3-media2.fl.yelpcdn.com/bphoto/DnReRUkXRtsmKycQEYl0pg/o.jpg">
                                        <input type="hidden" name="og:image:height" value="1088">
                                        <input type="hidden" name="og:image:width" value="1600">
                                        <input type="hidden" name="og:site_name" value="Yelp">
                                        <input type="hidden" name="og:title" value="Peter Luger - Williamsburg - South Side - Brooklyn, NY">
                                        <input type="hidden" name="og:type" value="yelpyelp:business">
                                        <input type="hidden" name="og:url" value="https://www.yelp.com/biz/peter-luger-brooklyn-2">

                                <div id="recaptcha-container" style="display: none;">
                                            <div id="recaptcha-widget" style="display: none;"></div>

        <noscript>
        <div>
            <div class="g-recaptcha-challenge-outer-wrap-noscript">
                <div class="g-recaptcha-challenge-inner-wrap-noscript">
                    <iframe
                        class="g-racaptcha-challenge-frame-noscript"
                        src="https://www.google.com/recaptcha/api/fallback?k=6Le5OSYTAAAAADy8seTrWT0eqpS795iV32Gm9Ag1"
                        frameborder="0" scrolling="no"
                    ></iframe>
                </div>
            </div>
            <div class="g-recaptcha-response-wrap-noscript">
                <textarea
                    id="g-recaptcha-response" name="g-recaptcha-response"
                    class="g-recaptcha-response"
                ></textarea>
            </div>
        </div>
    </noscript>


                                </div>
    </form>
                        </div>
                        <div class="arrange_unit u-flex-inline">
                            <button type="submit" value="submit" class="ybtn ybtn--primary ybtn--small sms-modal-button"><span>Text Link</span></button>
                        </div>
                        <div class="u-legal-text u-text-subtle">
                            Your carrier’s rates may apply
                        </div>
                    </div>
                </div>
            </div>
    </div>
    </div>

    



    

    

        <script>
            (function() {
                var main = null;

                var main=function(a){adroll_adv_id="BHPKS4B4ONEJJMGH4QCJZR";adroll_pix_id="QB5JPFIKRZDSBOZSULG4YB";adroll_user_identifier=a.user_identifier;adroll_email=a.hashed_email;adroll_shop_id="biz_id" in a.custom_data?a.custom_data.biz_id:undefined;window.adroll_callbacks=[function(){if(a.redirect_url){setTimeout(function(){location.replace(a.redirect_url)},"redirect_wait_time_in_seconds" in a?a.redirect_wait_time_in_seconds*1000:0)}}];(function(){var b=function(){if(document.readyState&&!/loaded|complete/.test(document.readyState)){setTimeout(b,10);
return}if(!window.__adroll_loaded){__adroll_loaded=true;setTimeout(b,50);return}var d=document.createElement("script");var c="https://s.adroll.com";d.setAttribute("async","true");d.type="text/javascript";d.src=c+"/j/roundtrip.js";((document.getElementsByTagName("head")||[null])[0]||document.getElementsByTagName("script")[0].parentNode).appendChild(d)};if(window.addEventListener){window.addEventListener("load",b,false)}else{window.attachEvent("onload",b)}}());adroll_custom_data=a.custom_data};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main({"user_identifier": "DE1CA55AE6AF7DF3", "hashed_email": null, "custom_data": {"city": "Brooklyn", "category_aliases": "steak", "adroll_geoquad_v1": 803417, "user_type": 0, "state": "NY", "biz_id": "4yPqqJDJOQX69gC66YUDkA"}});
            })();
    </script>

    <script>
            (function() {
                var main = null;

                var main=function(){!function(h,a,i,c,j,d,g){if(h.fbq){return}j=h.fbq=function(){j.callMethod?j.callMethod.apply(j,arguments):j.queue.push(arguments)};if(!h._fbq){h._fbq=j}j.push=j;j.loaded=!0;j.version="2.0";j.queue=[];d=a.createElement(i);d.async=!0;d.src=c;g=a.getElementsByTagName(i)[0];g.parentNode.insertBefore(d,g)}(window,document,"script","https://connect.facebook.net/en_US/fbevents.js");fbq("init","102029836881428");fbq("track","PageView")};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main();
            })();
    </script>

    

    

            

    
    
    
                    </div>
                </div>

    </div>


                    <!-- <a href="/biz/outlook-autumn-market-fundamental-catwalk-flimsy-roost-legibility-individualism-grocer-predestination-5">yelp</a> -->
                    <!-- google_ad_section_start(weight=ignore) -->
    <div class="main-content-wrap main-content-wrap--separated">
        <div class="content-container">
            <div class="main-footer webview-hidden">
    <div class="main-footer_section main-footer_menu arrange arrange--equal arrange--30 arrange--stack-small u-sm-space-b0">
                    <div class="main-footer_item responsive-hidden-small">
            <div class="footer-menu responsive-hidden-small">
        <h3 class="footer-menu_header">About</h3>
        <ul class="footer-menu_list">
                <li class="footer-menu_item">
                    <a href="https://yelp.com/about">About Yelp</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/careers/home">Careers</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/press">Press</a>
                </li>
                <li class="footer-menu_item">
                    <a href="http://yelp-ir.com/">Investor Relations</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/guidelines">Content Guidelines</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/static?p=tos">Terms of Service</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/tos/privacy_policy">Privacy Policy</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/static?locale=en_US&amp;p=privacy#third-parties">Ad Choices</a>
                </li>
        </ul>
    </div>

    </div>

                    <div class="main-footer_item responsive-hidden-small">
            <div class="footer-menu responsive-hidden-small">
        <h3 class="footer-menu_header">Discover</h3>
        <ul class="footer-menu_list">
                <li class="footer-menu_item">
                    <a href="/costs">Yelp Project Cost Guides</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/collections">Collections</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/talk">Talk</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/events">Events</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/local_yelp">The Local Yelp</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://officialblog.yelp.com/">Yelp Blog</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://www.yelp-support.com/?l=en_US">Support</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/yelpmobile?source=footer">Yelp Mobile</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/developers?country=US">Developers</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/rss">RSS</a>
                </li>
        </ul>
    </div>

    </div>

                    <div class="main-footer_item responsive-hidden-small">
            <div class="footer-menu responsive-hidden-small">
        <h3 class="footer-menu_header">Yelp for Business Owners</h3>
        <ul class="footer-menu_list">
                <li class="footer-menu_item">
                    <a href="https://biz.yelp.com/?utm_campaign=claim_business&amp;utm_content=claim_footer_link&amp;utm_medium=www&amp;utm_source=footer">Claim your Business Page</a>
                </li>
                <li class="footer-menu_item">
                    <a href="/advertise">Advertise on Yelp</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://www.yelpreservations.com/">Yelp Reservations</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://www.yelpwifi.com/">Yelp WiFi</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://www.nowait.com/restaurants/demo/">Yelp Nowait</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://biz.yelp.com/support/case_studies">Business Success Stories</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://www.yelp-support.com/Yelp_for_Business_Owners?l=en_US">Business Support</a>
                </li>
                <li class="footer-menu_item">
                    <a href="https://www.yelpblog.com/section/business">Yelp Blog for Business Owners</a>
                </li>
        </ul>
    </div>

    </div>

                    <div class="main-footer_item">
        <div class="footer-menu languages-menu">
                <div class="footer-menu_section footer-language">
                    <h3 class="footer-menu_header responsive-hidden-small">Languages</h3>
                    

    <div class="dropdown js-dropdown dropdown--hover dropdown--boxed-on-mobile dropdown--separate-groups dropdown--restricted">
        

    <div class="dropdown_toggle js-dropdown-toggle" tabindex="-1">
        <a
            class="dropdown_toggle-action"
                href="javascript:;"
                data-dropdown-prefix="English"
            aria-haspopup="true"
            role="button"
        >
                <span class="dropdown_prefix">
                    English
                </span>
            <span class="dropdown_toggle-text js-dropdown-toggle-text" data-dropdown-initial-text=""></span>
            <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-triangle-down icon--size-14 icon--currentColor u-triangle-direction-down dropdown_arrow">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_triangle_down" />
    </svg>
</span>
        </a>
    </div>

            <div class="dropdown_menu-container js-dropdown_menu-container">
        <div class="dropdown_menu js-dropdown-menu">
            <div class="dropdown_menu-inner">
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://ms.yelp.my/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Bahasa Malaysia (Malaysia)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.cz/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Čeština (Česká republika)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.dk/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Dansk (Danmark)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.de/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Deutsch (Deutschland)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://de.yelp.ch/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Deutsch (Schweiz)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.at/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Deutsch (Österreich)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.au/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Australia)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.be/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Belgium)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.ca/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Canada)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.com.hk/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Hong Kong)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.my/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Malaysia)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.co.nz/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (New Zealand)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.com.ph/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Philippines)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.ie/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Republic of Ireland)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.sg/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Singapore)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.ch/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (Switzerland)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.co.uk/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (United Kingdom)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                English (United States)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.ar/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Español (Argentina)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.cl/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Español (Chile)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.es/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Español (España)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.mx/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Español (México)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://fil.yelp.com.ph/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Filipino (Pilipinas)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://fr.yelp.be/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Français (Belgique)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://fr.yelp.ca/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Français (Canada)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.fr/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Français (France)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://fr.yelp.ch/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Français (Suisse)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.it/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Italiano (Italia)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://it.yelp.ch/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Italiano (Svizzera)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://nl.yelp.be/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Nederlands (België)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.nl/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Nederlands (Nederland)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.no/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Norsk (Norge)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.pl/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Polski (Polska)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.br/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Português (Brasil)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.pt/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Português (Portugal)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://fi.yelp.fi/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Suomi (Suomi)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://sv.yelp.fi/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Svenska (Finland)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.se/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Svenska (Sverige)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.tr/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                Türkçe (Türkiye)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.co.jp/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                日本語 (日本)
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.tw/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                繁體中文 (台灣)
        </span>
    </a>

                            </li>
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://zh.yelp.com.hk/biz/peter-luger-brooklyn-2" rel="nofollow" role="menuitem">
        <span class="dropdown_label">
                繁體中文 (香港)
        </span>
    </a>

                            </li>
                    </ul>
            </div>
        </div>
    </div>

    </div>

                </div>
            <div class="footer-menu_section footer-country">
                <h3 class="footer-menu_header responsive-hidden-small">Countries</h3>
                

    <div class="dropdown js-dropdown dropdown--hover dropdown--boxed-on-mobile dropdown--restricted">
        

    <div class="dropdown_toggle js-dropdown-toggle" tabindex="-1">
        <a
            class="dropdown_toggle-action"
                href="javascript:;"
                data-dropdown-prefix="United States"
            aria-haspopup="true"
            role="button"
        >
                <span class="dropdown_prefix">
                    United States
                </span>
            <span class="dropdown_toggle-text js-dropdown-toggle-text" data-dropdown-initial-text=""></span>
            <span aria-hidden="true" style="width: 14px; height: 14px;" class="icon icon--14-triangle-down icon--size-14 icon--currentColor u-triangle-direction-down dropdown_arrow">
    <svg role="img" class="icon_svg">
        <use xlink:href="#14x14_triangle_down" />
    </svg>
</span>
        </a>
    </div>

            <div class="dropdown_menu-container js-dropdown_menu-container">
        <div class="dropdown_menu js-dropdown-menu">
            <div class="dropdown_menu-inner">
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.ar/" role="menuitem">
        <span class="dropdown_label">
                Argentina
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.au/" role="menuitem">
        <span class="dropdown_label">
                Australia
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.at/" role="menuitem">
        <span class="dropdown_label">
                Austria
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.be/" role="menuitem">
        <span class="dropdown_label">
                Belgium
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.br/" role="menuitem">
        <span class="dropdown_label">
                Brazil
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.ca/" role="menuitem">
        <span class="dropdown_label">
                Canada
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.cl/" role="menuitem">
        <span class="dropdown_label">
                Chile
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.cz/" role="menuitem">
        <span class="dropdown_label">
                Czech Republic
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.dk/" role="menuitem">
        <span class="dropdown_label">
                Denmark
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://fi.yelp.fi/" role="menuitem">
        <span class="dropdown_label">
                Finland
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.fr/" role="menuitem">
        <span class="dropdown_label">
                France
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.de/" role="menuitem">
        <span class="dropdown_label">
                Germany
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.com.hk/" role="menuitem">
        <span class="dropdown_label">
                Hong Kong
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.it/" role="menuitem">
        <span class="dropdown_label">
                Italy
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.co.jp/" role="menuitem">
        <span class="dropdown_label">
                Japan
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.my/" role="menuitem">
        <span class="dropdown_label">
                Malaysia
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.mx/" role="menuitem">
        <span class="dropdown_label">
                Mexico
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.co.nz/" role="menuitem">
        <span class="dropdown_label">
                New Zealand
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.no/" role="menuitem">
        <span class="dropdown_label">
                Norway
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.com.ph/" role="menuitem">
        <span class="dropdown_label">
                Philippines
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.pl/" role="menuitem">
        <span class="dropdown_label">
                Poland
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.pt/" role="menuitem">
        <span class="dropdown_label">
                Portugal
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.ie/" role="menuitem">
        <span class="dropdown_label">
                Republic of Ireland
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.sg/" role="menuitem">
        <span class="dropdown_label">
                Singapore
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.es/" role="menuitem">
        <span class="dropdown_label">
                Spain
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.se/" role="menuitem">
        <span class="dropdown_label">
                Sweden
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://en.yelp.ch/" role="menuitem">
        <span class="dropdown_label">
                Switzerland
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.tw/" role="menuitem">
        <span class="dropdown_label">
                Taiwan
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.nl/" role="menuitem">
        <span class="dropdown_label">
                The Netherlands
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com.tr/" role="menuitem">
        <span class="dropdown_label">
                Turkey
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.co.uk/" role="menuitem">
        <span class="dropdown_label">
                United Kingdom
        </span>
    </a>

                            </li>
                    </ul>
                    <ul class="dropdown_menu-group" role="menu" aria-hidden="false">
                            <li class="dropdown_item " role="presentation">
                                        <a class="dropdown_link js-dropdown-link" href="https://www.yelp.com/" role="menuitem">
        <span class="dropdown_label">
                United States
        </span>
    </a>

                            </li>
                    </ul>
            </div>
        </div>
    </div>

    </div>

            </div>
        </div>
    </div>

                
    </div>

                                    <div class="main-footer_section lesser-text responsive-hidden-small">
            <div class="footer-menu footer-menu--inline footer-menu--bordered footer-menu--separated footer-menu--directory">
                <span class="footer-menu_header">Site Map</span>
                <ul class="footer-menu_list">
                        <li class="footer-menu_item"><a href="/atlanta">Atlanta</a></li>
                        <li class="footer-menu_item"><a href="/austin">Austin</a></li>
                        <li class="footer-menu_item"><a href="/boston">Boston</a></li>
                        <li class="footer-menu_item"><a href="/chicago">Chicago</a></li>
                        <li class="footer-menu_item"><a href="/dallas">Dallas</a></li>
                        <li class="footer-menu_item"><a href="/denver">Denver</a></li>
                        <li class="footer-menu_item"><a href="/detroit">Detroit</a></li>
                        <li class="footer-menu_item"><a href="/honolulu">Honolulu</a></li>
                        <li class="footer-menu_item"><a href="/houston">Houston</a></li>
                        <li class="footer-menu_item"><a href="/la">Los Angeles</a></li>
                        <li class="footer-menu_item"><a href="/miami">Miami</a></li>
                        <li class="footer-menu_item"><a href="/minneapolis">Minneapolis</a></li>
                        <li class="footer-menu_item"><a href="/nyc">New York</a></li>
                        <li class="footer-menu_item"><a href="/philadelphia">Philadelphia</a></li>
                        <li class="footer-menu_item"><a href="/portland">Portland</a></li>
                        <li class="footer-menu_item"><a href="/sacramento">Sacramento</a></li>
                        <li class="footer-menu_item"><a href="/san-diego">San Diego</a></li>
                        <li class="footer-menu_item"><a href="/sf">San Francisco</a></li>
                        <li class="footer-menu_item"><a href="/san-jose">San Jose</a></li>
                        <li class="footer-menu_item"><a href="/seattle">Seattle</a></li>
                        <li class="footer-menu_item"><a href="/dc">Washington, DC</a></li>
                        <li class="footer-menu_item"><a href="/locations">More Cities</a></li>
                </ul>
            </div>
        </div>

            
                <div class="main-footer_mobile-links responsive-visible-small-block hidden-non-responsive-block">
            <ul class="footer-menu--inline u-text-centered">
            <li class="footer-menu_item">
                <a href="https://yelp.com/about">About</a>
            </li>
            <li class="footer-menu_item">
                <a href="https://officialblog.yelp.com/">Blog</a>
            </li>
            <li class="footer-menu_item">
                <a href="https://www.yelp-support.com/?l=en_US">Support</a>
            </li>
            <li class="footer-menu_item">
                <a href="/static?p=tos">Terms</a>
            </li>
    </ul>


    </div>

            
    <div class="main-footer_city-landscape-img" role="presentation"></div>

                <small class="main-footer_copyright">
            Copyright © 2004–2019 Yelp Inc. Yelp, <img src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/17089be275f0/assets/img/logos/logo_desktop_xsmall_outline.png" alt="Yelp logo" class="main-footer_logo-copyright" srcset="https://s3-media1.fl.yelpcdn.com/assets/srv0/yelp_styleguide/0aade8725c91/assets/img/logos/logo_desktop_xsmall_outline@2x.png 2x">, <img src="https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_styleguide/58cfc999e1f5/assets/img/logos/burst_desktop_xsmall_outline.png" alt="Yelp burst" class="main-footer_logo-burst" srcset="https://s3-media1.fl.yelpcdn.com/assets/srv0/yelp_styleguide/dcb526e86d86/assets/img/logos/burst_desktop_xsmall_outline@2x.png 2x"> and related marks are registered trademarks of Yelp.
    </small>

                <!--Source: Freebase (https://www.freebase.com/), licensed under CC-BY (https://creativecommons.org/licenses/by/2.5/)-->


            </div>
        </div>
    </div>
<!-- google_ad_section_end -->

            <div style="display: none;" id="page-popups">
            </div>

                    
        
        

        </div> 

        

    <script>            yConfig = {"cookies": {"RECENT_LOCATIONS": "recentlocations", "SEARCH_PREFS": "searchPrefs", "JS_TRACK": "js_track", "DEBUG": "debug", "SELECTED_ACTIVITY_FEED": "fd", "SIGNUP_TRACK": "strack", "EXPR_OVERRIDE": "expr", "COOKIE_INFO": "cim", "LOCATION": "location", "APP_PITCH_CLOSED": "ap", "ADMIN_CM_SCOREBOARD_SETTINGS": "acms", "SEARCH_SUGGEST_INFO": "search_suggest", "DISMISSIBLE_MESSAGE_COOKIE": "dm", "UPGRADES_IN_ARREARS": "uia", "USED_LOCALE_SELECTOR": "used_locale_selector", "PRIVACY_POLICY_NOTICE": "ppn", "YR_VALENTINES_PROMOTIONS": "yrvp", "ADMIN_SEARCH_USERDATA": "ud", "RESERVATION_SEARCH_PARAMS": "rsp", "PREVIOUS_UNIQUE_REQUEST_ID": "pid", "XCJ": "xcj", "HOMEPAGE_SERVICE_ACTIVITY_FEED": "hsfd", "MESSAGE_SENDER_PASSED_CAPTCHA": "mspc", "QUANTCAST": "qntcst", "AD_CLICKS": "adc"}, "webviewFlow": null, "enabledSitRepChannels": {"vertical_search_platform": true, "biz_map_view": true, "ytp_session_events": true, "message_the_business": true, "call_to_action": true, "search_suggest_events": true, "highlights": true, "biz_services": true, "traffic_quality": true, "biz_directions": true, "vertical_search_reservation": true}, "isWebviewRequest": false, "uniqueRequestId": "5e84cbeca1806274", "cdn": {"hostnamePattern": {"s3": "s3-media%s.fl.yelpcdn.com", "yelp": "media%s.fl.yelpcdn.com"}, "shardCount": 4}, "yelpcodeTemplateVersion": "e8b0b8210c", "isLoggedIn": false, "imageUrls": {"mapMarkers": "https://s3-media1.fl.yelpcdn.com/assets/2/www/img/fca83ddb80f2/map/map_markers_sprite.png", "stars": "https://s3-media2.fl.yelpcdn.com/assets/srv0/yelp_design_web/9b34e39ccbeb/assets/img/stars/stars.png"}, "mapMarkerSerial": "20160801", "uaInfo": {"device": {"family": "Other"}, "os": {"major": null, "patch_minor": null, "minor": null, "family": "Other", "patch": null}, "string": "curl/7.54.0", "user_agent": {"major": null, "minor": null, "family": "Other", "patch": null}, "attributes": {}}, "isSitRepEnabled": true, "comscore": {"c15": "", "c3": "", "c2": 7130511, "c1": 2, "c6": "", "c5": "", "c4": ""}, "isBugsnagEnabled": true, "webviewInfo": {}, "deprecatedEncryptedYUV": "uUzVE3xrF4-y5Xh09NWX72vDZq0Y5FEIzIpDChwe-KAfRBHEunJxhN313wzLMns9t0smHRLWWB6mvgf1k_xN-NWoQ-sL8AA2", "vendorExternalURLs": {}, "cookieDomain": ".www.yelp.com", "recaptchaPublicKey": "6Le5OSYTAAAAADy8seTrWT0eqpS795iV32Gm9Ag1", "isLive": true, "recaptchaV3PublicKey": "6Le2lKQUAAAAAMMHRuHPaDwOMSodMg2FiVTGSw2Y", "isClientErrorsEnabled": false, "googlePlacesUrl": "//maps.googleapis.com/maps/api/js?key=AIzaSyByT6TXVL4jSf0MrzRwg0eG66A8-P4r0ps\u0026language=en\u0026libraries=places"};
</script>
        
        

        

            <noscript><img src="https://sb.scorecardresearch.com/p?cj=1&amp;c15=&amp;c3=&amp;c2=7130511&amp;c1=2&amp;c6=&amp;c5=&amp;c4="></noscript>

                <script>
            (function() {
                var main = null;

                var main=(function(){function a(c,b){window._qevents=window._qevents||[];window._qevents.push({qacct:b});window.yDFP.quantcast={};window.yDFP.quantcast.cookieName=c;window.yDFP.quantcast.adTagName="qncst_segs";window.yDFP.quantcast.getTargetingData=function(){var g=new RegExp(window.yDFP.quantcast.cookieName+"=([^;]*)");var f=document.cookie.match(g);if(f){var e=decodeURIComponent(f[1]);var d=e.split(",");return{key:window.yDFP.quantcast.adTagName,value:d}}else{return null}};if(Boolean(window.googletag)&&Boolean(window.googletag.cmd)){window.googletag.cmd.push(function(){var d=window.yDFP.quantcast.getTargetingData();
if(d){window.googletag.pubads().setTargeting(d.key,d.value)}})}}window.yDFP=window.yDFP||{};window.yDFP.initQuantcast=a;return a})();

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main("qntcst","p-M4yfUTCPeS3vn");
            })();
    </script>
                    <noscript>
        <img
            style="display: none;"
            src="https://pixel.quantserve.com/pixel/p-M4yfUTCPeS3vn.gif"
            alt="Quantcast"
            border="0"
            height="1"
            width="1"
        />
    </noscript>

                

            <script src="https://s3-media2.fl.yelpcdn.com/assets/2/www/js/74906f832f45/assets/vendor_external/closure_localizations_en_US.min.js" type="text/javascript"></script>

        <script id="__LOADABLE_REQUIRED_CHUNKS__" type="application/json">[]</script>

        <script src="https://cdnjs.cloudflare.com/ajax/libs/babel-polyfill/6.23.0/polyfill.min.js" crossorigin="anonymous" integrity="sha384-FbHUaR69a828hqWjPw4PFllFj1bvveKOTWORGkyosCw720HXy/56+2hSuQDaogMb"></script><script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/1.8.3/jquery.min.js"></script><script>if(document.readyState === 'interactive') jQuery.ready();
        yConfig = window.yConfig || {};
        yConfig.vendorExternalURLs = yConfig.vendorExternalURLs || {};
        </script><script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.8.6/umd/react.production.min.js" crossorigin="anonymous" integrity="sha384-qn+ML/QkkJxqn4LLs1zjaKxlTg2Bl/6yU/xBTJAgxkmNGc6kMZyeskAG0a7eJBR1"></script><script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.8.6/umd/react-dom.production.min.js" crossorigin="anonymous" integrity="sha384-85IMG5rvmoDsmMeWK/qUU4kwnYXVpC+o9hoHMLi4bpNR+gMEiPLrvkZCgsr7WWgV"></script><script src="https://s3-media0.fl.yelpcdn.com/assets/public/module_yelp_main_webpack_runtime.yji-90bc2d843874f6a4b657.js" crossorigin="anonymous"></script><script src="https://s3-media0.fl.yelpcdn.com/assets/public/module_yelp_js_packages.yji-5119cd813f2478eec2a8.chunk.js" crossorigin="anonymous"></script><script src="https://s3-media0.fl.yelpcdn.com/assets/public/common_chunk.yji-06a2a7c904834af30dc6.chunk.js" crossorigin="anonymous"></script><script src="https://s3-media0.fl.yelpcdn.com/assets/public/module_core.yji-6e3fb2edecad9b925cb2.chunk.js" crossorigin="anonymous"></script><script src="https://s3-media0.fl.yelpcdn.com/assets/public/module_react_biz_details.yji-b15ac7d337196c827fbb.chunk.js" crossorigin="anonymous"></script><script src="https://s3-media0.fl.yelpcdn.com/assets/public/module_misc_react.yji-f6d81dddf895badc1605.chunk.js" crossorigin="anonymous"></script><script src="https://s3-media0.fl.yelpcdn.com/assets/public/module_biz_details.yji-822fdfe347d9517cad1e.chunk.js" crossorigin="anonymous"></script><script>
                yConfig.vendorExternalURLs["plugin-detect"]=['https://s3-media4.fl.yelpcdn.com/assets/2/www/js/bfd62b330c7a/assets/vendor_external/plugin-detect-0.6.3.min.js']
                
                yConfig.vendorExternalURLs["fast-click"]=['https://s3-media1.fl.yelpcdn.com/assets/2/www/js/960df06ce458/assets/vendor_external/fastclick.min.js']
                yelp.www.util.analytics.initJSCustomDimensions();
yelp.www.init.initPromotions();
yelp.www.init.initSearchSuggest("#find_desc", ".find-decorator", "#dropperText_Mast", true);
yelp.www.init.initAccountMenu("#topbar-account-item");
yelp.www.init.initHeaderNavTracking();
yelp.www.init.initTimeOnPageAnalyticsBeforeUnload();
yelp.init.initAdAcknowledgment("/ad_acknowledgment", "fc49573141e40862af9e40fa4fd3947dcdcc56af4df0fef5d1b307d934ed7265");
yelp.common.init.adVisibilityTracker({"servletURI": "/ad_visibility", "loggingCSRF": "ae53f9113c6f1685a2d12315b098f1bb31f62f3580df2738f30d386a5eca7862"});
yelp.init.initDelegatedDropdowns();
yelp.init.attachGhostHovercard("qype", "    \u003cdiv id=\"qype-ghost-user-hovercard\"\u003e\n        \u003cdiv class=\"whatsthis-hovercard\"\u003e\n            \u003ch3\u003eWhat's This?\u003c/h3\u003e\n            \u003cp\u003eThis user has arrived from Qype, a company acquired by Yelp in 2012. We have integrated the two sites to bring you one great local experience.\u003c/p\u003e\n        \u003c/div\u003e\n    \u003c/div\u003e\n");
yelp.init.attachGhostHovercard("rk", "    \u003cdiv id=\"rk-ghost-user-hovercard\"\u003e\n        \u003cdiv class=\"whatsthis-hovercard\"\u003e\n            \u003ch3\u003eWhat's This?\u003c/h3\u003e\n            \u003cp\u003eThis user has arrived from Restaurant-Kritik.de, a company acquired by Yelp in 2014. We have integrated the two sites to bring you one great local experience.\u003c/p\u003e\n        \u003c/div\u003e\n    \u003c/div\u003e\n");
yelp.init.attachGhostHovercard("cityvox", "    \u003cdiv id=\"cityvox-ghost-user-hovercard\"\u003e\n        \u003cdiv class=\"whatsthis-hovercard\"\u003e\n            \u003ch3\u003eWhat's This?\u003c/h3\u003e\n            \u003cp\u003eThis user has arrived from Cityvox, a company acquired by Yelp in 2014. We have integrated the two sites to bring you one great local experience.\u003c/p\u003e\n        \u003c/div\u003e\n    \u003c/div\u003e\n");
yelp.www.init.initAsyncImageLoader();
yelp.pages.init();
yelp.common.init.facebook("97534753161", "//connect.facebook.net/en_US/sdk.js", "v3.1", "v3", ["email", "user_birthday", "user_gender"], [], ["user_friends"], true);
yelp.init.quantcast();
window.YelpReactBizDetails.hydrateReactComponents();
window.PageSpeedLogging.initBunsenLogging("biz_details#default");
window.YelpReactLSATHeader.hydrateReactComponent();

                yConfig.vendorExternalURLs["video.js"]=['https://s3-media4.fl.yelpcdn.com/assets/2/www/js/b10088a14a6e/assets/vendor_external/video.min.js']
                
                yConfig.vendorExternalURLs["jquery-flot"]=['https://s3-media4.fl.yelpcdn.com/assets/2/www/js/af411e8c182d/assets/vendor_external/jquery.flot.min.js','https://s3-media1.fl.yelpcdn.com/assets/2/www/js/7b2209ac9517/assets/vendor_external/jquery.flot.time.min.js','https://s3-media4.fl.yelpcdn.com/assets/2/www/js/5b0eb1180972/assets/vendor_external/jquery.flot.pie.min.js']
                yelp.www.init.biz_details.initWarButton(".biz-page-actions .war-button", false);
yelp.www.init.biz_details.initBizWebsiteTracking("4yPqqJDJOQX69gC66YUDkA");
yelp.www.init.biz_details.boyModule("/biz/peter-luger-brooklyn-2/best_of_yelp/");
yelp.www.init.biz_details.initQAExpanderLink(".js-answer-expander");
yelp.www.init.biz_details.initQABizOwnerPicModal(".js-answer-biz-owner-photo");
yelp.www.init.biz_details.initBookmark(false, "/bookmark/remove_biz_confirm", "81af02903532ce37baa119c26cdf0aeefc983566c5377b5807dfed3abc4dee0a", "ed1d77f098c7fc15bd7a3a50e4197ba2c3e5f37c57698bdc21400814c9d8a25d", "fa0385097195438ba01046d8d4f8ea3ceb725497c54b8f900b65264a307d36f5", "4yPqqJDJOQX69gC66YUDkA");
window.YelpBizDetails.initCollectionModal();
yelp.www.init.biz_details.setupEmbeddedGATracking();
yelp.www.init.biz_details.reservations("4yPqqJDJOQX69gC66YUDkA", false, "5e84cbeca1806274");
yelp.www.init.biz_details.initPlatform("BIZ_DETAILS", true);
yelp.www.init.biz_details.initPositionWithScroll();
yelp.www.init.biz_details.initMenu("4yPqqJDJOQX69gC66YUDkA");
yelp.www.init.biz_details.fromBusinessOwner();
yelp.www.init.biz_details.platformBannerGAInit();
yelp.www.init.biz_details.reviewFeedController("/biz/peter-luger-brooklyn-2/review_feed/", "/biz/peter-luger-brooklyn-2/review_feed_auto_fetch/", "/biz/peter-luger-brooklyn-2/review_highlight_feed/", "/biz/impression/peter-luger-brooklyn-2", "5f8405dd2cc925de7cc4595d6f7b74fd2facec5ebd66605399d7aa55251c9bd2", false, null, 3, 8, "20130924", null);
yelp.www.init.biz_details.initMapLightbox();
yelp.www.init.biz_details.initConsumerAlert();
yelp.www.init.biz_details.initDealExpander();
yelp.www.init.biz_details.mapBunsenLogger("4yPqqJDJOQX69gC66YUDkA");
yelp.www.init.biz_details.mapSitRepLogger("4yPqqJDJOQX69gC66YUDkA");
yelp.www.init.biz_details.initMediaShowcaseImpressionTracker(true, "/biz/impression/peter-luger-brooklyn-2", "5f8405dd2cc925de7cc4595d6f7b74fd2facec5ebd66605399d7aa55251c9bd2", false);
yelp.www.init.biz_details.initShowcaseMediaLightboxTrigger();
yelp.www.init.biz_details.initMessageTheBusiness({"mtb_search_ad": "/mtb_composer/business/4yPqqJDJOQX69gC66YUDkA/entry_point/mtb_search_ad?source_request_id=5e84cbeca1806274", "third_party_deep_link": "/mtb_composer/business/4yPqqJDJOQX69gC66YUDkA/entry_point/third_party_deep_link?source_request_id=5e84cbeca1806274", "mtb_landing_page": "/mtb_composer/business/4yPqqJDJOQX69gC66YUDkA/entry_point/mtb_landing_page?source_request_id=5e84cbeca1806274", "mtb_service_offerings": "/mtb_composer/business/4yPqqJDJOQX69gC66YUDkA/entry_point/mtb_service_offerings?source_request_id=5e84cbeca1806274", "mtb_biz_mapbox": "/mtb_composer/business/4yPqqJDJOQX69gC66YUDkA/entry_point/mtb_biz_mapbox?source_request_id=5e84cbeca1806274", "mtb_biz_widget": "/mtb_composer/business/4yPqqJDJOQX69gC66YUDkA/entry_point/mtb_biz_widget?source_request_id=5e84cbeca1806274"});
yelp.www.init.biz_details.initCoachmark();
window.YelpBizDetails.initFullStory(false, false, "", "");
yelp.www.init.biz_details.initPlatformPopup("/transaction_platform/start_order/4yPqqJDJOQX69gC66YUDkA");
yelp.www.init.biz_details.initPopularDishes("4yPqqJDJOQX69gC66YUDkA");
yelp.www.init.biz_details.initClaimStatusHover(true);
yelp.www.init.biz_details.initAdClicksStorage({"business": null, "campaign": null, "timestamp": null});
yelp.www.init.biz_details.initAdClicksStorageReader(".js-add-url-tagging a", "4yPqqJDJOQX69gC66YUDkA");
yelp.www.init.biz_details.initNotRecommendedDrawer("/not_recommended_reviews/preview?biz_id=4yPqqJDJOQX69gC66YUDkA", "/biz/impression/peter-luger-brooklyn-2", "5f8405dd2cc925de7cc4595d6f7b74fd2facec5ebd66605399d7aa55251c9bd2", false);
yelp.www.init.biz_details.relatedBusinesses();
yelp.www.init.biz_details.initRatingDetails();
yelp.www.init.biz_details.initBizToPhone();
yelp.www.init.biz_details.initLinkedServiceOffering();
yelp.www.init.biz_details.initBunsenPageViewLogging("4yPqqJDJOQX69gC66YUDkA");
</script><script src="//sb.scorecardresearch.com/beacon.js"></script><script>
                if (window.COMSCORE && window.COMSCORE['beacon']) {
                    window.COMSCORE['beacon']({'c15': '', 'c3': '', 'c2': 7130511, 'c1': 2, 'c6': '', 'c5': '', 'c4': ''});
                }
                </script>

        <script>
            (function() {
                var main = null;

                main = function (serviceName, sha) {window.yelp_react = {serviceName: serviceName,sha: sha};};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main("yelp_main.www","e8b0b8210c131c668c7eb764d3e7fe7d7a4650b8");
            })();
    </script>

    <script>                window.webLogInfo = true;
</script>

    <script>                window.yClientLogInfo = true;
</script>

                <script>
            (function() {
                var main = null;

                var main=function(){if(["iPhone","iPod","iPad"].indexOf(navigator.platform)===-1||!window.requestAnimationFrame){return}window.bannerDetectionInitialScroll=window.bannerDetectionInitialScroll||window.scrollY||-100;requestAnimationFrame(function(){window.scrollTo(0,-100);requestAnimationFrame(function(){window.scrollBy(0,1);requestAnimationFrame(function(){window.userHasYelpApp=!window.scrollY;window.scrollTo(0,window.bannerDetectionInitialScroll)})})})};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main();
            })();
    </script>
            <script>
            (function() {
                var main = null;

                var main=function loadSpice(i){var l=0;var m=1;var j=2;var g="0";var k="1";var n=function(){if(document.querySelector){return document.querySelector(".delay-spice")}var o=document.getElementsByTagName("div");var r=o.length;var q=new RegExp("(^|\\s)delay-spice(\\s|$)");for(var p=0;p<r;++p){if(q.test(o[p].className)){return o[p]}}return null};function e(){return(window._gaUserPrefs&&window._gaUserPrefs["ioo"]&&typeof(window._gaUserPrefs["ioo"])==="function"&&window._gaUserPrefs["ioo"]()===true||false)?"1":"0"
}function f(){var o=["__webdriver_evaluate","__selenium_evaluate","__webdriver_script_function","__webdriver_script_func","__webdriver_script_fn","__fxdriver_evaluate","__driver_unwrapped","__webdriver_unwrapped","__driver_evaluate","__selenium_unwrapped","__fxdriver_unwrapped",];var q=["_phantom","__nightmare","_selenium","callPhantom","callSelenium","_Selenium_IDE_Recorder",];var r=0;for(var p=0;p<q.length;p++){if(typeof window[q[p]]!=="undefined"){r+=1}}for(var p=0;p<o.length;p++){if(typeof window.document[o[p]]!=="undefined"){r+=1
}}return r}function d(){if(typeof window.navigator.permissions!=="undefined"){var o=window.navigator.permissions;return(o.query.toString()!=="function query() { [native code] }"||o.query.toString.toString()!=="function toString() { [native code] }"||o.hasOwnProperty("query"))?1:0}return 2}function h(q,p){var r;if(p===l){r=["/ad_sp","ice?","r=",i]}else{r=["/sp","ice?","r=",i];if(p===j){r.push("&log_ad_spice=1")}}r.push("&pagevis="+q);r.push("&gablock="+e());if(window.webLogInfo&&"navigator" in window){r.push("&webdriver="+(("webdriver" in window.navigator&&window.navigator.webdriver)?1:0))
}if(window.yClientLogInfo){if("navigator" in window){if(typeof window.navigator.languages!=="undefined"){r.push("&languages="+window.navigator.languages.length)}if(typeof window.navigator.plugins!=="undefined"){r.push("&plugins="+window.navigator.plugins.length)}}r.push("&chrome="+((!!window.chrome)?1:0));r.push("&permissionsSpoofed="+d());r.push("&headlessVarCount="+f())}var o=new Image();o.src=r.join("")}function b(o){if(n()){h(o,m);c(o)}else{h(o,j)}}function c(o){if(n()){window.setTimeout(function(){c(o)
},500)}else{h(o,l)}}if(!document.webkitHidden){b(g)}else{var a=false;document.addEventListener("webkitvisibilitychange",function(){if(document.webkitHidden||a){return}a=true;b(k)})}};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main("5e84cbeca1806274");
            })();
    </script>

                    <script>
            (function() {
                var main = null;

                var main=function(){var a=new Image(1,1);a.onerror=a.onload=function(){a.onerror=a.onload=null};a.src=["//secure-us.imrworldwide.com/cgi-bin/m?ci=us-804377h&cg=0&cc=1&si=",escape(window.location.href),"&rp=",escape(document.referrer),"&ts=compact&rnd=",(new Date()).getTime()].join("")};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main();
            })();
    </script>
        <noscript>
            <img src="//secure-us.imrworldwide.com/cgi-bin/m?ci=us-804377h&amp;cg=0&amp;cc=1&amp;ts=noscript" width="1" height="1" alt="">
        </noscript>


    <script>            yConfig.csrf = {"FacebookTokenRefresh": "7aa5a5142029ce75e3b94fbb05548e30898203e3577f70c9cf6717d746ad26ce"};
</script>


            <!-- css-middleware: insert stylesheets -->

            <div id="ttdUniversalPixelTag290e816a69e9439f960a9588bc2ffb54" style="display:none">
                <script src="https://js.adsrvr.org/up_loader.1.1.0.js" type="text/javascript"></script>
                <script>
            (function() {
                var main = null;

                var main=function(){(function(b){if(typeof TTDUniversalPixelApi==="function"){var a=new TTDUniversalPixelApi();a.init("igcouad",["mvh4ai1"],"https://insight.adsrvr.org/track/up","ttdUniversalPixelTag290e816a69e9439f960a9588bc2ffb54")}})(this)};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main();
            })();
    </script>
            </div>

            <noscript>
                <img height="1" width="1" style="display:none" src="https://www.facebook.com/tr?id=102029836881428&ev=PageView&noscript=1"/>
            </noscript>

                <script>
            (function() {
                var main = null;

                var main=function(){var c=Math.random()+"";var b=c*10000000000000;document.write('<iframe src="https://6372968.fls.doubleclick.net/activityi;src=6372968;type=invmedia;cat=qr3hlsqk;dc_lat=;dc_rdid=;tag_for_child_directed_treatment=;ord='+b+'?" width="1" height="1" frameborder="0" style="display:none"></iframe>')};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main();
            })();
    </script>
    <noscript>
        <iframe src="https://6372968.fls.doubleclick.net/activityi;src=6372968;type=invmedia;cat=qr3hlsqk;dc_lat=;dc_rdid=;tag_for_child_directed_treatment=;ord=1?" width="1" height="1" frameborder="0" style="display:none"></iframe>
    </noscript>


    <script>
            (function() {
                var main = null;

                var main=function(a){window.yelp_bunsen_logger.bunsen_logger.init(a)};

                if (main === null) {
                    throw 'invalid inline script, missing main declaration.';
                }
                main({"context": [{"schema_id": 279661, "payload_data": {"guv": "DE1CA55AE6AF7DF3"}}, {"schema_id": 171864, "payload_data": {"unique_request_id": "5e84cbeca1806274"}}, {"schema_id": 280941, "payload_data": {"internal_ip": false}}, {"schema_id": 280599, "payload_data": {"country": "US", "language": "en"}}, {"schema_id": 283288, "payload_data": {"user_id": null}}, {"schema_id": 311701, "payload_data": {"platform": "www"}}, {"schema_id": 284174, "payload_data": {"normed_app_version": null}}, {"schema_id": 284175, "payload_data": {"normed_device_type": null}}, {"schema_id": 312263, "payload_data": {"product": "consumer", "product_version": "re8b0b8210c-deploy-safe-safe-safe"}}, {"schema_id": 309536, "payload_data": {"unique_view_id": "5e84cbeca1806274", "previous_unique_view_id": null}}, {"schema_id": 311448, "payload_data": {"biz_user_id_encid": null}}, {"schema_id": 311702, "payload_data": {"os_version": "x.x.x", "os_name": "unknown"}}, {"schema_id": 306745, "payload_data": {"name": "www"}}, {"schema_id": 311859, "payload_data": {"page_id": "9c8b6d9f740d50f59a4c50db8f885ad31447f635"}}, {"schema_id": 311757, "payload_data": {"is_interactive": true}}, {"schema_id": 311510, "payload_data": {"user_id_encid": null}}, {"schema_id": 311758, "payload_data": {"interface_version": "x.x.x", "interface_name": "unknown"}}, {"schema_id": 312281, "payload_data": {"offset": null}}, {"schema_id": 311756, "payload_data": {"hardware_model": null}}]});
            })();
    </script>
    </body>
</html>
`

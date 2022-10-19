<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

class ButtonLink extends BaseHTMLHelper
{
    public static $aDefaultAttrs = [
        "class" => "btn",
        "role" => "button",
    ];

    public function __invoke($sTitle, $sHref="#", $sClassType="secondary", $aAttrs=[])
    {
        $aAttrs['href'] = $sHref;
        $aAttrs = static::fnPrepareAttrs($aAttrs, static::$aDefaultAttrs);
        $aAttrs['class'] .= " btn-{$sClassType}";

        static::fnPrint(static::fnRenderTag(static::T_A, false, $aAttrs, $sTitle));
    }
}


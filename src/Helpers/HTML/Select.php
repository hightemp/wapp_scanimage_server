<?php

namespace Hightemp\WappSnanimageServer\Helpers\HTML;

class Select extends BaseHTMLHelper
{
    public static $aDefaultAttrs = [
        "class" => "form-select",
    ];

    public function __invoke($sName, $aList, $sSelectedValue="", $aAttr=[])
    {
        $aHTML = [];
        $aAttr['name'] = $sName;
        $aAttrs = static::fnPrepareAttrs($aAttr, static::$aDefaultAttrs);

        foreach ($aList as $sK => $mV) {
            if (is_array($mV)) {
                $aSelected = $sSelectedValue==$mV[0] ? ["selected" => "true"] : [];
                // NOTE: $mV = [ '1001', 'Подпись для значения' ]
                $aHTML[] = static::fnRenderTag('option', false, [ "value" => $mV[0], ...$aSelected ], $mV[1]);
            } else {
                $aSelected = $sSelectedValue==$sK ? ["selected" => "true"] : [];
                // NOTE: $mV = 'Подпись для значения'
                $aHTML[] = static::fnRenderTag('option', false, [ "value" => $sK, ...$aSelected ], $mV);
            }
        }

        static::fnPrint(static::fnRenderTag(static::T_SELECT, false, $aAttrs, join("", $aHTML)));
    }
}
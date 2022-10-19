<?php

namespace Hightemp\WappSnanimageServer;

use RedBeanPHP\Facade as R;

class Model
{
    public const TABLE = '';

    public const D_PAGE = 1;
    public const D_PAGE_SIZE = 10;
    public const CN_INDEX_COUMN = 'id';
    public const CN_CREATED_AT_COUMN = 'created_at';
    public const CN_UPDATED_AT_COUMN = 'updated_at';
    public const COLUMNS = [];

    public static function fnWipe()
    {
        return R::wipe(static::TABLE);
    }

    public static function fnDelete($aIDs)
    {
        return R::trashBatch(static::TABLE, $aIDs);
    }

    public static function fnDeleteElements($aBeans)
    {
        return R::trashAll($aBeans);
    }

    public static function fnFindOne($sql = '', $bindings=[])
    {
        return R::findOne(static::TABLE, $sql, $bindings);
    }

    public static function fnFindAll($sql = '', $bindings=[])
    {
        return R::findAll(static::TABLE, $sql, $bindings);
    }

    public static function fnExport($sql = '', $bindings=[])
    {
        return R::exportAll(static::fnFindAll($sql, $bindings));
    }

    public static function fnFindOrCreate($like = array(), $sql = '', &$hasBeenCreated = false)
    {
        return R::findOrCreate(static::TABLE, $like, $sql, $hasBeenCreated);
    }

    public static function fnCreateOrUpdate($aData=[])
    {
        $aList = [];
        if (isset($aData[static::CN_INDEX_COUMN])) {
            $aList = R::findLike(static::TABLE, $aData);
        }
        $oItem = null;

        $aData[static::CN_UPDATED_AT_COUMN] = date(DATE_FORMAT);

        if ($aList) {
            $oItem = array_shift($aList);
        } else {
            $aData[static::CN_CREATED_AT_COUMN] = date(DATE_FORMAT);
            $oItem = R::dispense(static::TABLE);
        }

        $oItem->import($aData);
        R::store($oItem);

        return $oItem;
    }

    public static function fnCreateOrUpdateFromList($aList=[])
    {
        R::begin();
        foreach ($aList as $aItem) {
            static::fnCreateOrUpdate($aItem);
        }
        R::commit();
    }

    public static function fnPagination($iPage, $iRows, $bUseOffset=false)
    {
        if ($bUseOffset) return " LIMIT {$iPage}, {$iRows}";
        $iF = ($iPage-1)*$iRows;
        return " LIMIT {$iF}, {$iRows}";
    }

    public static function fnGenerateFilterRules($aFilterRules)
    {
        $sSQL = "";

        foreach ($aFilterRules as $aRule) {
            $aRule = (array) $aRule;
            if ($aRule["op"] == "contains") {
                $sSQL .= " {$aRule["field"]} LIKE '%{$aRule["value"]}%' ";
            }
        }

        return $sSQL;
    }

    public static function fnGenerateFilterRulesForSearch($sSearch)
    {
        $aSQL = [];

        foreach (static::COLUMNS as $sColumnName => $sColumnClass) {
            $aSQL[] = "{$sColumnName} LIKE '%{$sSearch}%'";
        }

        return join(" OR ", $aSQL);
    }

    public static function fnGenerateFilterRulesForFilter($mFilter)
    {
        $aSQL = [];

        $aFilter = is_array($mFilter) ? $mFilter : json_decode($mFilter);

        foreach ($aFilter as $sColumnName => $sFilter) {
            if (isset(static::COLUMNS[$sColumnName])) {
                if (str_ends_with($sColumnName, "_id")) {
                    $aSQL[] = "{$sColumnName} = '{$sFilter}'";
                } else {
                    $aSQL[] = "{$sColumnName} LIKE '%{$sFilter}%'";
                }
            }
        }

        return join(" OR ", $aSQL);
    }

    public static function fnListWithPagination($aParams=[], $bUseTags=null)
    {
        $sFilterRules = " 1 = 1";
        $sSort = " ORDER BY id DESC";
        $sOffset = "";

        if (isset($aParams['filterRules']) && $aParams['filterRules']) {
            $aParams['filterRules'] = json_decode($aParams['filterRules']);
            $sFilterRules = static::fnGenerateFilterRules($aParams['filterRules']);
        }

        if (isset($aParams['search']) && $aParams['search']) {
            $sFilterRules = static::fnGenerateFilterRulesForSearch($aParams['search']);
        }

        if (isset($aParams['filter']) && $aParams['filter']) {
            $sFilterRules = static::fnGenerateFilterRulesForFilter($aParams['filter']);
        }

        $iPage = static::D_PAGE;
        $iLimit = static::D_PAGE_SIZE;
        $iOffset = ($iPage-1)*$iLimit;

        if (isset($aParams['offset'])) {
            $iOffset = (int) $aParams['offset'];
            if (isset($aParams['limit']) && $aParams['limit'] > 0) {
                $iLimit = (int) $aParams['limit'];
            }
            $iPage = ceil($iOffset/$iLimit);
        }
        if (isset($aParams['page'])) {
            $iPage = (int) $aParams['page'];
            if (isset($aParams['rows']) && $aParams['rows'] > 0) {
                $iLimit = (int) $aParams['rows'];
            }
            $iOffset = ($iPage-1)*$iLimit;
        }

        $sOffset = " LIMIT {$iOffset}, {$iLimit}";

        if (isset($aParams['sort'])) {
            $sSort = " ORDER BY ".$aParams['sort'];
            $sOrder = "DESC";

            if (isset($aParams['order'])) {
                $sSort = $sSort." ".(strtolower($aParams['order'][0]) == 'd' ? "DESC" : "ASC");
            }
        }

        $aResult = [];

        $aItems = R::findAll(static::TABLE, "{$sFilterRules} {$sSort} {$sOffset}", []);

        $aResult['total'] = R::count(static::TABLE, "{$sFilterRules}");
        $aResult['totalNotFiltered'] = R::count(static::TABLE, "1 = 1");

        $aResult['current_page'] = $iPage;
        $aResult['total_pages'] = ceil($aResult['total'] / $iLimit);

        $aResult['urls'] = [];

        // if ((is_null($bUseTags) && $this->$bUseTags) || $bUseTags === true) {
        //     foreach ($aItems as $oItem) {
        //         $oItem->tags = $this->fnGetTagsAsStringList($oItem->id, $this->isset($sTableName)) ?: '';
        //     }
        // }

        $aResult['rows'] = array_values((array) $aItems);

        return $aResult;
    }
}
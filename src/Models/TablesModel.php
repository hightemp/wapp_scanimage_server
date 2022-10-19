<?php

namespace Hightemp\WappSnanimageServer\Models;

use RedBeanPHP\Facade as R;
use Hightemp\WappSnanimageServer\Model;

class TablesModel extends Model
{
    public const TABLE = 'tables';

    public const COLUMNS = [
        'created_at' => [],
        'updated_at' => [],
        'name' => [],
        'description' => [],
    ];
}